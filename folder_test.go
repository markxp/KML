package kml

import (
	"encoding/xml"
	"testing"
)

func Test_XmlMarshal_Folder_1(t *testing.T) {
	f := &Folder{
		Name: "New expt",
	}
	p := Placemark{
		Name: "Taipei 101",
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
	t.Log(f)
}

func Test_XmlUnmarshal_Folder_1(t *testing.T) {
	raw := `
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
  `
	var f Folder
	err := xml.Unmarshal([]byte(raw), &f)
	if err != nil {
		t.Error("UnMarshal Folder failed")
	} else if f.Name != "New expt" ||
		len(f.Placemarks) != 1 ||
		f.Placemarks[0].Point.CoordData != "1.0,-1.0,+0.35" {
		t.Errorf("UnMarshal Folder failed. inner error \n%v", f)
	}
}
