package mapxml

import (
	// "fmt"
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
	tm := readTiledMap(dat)

	if len(tm.Properties.Properties) != 6 {
		t.Error(`not 6 properties`)
	}
	if tm.Width != 516 {
		t.Error(`wrong width`)
	}
	if !tm.Infinite {
		t.Error(`not infinite`)
	}
	fmt.Printf(`The name %v`,tm.Tilesets[0].Name)

}
