package KML

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
		<LabelStyle>
			<Scale>0</Scale>
		</LabelStyle>
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

func Test_XmlMarshal_IconStyle_1(t *testing.T) {
	s := IconStyle{
		XMLName:  xml.Name{Space: "", Local: "IconStyle"},
		Scale:    0.1,
		KMLColor: "12452450",
		IconHref: "http://i-am.url.com"}
	d, err := xml.Marshal(&s)
	if err != nil {
		t.Error("Marshal Style-IconStyle failed")
	} else {
		t.Logf("%s", string(d))
	}

}

func Test_XmlUnmarshal_Style_1(t *testing.T) {
	raw := `
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
  `
	s := Style{
		XMLName: xml.Name{Space: "", Local: "Style"},
		IconStyle: IconStyle{
			XMLName:  xml.Name{Space: "", Local: "IconStyle"},
			Scale:    0.1,
			KMLColor: "12452450",
			IconHref: "http://i-am.url.com"},
		LabelStyleScale: 0.0}
	var unmarshal Style
	xml.Unmarshal([]byte(raw), &unmarshal)
	if s != unmarshal {
		t.Errorf("Unmarshal Point , raw data: %s\n, Expect %#v\n, Result %#v", raw, s, unmarshal)
	} else {
		t.Log("Test_XmlUnmarshal_Style_1 PASS")
	}
}

func Test_XmlMarshal_Style_1(t *testing.T) {
	s := Style{
		XMLName: xml.Name{Space: "", Local: "Style"},
		IconStyle: IconStyle{
			XMLName:  xml.Name{Space: "", Local: "IconStyle"},
			Scale:    0.1,
			KMLColor: "12452450",
			IconHref: "http://i-am.url.com"},
		LabelStyleScale: 0.0}
	d, err := xml.Marshal(&s)
	if err != nil {
		t.Error("Marshal Style failed")
	} else {
		t.Logf("%s", string(d))
	}
}

func Test_Style_Color_ToRgba(t *testing.T) {
	i := IconStyle{
		XMLName:  xml.Name{Local: "IconStyle"},
		KMLColor: "1234aced",
		Scale:    0.1,
		IconHref: "http://any.url.com",
	}

	expect := "edac3412"
	if expect != ToRgba(i.KMLColor) {
		t.Error("Color : ToRgba() failed")
	}
}

func Test_Style_Color_SetRgba(t *testing.T) {
	rgba := "aabb1122"
	KMLColor := "2211bbaa"
	if ToKmlColor(rgba) != KMLColor {
		t.Error("Color ToKmlColor failed")
	}

}
