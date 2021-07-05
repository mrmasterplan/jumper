// this file provides
// - a data structure that holds a tileset
//   - suitably optimized for use in game
// - serialization of tlied xml (TSX)
//   - tsx to datastructure (discards non-conformant data)
//   - datastructure to tsx (lossless)
//   - TSX has these advantages
//     - can be edited in Tiled
// - serialization to protobuf
//   - proto to datastructure
//   - datastructure to proto
//   - proto has these advantages:
//     - self-contained
//       - no external references
//       - can be moved anywhere
//       - can be read from anywhere
//       - contains all pictures
//     - small size and efficient serialization

package tileset

import (
	// "../../config"
	"fmt"
	"image"

	"github.com/mrmasterplan/jumper/worldengine/config"
	"github.com/mrmasterplan/jumper/worldengine/tiled/mapxml"
	// "github.com/mrmasterplan/jumper/src/config"
	// "github.com/mrmasterplan/jumper/src/mapxml"
)

func ParseTMXTileSet(tsx *mapxml.TMXTileset) (*Tileset, error) {
	// If we were read as part of a level, the level will have taken care of the FirstGID. We only need to provide the tileset.
	// if we are in a level, the tileset will not normally be embedded and have a source filled in.
	if tsx.Source != "" {
		newtsx, err := mapxml.ReadTileSetFile(tsx.Source)
		if err != nil {
			return nil, err
		}
		tsx = newtsx
	}

	ts := &Tileset{}
	ts.Name = tsx.Name

	if tsx.TileWidth != config.TileSize || tsx.TileHeight != config.TileSize {
		return nil, fmt.Errorf("XML Tileset uses incompatible tile size")
	}

	// if there is a tileset image, we need to create a cache of images.
	

	return ts, nil
}

type Tileset struct {
	Name string // the human readable name of this tileset. Only used for nice XML export. Unused otherwise
	// tileWidth and height are global constants
	// spacing and margin are xml serialization datails

	// the complete list of tiles. Tiles contain their own details
	tiles []TileInSet

	// the list of scripts

}

type WorldStateUpdate interface{}

type WorldAction interface {
	TransformWorld(win *interface{}) (WorldStateUpdate, error)
}

type TileInSet struct {
	Img image.Image // the sprite

	Animated      bool // should animation even be considered
	TimeInFrame   int  // ms to show this frame
	NextFrameTile int  // index (in tileset) of tile for next frame

	Solid bool // is it impassible (important optimization)

	// does passing into this cell activate an action? if so, go do it
	// solid tiles are activated by touch
	TouchAction WorldAction
	// non-solid tiles are activated by being inside (CG of player)
	// or by touching their hit-box (overlap with player hitbox) if non-nil
	Hitbox image.Rectangle // collision box of tile

	// the folloing two attributes are only relevant for serialization and not used in game execution
	// actions are either scripted (lua) or pre-compiled in a function pointer map
	TouchActionScriptIndex string // which index in the tileset script list generates the action for this Tile
	TouchLibaryAction      string // which pre-compiled library action does this tile use

	TickAction            WorldAction
	TickActionScriptIndex string // which index in the tileset script list generates the action for this Tile
	TickLibaryAction      string // which pre-compiled library action does this tile use
}
