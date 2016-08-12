package kml

import (
	"encoding/xml"
	"testing"
)

func Test_XmlUnmarshal_IconStyle_1(t *testing.T) {
	raw := `
  <IconStyle>
    <color>12452450</color>
    <scale>0.1</scale>
    <Icon>
      <href>http://i-am.url.com</href>
    </Icon>
  </IconStyle>
  `
	s := IconStyle{
		XMLName:  xml.Name{Space: "", Local: "IconStyle"},
		Scale:    0.1,
		KMLColor: "12452450",
		IconHref: "http://i-am.url.com"}

	var unmarshal IconStyle
	xml.Unmarshal([]byte(raw), &unmarshal)
	if s != unmarshal {
		t.Errorf("Unmarshal Point, raw data: %s, Expect %#v\n, Result %#v", raw, s, unmarshal)
	} else {
		t.Log("Test_XmlUnmarshal_IconStyle_1 PASS")
	}
}

func Test_XmlUnmarshal_Style_1(t *testing.T) {
	raw := `
  <Style>

  </Style>
  `
	s := Style{XMLName: xml.Name{Space: "", Local: "Style"},

		LabelStyleScale: 0.0}
	var unmarshal Style
	xml.Unmarshal([]byte(raw), &unmarshal)
	if s != unmarshal {
		t.Errorf("Unmarshal Point , raw data: %s\n, Expect %#v\n, Result %#v", raw, s, unmarshal)
	} else {
		t.Log("Test_XmlUnmarshal_Style_1 PASS")
	}

}
