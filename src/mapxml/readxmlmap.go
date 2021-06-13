package mapxml

import (
	"encoding/xml"
)


func ReadTiledMap(ba []byte) TiledMap {
	// m := fullMap{}
	var tiledMap TiledMap

	if err := xml.Unmarshal(ba, &tiledMap); err != nil {
		panic(err)
	}


	return tiledMap
}
