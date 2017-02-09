package KML

import (
	"encoding/xml"
)

type Placemark struct {
	XMLName xml.Name `xml:"Placemark"`
	Name    string   `xml:"name,emitempty"`
	Style   Style    `xml:"Style"`
	ExtData ExtData  `xml:"ExtendedData,emitempty"`
	Point   Point    `xml:"Point"`
}
