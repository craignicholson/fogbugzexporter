// Copyright 2015 Craig Nicholson. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Quick app to test the package for ev

package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/craignicholson/fogbugz/fogbugz"
	"gopkg.in/yaml.v2"
)

// config stores the configurations for the api package
type config struct {
	User     string `yaml:"user,omitempty"`
	Password string `yaml:"password,omitempty"`
	Rootsite string `yaml:"rootsite,omitempty"`
	Timezone string `yaml:"timezone,omitempty"`
}

var api fogbugz.API
var configs config

func main() {

	setupAPI()

	// Receives YYYY-MM-DD in your local timezone
	// Check for required agruments
	if len(os.Args) != 3 {
		fmt.Println("Missing Parameters")
		//TODO: add more specific instructions output for user
		os.Exit(3)
	} else {
		var from = os.Args[1]
		var to = os.Args[2]

		if validateArgs(from, to, configs.Timezone) {
			fmt.Println("Parameters good - fetching data")
			export(from, to)
		}
		fmt.Println("Done - Bye Bye")
	}
}

// validateArgs checks the arguments for yyyy-mm-dd
// format and returns false is any of the arguements
// does not match this pattern
func validateArgs(startDateLocal string, endDateLocal string, timezone string) bool {
	//Try and Parse the strings into dates.
	//If any of the strings fails to pass return false
	isValid := true
	loc, errloc := time.LoadLocation(timezone)
	if errloc != nil {
		fmt.Println(errloc)
		isValid = false
	}

	const shortFormlayout = "2006-01-02"
	_, errStart := time.ParseInLocation(shortFormlayout, startDateLocal, loc)
	if errStart != nil {
		fmt.Println(errStart)
		isValid = false
	}
	_, errEnd := time.ParseInLocation(shortFormlayout, endDateLocal, loc)
	if errEnd != nil {
		fmt.Println(errEnd)
		isValid = false
	}

	return isValid
}

// setupAPI loads the configs and assigns the configs to the api
func setupAPI() {
	data, _ := ioutil.ReadFile("app.yaml")
	err := yaml.Unmarshal(data, &configs)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("error: %v", err)
	}

	//Load the configs
	u, err := url.Parse(configs.Rootsite)
	if err != nil {
		log.Fatal(err)
	}
	api.Root = u
}

// export pulls the hours and writes the results to a csv and json file.
func export(from string, to string) {
	api.Login(configs.User, configs.Password)
	//hours := api.GetHours("2016-01-01", "2016-03-16", "America/Chicago")
	hours := api.GetHours(from, to, configs.Timezone)

	api.InvalidateToken()

	data, err := json.Marshal(hours)
	if err != nil {
		fmt.Println(err)
	}
	writeJSONFile(data, "hours.json")
	writeCsvFile(hours, "hours.csv")
}

// writeJSONFile write the data to disk in JSON format.
func writeJSONFile(data []byte, filename string) {
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("error:", err)
	}
}

// writeCsvFile write the data to disk in csv format.
func writeCsvFile(data []fogbugz.Hour, filename string) {
	// Create a csv file
	f, err := os.Create("hours.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	// Write Unmarshaled json data to CSV file
	w := csv.NewWriter(f)
	// header records
	//TODO: I should be able to get all the struct names dymanically
	header := []string{"RowID", "StartDate", "EndDate", "Title", "Duration", "Expense", "Employee", "Project", "MileStone", "Customer", "CaseNumber", "BillingPeriod", "Area", "Category", "StartNote", "Year", "Month", "Day", "DOW", "Tags"}
	w.Write(header)
	//for i := 0; i < len(data); i++ {
	for i, item := range data {
		var record []string
		// This could also be hour.ID which is ixPerson
		record = append(record, strconv.Itoa(i))
		//Note: When commans occur in the string, encoding/csv wraps the string in double quotes
		//TODO: Wrap all strings in double quotes
		record = append(record, item.StartDate.Format("2006-01-02 03:04:05 PM"))
		record = append(record, item.EndDate.Format("2006-01-02 03:04:05 PM"))
		record = append(record, item.Title)
		record = append(record, strconv.FormatFloat(item.Duration, 'f', 4, 64))
		record = append(record, item.Expense)
		record = append(record, item.Employee)
		record = append(record, item.Project)
		record = append(record, item.MileStone)
		record = append(record, item.Customer)
		record = append(record, strconv.Itoa(item.CaseNumber))
		record = append(record, item.BillingPeriod)
		record = append(record, item.Area)
		record = append(record, item.Category)
		record = append(record, item.StartNote)
		record = append(record, strconv.Itoa(item.Year))
		record = append(record, item.Month)
		record = append(record, strconv.Itoa(item.Day))
		record = append(record, item.DOW)
		record = append(record, item.Tags)
		w.Write(record)
	}
	w.Flush()
}
