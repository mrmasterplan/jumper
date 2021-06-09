package mapxml

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var fullTestFile string = "testdata/Voyager.tmx"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestReadTiledMap(t *testing.T) {
	dat, err := ioutil.ReadFile(fullTestFile)
	check(err)
	tm := ReadTiledMap(dat)

	if len(tm.Properties.Properties) != 6 {
		t.Error(`not 6 properties`)
	}
	if tm.Width != 516 {
		t.Error(`wrong width`)
	}
	if !tm.Infinite {
		t.Error(`not infinite`)
	}
	for _,ts :=range tm.Tilesets {
		fmt.Printf("TileSet name \"%v\"\n", ts)
	}

	
}
