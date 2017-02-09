package KML

import (
	"encoding/xml"
)

// BUG: cannot set xmlns:gx, xmlns:kml, xmlns:atom

type KML struct {
	XMLName xml.Name `xml:"kml"`

	Xmlns     string `xml:"xmlns,attr"`
	XmlnsGx   string `xml:"xmlns:gx,attr"`
	XmlnsKml  string `xml:"xmlns:kml,attr"`
	XmlnsAtom string `xml:"xmlns:atom,attr"`

	Folder Folder `xml:"Folder,emitempty"`
}
