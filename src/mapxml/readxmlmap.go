package mapxml

import (
	// "fmt"
	"encoding/xml"
)


func readTiledMap(ba []byte) TiledMap {
	// m := fullMap{}
	var tiledMap TiledMap

	if err:=xml.Unmarshal(ba,&tiledMap); err !=nil{
		panic(err)
	}


	
	return tiledMap
}

