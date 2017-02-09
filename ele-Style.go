package KML

import "encoding/xml"

type Style struct {
	XMLName   xml.Name  `xml:"Style"`
	IconStyle IconStyle `xml:"IconStyle,emitempty"`
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

func ToRgba(KMLcolor string) string {
	a := KMLcolor[0:2]
	b := KMLcolor[2:4]
	g := KMLcolor[4:6]
	r := KMLcolor[6:8]
	rgba := r + g + b + a
	return rgba
}

func ToKmlColor(rgba string) string {
	r := rgba[0:2]
	g := rgba[2:4]
	b := rgba[4:6]
	a := rgba[6:8]
	KMLcolor := a + b + g + r
	return KMLcolor
}
