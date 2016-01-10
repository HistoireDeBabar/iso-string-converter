# iso-string-converter
Provides a simple converter to change an ISO time to a Time object in Go.

#### Documentation

Given an ISO string format, e.g. "2016-Feb-21", converts to a Time object.

#### Import

`
    import (
      "github.com/HistoireDeBabar/iso-string-converter"
    )
`

#### Usage

`
    time := isoConverter.IsoStringToDate("2011-Oct-14")
`

Time will be default (Zero) time if string can't be parsed.
See Time godoc https://golang.org/pkg/time/
