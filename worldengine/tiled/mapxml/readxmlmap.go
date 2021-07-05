package mapxml

import (
	"io/ioutil"
	"encoding/xml"
)


func ReadTiledMapFile(path string) TiledMap {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	// m := fullMap{}
	var tiledMap TiledMap

	if err := xml.Unmarshal(dat, &tiledMap); err != nil {
		panic(err)
	}


	return tiledMap
}
