package kml

import (
	"encoding/xml"
	"errors"
	"strconv"
	"strings"
)

type Point struct {
	XMLName   xml.Name `xml:"Point"`
	CoordData string   `xml:"coordinates"`
	Mode      string   `xml:"altitudeMode,emitempty"`
}

const (
	RELATIVE = "relativeToGround" //default Mode
	ABSOLUTE = "absolute"
)

func (p Point) Coordinates() (ret [3]float64, err error) {

	strArr := strings.Split(p.CoordData, ",")
	s := len(strArr)
	if s < 2 || s > 3 {
		err = errors.New("Invalid coordinates in Package kml, func (Point) Coordinates() (ret [3]float64, err error)")
		return
	}
	// parser string to float64, could be 2 (x,y) or 3 (x,y,z)
	// could be buggy here
	for i := 0; i != s; i++ {
		ret[i], err = strconv.ParseFloat(strArr[i], 64)
		if err != nil {
			panic("")
		}
	}
	// if no altitude (height) value, set to 0
	if s == 2 {
		ret[2] = 0.0
	}
	return
}

func (p *Point) setCoordinates(c [3]float64) {
	var strArr []string
	for i := 0; i != 3; i++ {
		strArr[i] = strconv.FormatFloat(c[i], 'f', -1, 64)
	}
	p.CoordData = strings.Join(strArr, ",")
}
