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


Example Output

    Parameters good - fetching data
    map[cmd:[logon] email:[Company] password:[*******************]]
    Error:
    Token: ÆÆÆÆÆÆÆÆÆÆÆÆÆÆÆÆÆÆÆÆÆÆÆÆÆ
    parsing time "" as "2006-01-02T15:04:05Z07:00": cannot parse "" as "2006"
    2016-03-01 00:00:00 -0600 CST
    2016-03-23 00:00:00 -0500 CDT
    Mar 1, 2016 at 6:00am (UTC)
    Mar 23, 2016 at 5:00am (UTC)
    Done - Bye Bye

Notes
----------------------
You will have to edit the app.yaml file and provide the correct
credentials and site.

Currently you can just edit the code to change the dates.  I will
add CLI parameters later.

Testing on windows I had to use to pull down the code.

  go get -v github.com/craignicholson/fogbugz/fogbugz

```go

  export("2016-01-01", "2016-03-24")

```
More information about the timezone can be found in the
github.com/craignicholson/fogbugz/fogbugz package.
