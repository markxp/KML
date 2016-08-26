package kml

import (
	"encoding/xml"
	"testing"
)

const raw = `
		<ExtendedData>
			<Data name="CGI">
				<value>466561000000001</value>
			</Data>
			<Data name="CELLNAME">
			</Data>
			<Data name="CELLID">
				<value>1</value>
			</Data>
			<Data name="eNB">
				<value>521</value>
			</Data>
			<Data name="TAC">
				<value>10000</value>
			</Data>
			<Data name="TECHNOLOGY">
				<value>4G</value>
			</Data>
			<Data name="MODE">
				<value>LTE</value>
			</Data>
			<Data name="RSRP">
				<value>-80 dBm</value>
			</Data>
			<Data name="RSRQ">
				<value>-7</value>
			</Data>
			<Data name="LAYER">
				<value>1</value>
			</Data>
			<Data name="PC">
				<value>0</value>
			</Data>
			<Data name="DL_BITRATE">
				<value>0 kbps</value>
			</Data>
			<Data name="UL_BITRATE">
				<value>0 kbps</value>
			</Data>
			<Data name="SNR">
				<value>30.0</value>
			</Data>
			<Data name="LTE RSSI">
				<value>-66</value>
			</Data>
			<Data name="PHONE STATE">
				<value>I</value>
			</Data>
			<Data name="SPEED">
				<value>17km/h</value>
			</Data>
			<Data name="HEIGHT">
				<value>185m</value>
			</Data>
			<Data name="ACCURACY">
				<value>5m</value>
			</Data>
			<Data name="TIME">
				<value>2016.07.19_11.46.58</value>
			</Data>
		</ExtendedData>
  `

func Test_XmlUnmarshal_ExtData_1(t *testing.T) {
	v := &ExtData{}
	err := xml.Unmarshal([]byte(raw), v)
	if (v.XMLName == xml.Name{Space: "", Local: ""} || len(v.Data) == 0 || err != nil) {
		t.Errorf("UnMarshal ExtendedData failed")
	} else if len(v.Data) != 20 {
		t.Errorf("UnMarshal ExtendedData failed. Partially lost")
	} else {
		t.Log("UnMarshal ExtendedData OK")
	}
}

func Test_XmlMarshal_ExtData_1(t *testing.T) {
	v := ExtData{
		XMLName: xml.Name{Local: "ExtendedData"},
	}

	v.Data = append(v.Data, DataType{Name: "BBA", Value: "-1"})
	v.Data = append(v.Data, DataType{Name: "AAB", Value: "hello"})

	marshal, err := xml.Marshal(&v)
	if err != nil {
		t.Error("Marshal ExtData failed")
	}

	t.Logf("%s", string(marshal))
}

func Test_SetField(t *testing.T) {
	v := ExtData{
		XMLName: xml.Name{Local: "ExtendedData"},
	}

	v.Data = append(v.Data, DataType{Name: "BBA", Value: "-1"})
	v.Data = append(v.Data, DataType{Name: "AAB", Value: "hello"})

	v.SetField("AAB", "world")
	found := false
	for _, d := range v.Data {
		if d.Name == "AAB" &&
			d.Value == "world" {
			found = true
			t.Log("func SetField(kind string) string PASS")
			break
		}
	}
	//v.Data is empty
	if len(v.Data) == 0 {
		t.Error("ExtData error")
	}
	if !found {
		t.Error("Invalid field or SetField() failed")
	}

}

func Test_Field(t *testing.T) {
	v := ExtData{
		XMLName: xml.Name{Local: "ExtendedData"},
	}

	v.Data = append(v.Data, DataType{Name: "BBA", Value: "-1"})
	v.Data = append(v.Data, DataType{Name: "AAB", Value: "hello"})
	a1, _ := v.Field("AAB")
	a2, _ := v.Field("BBA")
	if a1 == "hello" && a2 == "-1" {
		t.Log("func Field(kind string) string  PASS")
	}
}

/*
func Test_XmlMarshal_ExtData(t *testing.T) {

}
*/
