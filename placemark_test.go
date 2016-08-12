package kml

import (
	"encoding/xml"
	"testing"
)

func Test_XmlUnmarshal_Placemark_1(t *testing.T) {
	raw := `
    <Placemark>
      <name>dayAndNight</name>
      <Style></Style>
      <Point></Point>
    </Placemark>
  `
	var v Placemark
	xml.Unmarshal([]byte(raw), v)
	t.Error(v)
}
