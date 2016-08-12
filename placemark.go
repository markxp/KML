package kml

import (
	"encoding/xml"
)

type Placemark struct {
	XMLName xml.Name `xml:"Placemark"`
	Name    string   `xml:"name,emitempty"`
	Style   `xml:"Style,innerxml"`
	ExtData `xml:"ExtendedData,innerxml"`
	Point   `xml:"Point,innerxml"`
}
