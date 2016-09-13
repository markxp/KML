package kml

type Placemark struct {
	Name    string  `xml:"name,emitempty"`
	Style   Style   `xml:"Style"`
	ExtData ExtData `xml:"ExtendedData,emitempty"`
	Point   Point   `xml:"Point"`
}
