package kml

type Placemark struct {
	Name    string `xml:"name,emitempty"`
	Style   Style
	ExtData ExtData `xml:",emitempty"`
	Point   Point
}
