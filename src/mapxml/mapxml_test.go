package mapxml

import (
	"fmt"
	"testing"
)

func TestReadTiledMap(t *testing.T) {

	tm := ReadTiledMapFile("testdata/Voyager.tmx")

	if len(tm.Properties.Properties) != 6 {
		t.Error(`not 6 properties`)
	}
	if tm.Width != 516 {
		t.Error(`wrong width`)
	}
	if !tm.Infinite {
		t.Error(`not infinite`)
	}
	for _, ts := range tm.Tilesets {
		fmt.Printf("TileSet name \"%v\"\n", ts)
	}

}

func TestReadTileSet(t *testing.T) {

	ts := ReadTileSetFile("../../tilesets/jumper1.tsx")

	prnt := func(ts TMXTileset) {
		// fmt.Println(ts)
		fmt.Println(ts.Name)
		if ts.Image != nil {
			fmt.Println(ts.Image.Source)

		}
	}
	prnt(ts)

	ts = ReadTileSetFile("../../tilesets/cross.tsx")
	prnt(ts)

}
