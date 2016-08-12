package kml

import (
	//	"encoding/xml"

	"encoding/xml"
	"testing"
	//	. "github.com/markxp/parser/kml"
)

func Test_XmlUnmarshal_Point_1(t *testing.T) {
	raw := `
  <Point>
    <coordinates>1.0,-1.0,+0.35</coordinates>
    <altitudeMode>relativeToGround</altitudeMode>
  </Point>
`
	p := Point{XMLName: xml.Name{Space: "", Local: "Point"},
		CoordData: "1.0,-1.0,+0.35",
		Mode:      "relativeToGround"}

	var unmarshal Point
	xml.Unmarshal([]byte(raw), &unmarshal)
	if p != unmarshal {
		t.Errorf("Unmarshal Point , raw data: %s\n, Expect %v\n, Result %v", raw, p, unmarshal)
	} else {
		t.Log("Test_XmlUnmarshal_Point_1 PASS")
	}

}

func Test_XmlUnmarshal_Point_2(t *testing.T) {
	raw := `
  <Point>
    <coordinates>1.0,-1.0,+0.35</coordinates>
  </Point>
`
	p := Point{XMLName: xml.Name{Space: "", Local: "Point"},
		CoordData: "1.0,-1.0,+0.35",
		Mode:      ""}

	var unmashal Point
	xml.Unmarshal([]byte(raw), &unmashal)
	if p != unmashal {
		t.Errorf("Unmarshal Point , raw data: %s\n, Expect %#v\n, Result %#v", raw, p, unmashal)
	} else {
		t.Log("Test_XmlUnmarshal_Point_2 PASS")
	}

}
