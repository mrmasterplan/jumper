package mapxml

import (
	"encoding/xml"
)


type TMXTileset struct {
	// 	<tileset>
	// ---------
	XMLName xml.Name `xml:"tileset"`

	// -  **firstgid:** The first global tile ID of this tileset (this global ID
	//    maps to the first tile in this tileset).
	FirstGID int `xml:"firstgid,attr"`
	// -  **source:** If this tileset is stored in an external TSX (Tile Set XML)
	//    file, this attribute refers to that file. That TSX file has the same
	//    structure as the ``<tileset>`` element described here. (There is the
	//    firstgid attribute missing and this source attribute is also not
	//    there. These two attributes are kept in the TMX map, since they are
	//    map specific.)
	Source string `xml:"source,attr"`
	// -  **name:** The name of this tileset.
	Name string `xml:"name,attr"`
	// -  **tilewidth:** The (maximum) width of the tiles in this tileset.
	TileWidth int `xml:"tilewidth,attr"`
	// -  **tileheight:** The (maximum) height of the tiles in this tileset.
	TileHeight int `xml:"tileheight,attr"`
	// -  **spacing:** The spacing in pixels between the tiles in this tileset
	//    (applies to the tileset image, defaults to 0)
	Spacing int `xml:"spacing,attr"`
	// -  **margin:** The margin around the tiles in this tileset (applies to the
	//    tileset image, defaults to 0)
	Margin int `xml:"margin,attr"`
	// -  **tilecount:** The number of tiles in this tileset (since 0.13)
	TileCount int `xml:"tilecount,attr"`
	// -  **columns:** The number of tile columns in the tileset. For image
	//    collection tilesets it is editable and is used when displaying the
	//    tileset. (since 0.15)
	Columns int `xml:"columns,attr"`
	// -  **objectalignment:** Controls the alignment for tile objects.
	//    Valid values are ``unspecified``, ``topleft``, ``top``, ``topright``,
	//    ``left``, ``center``, ``right``, ``bottomleft``, ``bottom`` and
	//    ``bottomright``. The default value is ``unspecified``, for compatibility
	//    reasons. When unspecified, tile objects use ``bottomleft`` in orthogonal mode
	//    and ``bottom`` in isometric mode. (since 1.4)
	ObjectAlignment string `xml:"objectalignment,attr"`

	// If there are multiple ``<tileset>`` elements, they are in ascending
	// order of their ``firstgid`` attribute. The first tileset always has a
	// ``firstgid`` value of 1. Since Tiled 0.15, image collection tilesets do
	// not necessarily number their tiles consecutively since gaps can occur
	// when removing tiles.

	// Image collection tilesets have no ``<image>`` tag. Instead, each tile has
	// an ``<image>`` tag.

	// Can contain at most one: :ref:`tmx-image`, :ref:`tmx-tileoffset`,
	// :ref:`tmx-grid` (since 1.0), :ref:`tmx-properties`, :ref:`tmx-terraintypes`,
	// :ref:`tmx-wangsets` (since 1.1), :ref:`tmx-tileset-transformations` (since 1.5)
	Image           TMXImage           `xml:"image"`
	TileOffset      TMXTileOffset      `xml:"tileoffset"`
	Grid            TMXGrid            `xml:"grid"`
	Properties      TMXProperties      `xml:"properties"`
	TerrainType     TMXTerrainTypes    `xml:"terraintypes"`
	Wangsets        TMXWangSets        `xml:"wangsets"`
	Transformations TMXTransformations `xml:"transformations"`

	// Can contain any number: :ref:`tmx-tileset-tile`
	Tiles []TMXTileInTileset `xml:"tile"`
}

func (tm *TMXTileset) getName() string {
	return tm.XMLName.Local
}

type TMXTileOffset struct {
	// 	<tileoffset>
	// ~~~~~~~~~~~~
	XMLName xml.Name `xml:"tileoffset"`

	// -  **x:** Horizontal offset in pixels. (defaults to 0)
	X int `xml:"x,attr"`
	// -  **y:** Vertical offset in pixels (positive is down, defaults to 0)
	Y int `xml:"y,attr"`

	// This element is used to specify an offset in pixels, to be applied when
	// drawing a tile from the related tileset. When not present, no offset is
	// applied.

}

type TMXGrid struct {
	// 	<grid>
	// ~~~~~~
	XMLName xml.Name `xml:"grid"`

	// -  **orientation:** Orientation of the grid for the tiles in this
	//    tileset (``orthogonal`` or ``isometric``, defaults to ``orthogonal``)
	// -  **width:** Width of a grid cell
	// -  **height:** Height of a grid cell

	// This element is only used in case of isometric orientation, and
	// determines how tile overlays for terrain and collision information are
	// rendered.

}

