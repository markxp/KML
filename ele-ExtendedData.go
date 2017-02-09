package KML

import (
	"encoding/xml"
)

type ExtData struct {
	XMLName xml.Name   `xml:"ExtendedData"`
	Data    []DataType `xml:"Data"`
}

type DataType struct {
	XMLName xml.Name `xml:"Data"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,emitempty"`
}

type FieldError struct {
	Msg string
}

func (e *FieldError) Error() string {
	return e.Msg
}

func (dataHolder *ExtData) Field(kind string) (string, error) {
	// find 'kind' in 'dataHolder''s Data[]
	for _, d := range dataHolder.Data {
		if d.Name == kind {
			return d.Value, nil
		}
	}
	return "", dataHolder.fieldError("No Mapping field.")
}

func (d *ExtData) fieldError(msg string) error {
	return &FieldError{Msg: msg}
}

func (dataHolder *ExtData) SetField(kind string, val string) error {
	for k, d := range dataHolder.Data {
		if d.Name == kind {
			dataHolder.Data[k].Value = val
			return nil
		}
	}
	return dataHolder.fieldError("No Mapping field.")
}
