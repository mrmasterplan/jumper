package mapxml

import (
	"encoding/xml"
)

type TiledMap struct {
	XMLName          xml.Name   `xml:"map"`
	Properties       Properties `xml:"properties"`
	Version          string     `xml:"version,attr"`
	TiledVersion     string     `xml:"tiledversion,attr"`
	Orientation      string     `xml:"orientation,attr"`
	RenderOrder      string     `xml:"renderorder,attr"`
	CompressionLevel string     `xml:"compressionlevel,attr"`
	Width            int        `xml:"width,attr"`
	Height           int        `xml:"height,attr"`
	TileHeight       int        `xml:"tileheight,attr"`
	TileWidth        int        `xml:"tilewidth,attr"`
	// backgroundcolor
	// nextlayerid // TODO
	// nextobjectid // TODO
	Infinite bool `xml:"infinite,attr"`

	Tilesets []Tileset `xml:"tileset"`
}

type Tileset struct {
	// 	firstgid: The first global tile ID of this tileset (this global ID maps to the first tile in this tileset).
	// source: If this tileset is stored in an external TSX (Tile Set XML) file, this attribute refers to that file. That TSX file has the same structure as the <tileset> element described here. (There is the firstgid attribute missing and this source attribute is also not there. These two attributes are kept in the TMX map, since they are map specific.)
	// name: The name of this tileset.
	// tilewidth: The (maximum) width of the tiles in this tileset.
	// tileheight: The (maximum) height of the tiles in this tileset.
	// spacing: The spacing in pixels between the tiles in this tileset (applies to the tileset image, defaults to 0)
	// margin: The margin around the tiles in this tileset (applies to the tileset image, defaults to 0)
	// tilecount: The number of tiles in this tileset (since 0.13)
	// columns: The number of tile columns in the tileset. For image collection tilesets it is editable and is used when displaying the tileset. (since 0.15)
	// objectalignment: Controls the alignment for tile objects. Valid values are unspecified, topleft, top, topright, left, center, right, bottomleft, bottom and bottomright. The default value is unspecified, for compatibility reasons. When unspecified, tile objects use bottomleft in orthogonal mode and bottom in isometric mode. (since 1.4)
	// If there are multiple <tileset> elements, they are in ascending order of their firstgid attribute. The first tileset always has a firstgid value of 1. Since Tiled 0.15, image collection tilesets do not necessarily number their tiles consecutively since gaps can occur when removing tiles.

	// Image collection tilesets have no <image> tag. Instead, each tile has an <image> tag.

	// Can contain at most one: :ref:`tmx-image`, :ref:`tmx-tileoffset`, :ref:`tmx-grid` (since 1.0), :ref:`tmx-properties`, :ref:`tmx-terraintypes`, :ref:`tmx-wangsets` (since 1.1), :ref:`tmx-tileset-transformations` (since 1.5)

	// Can contain any number: :ref:`tmx-tileset-tile`

	XMLName xml.Name `xml:"tileset"`

	FirstGID        int    `xml:"firstgid,attr"`
	Source          string `xml:"source,attr"`
	Name            string `xml:"name,attr"`
	TileHeight      int    `xml:"tileheight,attr"`
	TileWidth       int    `xml:"tilewidth,attr"`
	Spacing         int    `xml:"spacing,attr"`
	Margin          int    `xml:"margin,attr"`
	TileCount       int    `xml:"tilecount,attr"`
	Columns         int    `xml:"columns,attr"`
	ObjectAlignment string `xml:"objectalignment,attr"`
}

type Properties struct {
	XMLName xml.Name `xml:"properties"`

	Properties []Property `xml:"property"`
}

type Property struct {
	// 	name: The name of the property.
	// type: The type of the property. Can be string (default), int, float, bool, color, file or object (since 0.16, with color and file added in 0.17, and object added in 1.4).
	// value: The value of the property. (default string is "", default number is 0, default boolean is "false", default color is #00000000, default file is "." (the current file's parent directory))
	// Boolean properties have a value of either "true" or "false".

	// Color properties are stored in the format #AARRGGBB.

	// File properties are stored as paths relative from the location of the map file.

	// Object properties can reference any object on the same map and are stored as an integer (the ID of the referenced object, or 0 when no object is referenced). When used on objects in the Tile Collision Editor, they can only refer to other objects on the same tile.

	// When a string property contains newlines, the current version of Tiled will write out the value as characters contained inside the property element rather than as the value attribute. It is possible that a future version of the TMX format will switch to always saving property values inside the element rather than as an attribute.
	XMLName xml.Name `xml:"property"`

	Name  string `xml:"name,attr"`
	Type  string `xml:"type,attr"`
	Value string `xml:"value,attr"`
}
