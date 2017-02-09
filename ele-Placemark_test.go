package KML

import (
	"encoding/xml"
	"testing"
)

func Test_XmlUnmarshal_Placemark_1(t *testing.T) {
	raw := `
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
  `
	var v Placemark
	xml.Unmarshal([]byte(raw), &v)
	expect := Placemark{
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
			Mode:      "relativeToGround"}}
	/*
		if v != expect {
			t.Errorf("Unmarshal Placemark , raw data: %s\n, Expect %#v\n, Result %#v", raw, expect, v)
		} else {
			t.Log("Test_XmlUnmarshal_Placemark_1 PASS")
		}
	*/
	t.Logf("\n%+v \n\n%+v", expect, v)
}

func Test_XmlUnmarshal_Placemark_2(t *testing.T) {
	raw := `
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
			<ExtendedData>
				<Data name="SNR">
					<value>-30 dbm</value>
				</Data>
			</ExtendedData>
			<Point>
		    <coordinates>1.0,-1.0,+0.35</coordinates>
		    <altitudeMode>relativeToGround</altitudeMode>
		  </Point>
    </Placemark>
  `
	var v Placemark
	xml.Unmarshal([]byte(raw), &v)
	expect := Placemark{
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
			Mode:      "relativeToGround"}}

	// Add ExtData XMLName (optional)
	expect.ExtData.XMLName.Local = "ExtendedData"
	arr := &expect.ExtData.Data

	//Add DataType contents
	//Add DataType XMLName (optional)
	*arr = append(*arr, DataType{XMLName: xml.Name{Local: "Data"}, Name: "SNR", Value: "-30 dbm"})

	//Human check
	t.Logf("\n%+v \n\n%+v", expect, v)
}

func Test_XmlMarshal_Placemark_1(t *testing.T) {
	raw := `
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
			<ExtendedData>
				<Data name="SNR">
					<value>-30 dbm</value>
				</Data>
			</ExtendedData>
			<Point>
				<coordinates>1.0,-1.0,+0.35</coordinates>
				<altitudeMode>relativeToGround</altitudeMode>
			</Point>
		</Placemark>
	`
	var v Placemark
	xml.Unmarshal([]byte(raw), &v)
	expect := Placemark{
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
			Mode:      "relativeToGround"}}

	// Add ExtData XMLName (optional)
	expect.ExtData.XMLName.Local = "ExtendedData"
	arr := &expect.ExtData.Data

	//Add DataType contents
	//Add DataType XMLName (optional)
	*arr = append(*arr, DataType{XMLName: xml.Name{Local: "Data"}, Name: "SNR", Value: "-30 dbm"})

	data, _ := xml.Marshal(expect)
	//Human check
	t.Logf("\n%+v \n\n%+v", expect, string(data))
}
