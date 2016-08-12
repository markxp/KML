package kml

import "encoding/xml"

type Style struct {
	XMLName   xml.Name `xml:"Style"`
	IconStyle `xml:"IconStyle,innerxml,emitempty"`
	//LineStyle
	//PolyStyle
	//BalloonStyle
	LabelStyleScale float64 `xml:"LabelStyle>scale,emitempty"`
	Id              string  `xml:"id,attr,emitempty"`
}

type IconStyle struct {
	XMLName  xml.Name `xml:"IconStyle"`
	KMLColor string   `xml:"color"`
	Scale    float64  `xml:"scale"`
	IconHref string   `xml:"Icon>href"`
}

//KML color format as "ABGR", ie. "aabbggrr"
type KMLColorString string

func (p KMLColorString) ToRgba() string {
	s := string(p)
	a := s[0:2]
	b := s[2:4]
	g := s[4:6]
	r := s[6:8]
	rgba := r + g + b + a
	return rgba
}

func SetRgba(rgba string) KMLColorString {
	r := rgba[0:2]
	g := rgba[2:4]
	b := rgba[4:6]
	a := rgba[6:8]
	s := a + b + g + r
	return KMLColorString(s)
}
