package kml

type Folder struct {
	Name       string      `xml:"name,emitempty"`
	Placemarks []Placemark `xml:"Placemark"`
}
