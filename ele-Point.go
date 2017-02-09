package KML

import (
	"encoding/xml"
	"errors"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	XMLName   xml.Name `xml:"Point"`
	CoordData string   `xml:"coordinates"`
	Mode      string   `xml:"altitudeMode,emitempty"`
	// if <altitudeMode></altitudeMode> is empty, Marshal() still
	// leaves empty tag on Point
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

func (p *Point) SetCoordinates(c [3]float64) {
	strArr := make([]string, 3)
	for i := 0; i != 3; i++ {
		strArr[i] = strconv.FormatFloat(c[i], 'f', -1, 64)
	}
	p.CoordData = strings.Join(strArr, ",")
}

func (src *Point) Distance(dst Point) (ret float64) {
	srcArr, _ := src.Coordinates()
	dstArr, _ := dst.Coordinates()

	return Distance(srcArr[0], srcArr[1], dstArr[0], dstArr[1])
}

func Distance(srcX, srcY, dstX, dstY float64) (ret float64) {
	earthRadius := 6357000.0
	srcLambda := srcX / 180 * math.Pi
	srcPhi := srcY / 180 * math.Pi
	dstLambda := dstX / 180 * math.Pi
	dstPhi := dstY / 180 * math.Pi
	theta := math.Acos(
		math.Sin(srcPhi)*math.Sin(dstPhi) +
			math.Cos(srcPhi)*math.Cos(dstPhi)*math.Cos(dstLambda-srcLambda),
	)
	ret = earthRadius * theta
	return ret
}
