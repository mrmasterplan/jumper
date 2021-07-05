
package config

// package provides global config parameters in case I ever want to change them

type XY struct {
	X int
	Y int
}

var GameWindow = XY{320, 240}
var GameWindowWorldView = XY{640, 480}
