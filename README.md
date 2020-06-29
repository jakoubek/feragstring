[![Go Report Card](https://goreportcard.com/badge/github.com/jakoubek/feragstring)](https://goreportcard.com/report/github.com/jakoubek/feragstring)

# feragstring

*feragstring* is a Go package for creating a FERAG string file from a JSON data structure. If you don't know what FERAG (the company) or a FERAG string is you probably don't need this package ;-)

## The shortest possible FERAG string

According to FERAG's documentation this is the shortest possible FERAG string:

```
%2440+40DEMO2009!
%2401+11E1_ROUTE_100 !
%2402+11E1_ROUTE_100 +590+91000000+20E1                            !
%2403+12R100RE001 +1300123!
%2406+11E1_ROUTE_100 !
%2441+40DEMO2009!
```

The variable values are:

- the title is 'DEMO2009'
- a route named 'E1_ROUTE_100'
- an edition called 'E1'
- a production drop 'R100RE001' with 123 copies

## Usage

```go
fs := feragstring.NewFeragString()
fs.SetTitleName("EDITION1")

fs.TitleInfo.SetPrintObjectName("EDITION1A")
fs.TitleInfo.SetPublicationDate("2020-05-31")
fs.TitleInfo.SetCountryCode("13")
fs.TitleInfo.SetPrintObjectColor("00000000")

pr1 := feragstring.NewProductReference()
pr1.SetProductName("MAIN")
pr1.SetCopiesAssigned(25000)
pr1.SetSupervision(1)
pr1.SetOverlap(5)
mp := feragstring.NewMissingParameter(1, 1)
pr1.SetMissingParameter(mp)
pr1.SetIssueReference("MAIN01")
fs.AddProductReference(pr1)

rl1 := feragstring.NewRouteListEntry()
rl1.SetRouteName("ROUTE001")
rl1.SetRouteCode(fs.NextRouteCode())
rl1.SetRampNumber(0)
rl1.SetCopiesInRoute(1500)
fs.AddRouteListEntry(rl1)
``` 

## Supported messages

- Title Info (%2440)
- Title End (%2441)
- Product Reference (%2450)
- Route List Entry (%2401)
- Route Info (%2402)
- Production Drop (%2403)
- Topsheet Data for TSL (%2414)
- Route End (%2406)

## Installation

```bash
go get -u github.com/jakoubek/feragstring
``` 