# FogBugz Exporter Application

Introduction
------------
This is a test application which pulls the hours worked from fogbugz
and writes the hours to a a csv and json file.  The purpose is for a replacement
to the Clark Kent report in FogBugz which is missing some of the data
needed for the reports we need to run.


Installation and usage
----------------------

The import path for the packages are *github.com/craignicholson/fogbugz/fogbugz*.

To install it, run:

    go get github.com/craignicholson/fogbugz/fogbugz


To run the application

    go run main.go "2016-01-01" "2016-01-03"


Example Output

    Parameters good - fetching data
    map[cmd:[logon] email:[Company] password:[*******************]]
    parsing time "" as "2006-01-02T15:04:05Z07:00": cannot parse "" as "2006"
    2016-03-01 00:00:00 -0600 CST
    2016-03-23 00:00:00 -0500 CDT
    Mar 1, 2016 at 6:00am (UTC)
    Mar 23, 2016 at 5:00am (UTC)
    Done - Bye Bye

Configuration of the app.yaml
----------------------
You will have to edit the app.yaml file and provide the correct
credentials and rootsite url.

Timezones References
----------------------
To find your timezone use the links below.  

Timezone Lists
* https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
* https://www.iana.org/time-zones
* http://golang.org/pkg/time/#LoadLocation
* http://www.geonames.org/export/web-services.html#timezone
* /usr/go/lib/time/zoneinfo.zip

Unknowns
----------------------

Testing on windows I had to use to pull down the code.

  go get -v github.com/craignicholson/fogbugz/fogbugz
