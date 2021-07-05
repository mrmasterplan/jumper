
package TmxTileset

import (
	"testing"
)

var testTilesets= [...]string{"../../../tilesets/jumper.tsx"}

func TestTilesetParsingp(t *testing.T) {
	for _, filename := range testTilesets {
		ReadTileSetFile(filename)
	} 
}