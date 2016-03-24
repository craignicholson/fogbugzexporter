# FogBugz Exporter Application

Introduction
------------
This is a test application which pulls the hours worked from fogbugz
and writes the hours to a file.  The purpose is for a replacement
to the Clark Kent report in FogBugz which does not have all the data
needed for the reports we need to run.


Installation and usage
----------------------

The import path for the packages are *gopkg.in/yaml.v2*.

To install it dependencies, run:

    go get gopkg.in/yaml.v2
    go get github.com/craignicholson/fogbugz/fogbugz


To run the application
    go run main.go "2016-01-01" "2016-01-03"


Notes
----------------------
You will have to edit the app.yaml file and provide the correct
credentials and site.

Currently you can just edit the code to change the dates.  I will
add CLI parameters later.

```go

  export("2016-01-01", "2016-03-24")

```
More information about the timezone can be found in the
github.com/craignicholson/fogbugz/fogbugz package.
