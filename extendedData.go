package kml

import (
	"encoding/xml"
)

type ExtData struct {
	XMLName xml.Name   `xml:"ExtendedData"`
	Data    []DataType `xml:"Data"`
}

type DataType struct {
	XMLName xml.Name `xml:"Data"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,emitempty"`
}

func (dataHolder ExtData) GetExtData(kind string) string {
	// find 'kind' in 'dataHolder''s Data[]
	for _, d := range dataHolder.Data {
		if d.Name == kind {
			return d.Value
		}
	}
	return ""
}
