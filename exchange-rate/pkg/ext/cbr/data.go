package cbr

import (
	"encoding/xml"
)

type ValCurs struct {
	XMLName    xml.Name `xml:"ValCurs"`
	ID         string   `xml:"ID,attr"`
	DateRange1 string   `xml:"DateRange1,attr"`
	DateRange2 string   `xml:"DateRange2,attr"`
	Records    []Record `xml:"Record"`
}

type Record struct {
	Date      string `xml:"Date,attr"`
	Nominal   int    `xml:"Nominal"`
	Value     string `xml:"Value"`
	VunitRate string `xml:"VunitRate"`
}
