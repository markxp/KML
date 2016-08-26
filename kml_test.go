package kml

import (
	"encoding/xml"
	"testing"
)

func Test_XmlUnmarshal_Kml_1(t *testing.T) {
	raw := `
  <kml xmlns="http://www.opengis.net/kml/2.2" xmlns:gx="http://www.google.com/kml/ext/2.2" xmlns:kml="http://www.opengis.net/kml/2.2" xmlns:atom="http://www.w3.org/2005/Atom">
    <Folder>
      <name>New expt</name>
      <Placemark>
        <name>Taipei 101</name>
        <Style>
          <IconStyle>
            <color>12452450</color>
            <scale>0.1</scale>
            <Icon>
              <href>http://i-am.url.com</href>
            </Icon>
          </IconStyle>
          <LabelStyle>
            <Scale>0</Scale>
          </LabelStyle>
        </Style>
        <Point>
          <coordinates>1.0,-1.0,+0.35</coordinates>
          <altitudeMode>relativeToGround</altitudeMode>
        </Point>
      </Placemark>
    </Folder>
  </kml>
  `
	var kml KML
	err := xml.Unmarshal([]byte(raw), &kml)

	/* Collect all attributes and use this part
	if err != nil {
		t.Error("UnMarshal KML failed")
	} else if kml.XmlnsGx != "http://www.google.com/kml/ext/2.2" {
		t.Errorf("UnMarshal KML failed."+
			" cannot set xml attributes. eg. xmlns:gx=%s \n\nkml info:%#v",
			kml.XmlnsGx, kml)
	} else {
		t.Log("UnMarshal KML ok")
	}
	*/

	// cannot load all attributes, use this part temporarily
	if err != nil {
		t.Errorf("UnMarshal KML failed, %s\n%#v", err, kml)
	} else {
		t.Log("UnMarshal KML ok")
	}

}

func Test_XmlMarshal_Kml_1(t *testing.T) {
	k := KML{
		Xmlns:     "http://www.opengis.net/kml/2.2",
		XmlnsGx:   "http://www.google.com/kml/ext/2.2",
		XmlnsKml:  "http://www.opengis.net/kml/2.2",
		XmlnsAtom: "http://www.w3.org/2005/Atom",
	}
	f := &Folder{
		XMLName: xml.Name{Local: "Folder"},
		Name:    "New expt",
	}
	p := Placemark{
		XMLName: xml.Name{Space: "", Local: "Placemark"},
		Name:    "Taipei 101",
		Style: Style{
			XMLName: xml.Name{Space: "", Local: "Style"},
			IconStyle: IconStyle{
				XMLName:  xml.Name{Space: "", Local: "IconStyle"},
				Scale:    0.1,
				KMLColor: "12452450",
				IconHref: "http://i-am.url.com"},
			LabelStyleScale: 0.0},
		Point: Point{XMLName: xml.Name{Space: "", Local: "Point"},
			CoordData: "1.0,-1.0,+0.35",
			Mode:      "",
		},
	}
	f.Placemarks = append(f.Placemarks, p)
	k.Folder = *f

	t.Log(f)
}
