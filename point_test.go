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

// BUG: altitudeMode tag: either Mode is set to ""or not set.
// Marshal() shows the tag.
//
func Test_XmlMarshal_Point_1(t *testing.T) {
	p := Point{XMLName: xml.Name{Space: "", Local: "Point"},
		CoordData: "1.0,-1.0,+0.35"}

	ans := `<Point><coordinates>1.0,-1.0,+0.35</coordinates><altitudeMode></altitudeMode></Point>`

	raw, err := xml.Marshal(p)

	if err == nil {
		if string(raw) != string(ans) {
			t.Logf("%#v \n\n %#v", string(ans), string(raw))
			t.Errorf("(p *Point) Marshal ([]byte,error) failed. Byte comparsion")
		} else {
			t.Log("OK")
		}
	} else {
		t.Errorf("(p *Point) Marshal ([]byte,error) failed. ")
	}
}

func Test_Coordinates(t *testing.T) {
	//
	//<Point>
	//  <coordinates>1.0,-1.0,+0.35</coordinates>
	//</Point>
	//
	p := Point{XMLName: xml.Name{Space: "", Local: "Point"},
		CoordData: "1.0,-1.0,+0.35",
		Mode:      ""}
	v, err := p.Coordinates()
	if err != nil {
		t.Errorf("func (p *Point) Coordinates() ([3]float64 ,error) failed.")
	}
	if len(v) != 3 {
		t.Errorf("func (p *Point) Coordinates() ([3]float64 ,error) failed." +
			" \"ArraSize\"")
	}
	ans := [3]float64{1.0, -1.0, 0.35}
	for i := 0; i != len(v); i++ {
		if v[i]-ans[i] > 0.00001 {
			t.Errorf("func (p *Point) Coordinates() ([3]float64 failed, Transform Unaccurately")
		}
	}

}

func Test_SetCoordinates(t *testing.T) {
	coord := [3]float64{10, -20, 0.00001}
	p := Point{XMLName: xml.Name{Space: "", Local: "Point"},
		CoordData: "1.0,-1.0,+0.35",
		Mode:      ""}

	p.SetCoordinates(coord)
	// Check (p.x, p.y, p.z)
	if p.CoordData != "10,-20,0.00001" {
		t.Errorf("%#v", p)
	}

}