type TMXImage struct {
	// 	<image>
	// ~~~~~~~
	XMLName xml.Name `xml:"image"`

	// -  **format:** Used for embedded images, in combination with a ``data``
	//    child element. Valid values are file extensions like ``png``,
	//    ``gif``, ``jpg``, ``bmp``, etc.
	// -  *id:* Used by some versions of Tiled Java. Deprecated and unsupported.
	// -  **source:** The reference to the tileset image file (Tiled supports most
	//    common image formats). Only used if the image is not embedded.
	// -  **trans:** Defines a specific color that is treated as transparent
	//    (example value: "#FF00FF" for magenta). Including the "#" is optional
	//    and Tiled leaves it out for compatibility reasons. (optional)
	// -  **width:** The image width in pixels (optional, used for tile index
	//    correction when the image changes)
	// -  **height:** The image height in pixels (optional)

	// Note that it is not currently possible to use Tiled to create maps with
	// embedded image data, even though the TMX format supports this. It is
	// possible to create such maps using ``libtiled`` (Qt/C++) or
	// `tmxlib <https://pypi.python.org/pypi/tmxlib>`__ (Python).

	// Can contain at most one: :ref:`tmx-data`

}

type TMXTerrainTypes struct {
	// 	<terraintypes>
	// ~~~~~~~~~~~~~~
	XMLName xml.Name `xml:"terraintypes"`

	// This element defines an array of terrain types, which can be referenced
	// from the ``terrain`` attribute of the ``tile`` element.

	// Can contain any number: :ref:`tmx-terrain`

}

type TMXTerrain struct {
	// 	<terrain>
	// ^^^^^^^^^
	XMLName xml.Name `xml:"terrain"`

	// -  **name:** The name of the terrain type.
	// -  **tile:** The local tile-id of the tile that represents the terrain
	//    visually.

	// Can contain at most one: :ref:`tmx-properties`

}

type TMXTransformations struct {
	// 	<transformations>
	// ~~~~~~~~~~~~~~~~~
	XMLName xml.Name `xml:"transformations"`

	// This element is used to describe which transformations can be applied to the
	// tiles (e.g. to extend a Wang set by transforming existing tiles).

	// - **hflip:** Whether the tiles in this set can be flipped horizontally (default 0)
	// - **vflip:** Whether the tiles in this set can be flipped vertically (default 0)
	// - **rotate:** Whether the tiles in this set can be rotated in 90 degree increments (default 0)
	// - **preferuntransformed:** Whether untransformed tiles remain preferred, otherwise
	//   transformed tiles are used to produce more variations (default 0)

}

type TMXTileInTileset struct {
	// 	<tile>
	// ~~~~~~
	XMLName xml.Name `xml:"tile"`

	// -  **id:** The local tile ID within its tileset.
	Id int `xml:"id,attr"`
	// -  **type:** The type of the tile. Refers to an object type and is used
	//    by tile objects. (optional) (since 1.0)
	Type string `xml:"type,attr"`
	// -  **terrain:** Defines the terrain type of each corner of the tile,
	//    given as comma-separated indexes in the terrain types array in the
	//    order top-left, top-right, bottom-left, bottom-right. Leaving out a
	//    value means that corner has no terrain. (optional)
	Terrain string `xml:"terrain,attr"`
	// -  **probability:** A percentage indicating the probability that this
	//    tile is chosen when it competes with others while editing with the
	//    terrain tool. (defaults to 0)
	Probability string `xml:"probablility,attr"`

	// Can contain at most one: :ref:`tmx-properties`, :ref:`tmx-image` (since
	// 0.9), :ref:`tmx-objectgroup`, :ref:`tmx-animation`
	Properties  TMXProperties  `xml:"properties"`
	Image       TMXImage       `xml:"image"`
	Objectgroup TMXObjectGroup `xml:"objectgroup"`
	Animation   TMXAnimation   `xml:"animation"`
}

type TMXAnimation struct {
	// 	<animation>
	// ^^^^^^^^^^^
	XMLName xml.Name `xml:"animation"`

	// Contains a list of animation frames.

	// Each tile can have exactly one animation associated with it. In the
	// future, there could be support for multiple named animations on a tile.

	// Can contain any number: :ref:`tmx-frame`
	Frames []TMXFrame `xml:"frame"`
}

type TMXFrame struct {
	// 	<frame>
	// '''''''
	XMLName xml.Name `xml:"frame"`

	// -  **tileid:** The local ID of a tile within the parent
	//    :ref:`tmx-tileset`.
	TileId int `xml:"tileid,attr"`
	// -  **duration:** How long (in milliseconds) this frame should be displayed
	//    before advancing to the next frame.
	Duration int `xml:"duration,attr"`
}
