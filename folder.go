package kml

import (
	"encoding/xml"
)

type Folder struct {
	XMLName    xml.Name    `xml:"Folder"`
	Name       string      `xml:"name,emitempty"`
	Placemarks []Placemark `xml:"Placemark"`
}
