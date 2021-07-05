package mapxml

import (
	// "encoding/xml"
)

type tmxXY struct {
	X int `xml:"x,attr"`
	Y int `xml:"y,attr"`
}
type tmxWidthHeight struct  {
	Width int `xml:"width,attr"`
	Height int `xml:"height,attr"`
}