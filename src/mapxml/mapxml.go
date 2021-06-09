package mapxml

// https://github.com/mapeditor/tiled/blob/master/docs/reference/tmx-map-format.rst
// ok, due to order and innerdata we need custom unmarshallers and marshallers
// TODO

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type TMXAnyLayer interface{}

type TiledMap struct {
	// 	<map>
	// -----
	XMLName xml.Name `xml:"map"`

	// -  **version:** The TMX format version. Was "1.0" so far, and will be
	//    incremented to match minor Tiled releases.
	Version string `xml:"version,attr"`

	// -  **tiledversion:** The Tiled version used to save the file (since Tiled
	//    1.0.1). May be a date (for snapshot builds). (optional)
	TiledVersion string `xml:"tiledversion,attr"`

	// -  **orientation:** Map orientation. Tiled supports "orthogonal",
	//    "isometric", "staggered" and "hexagonal" (since 0.11).
	Orientation string `xml:"orientation,attr"`

	// -  **renderorder:** The order in which tiles on tile layers are rendered.
	//    Valid values are ``right-down`` (the default), ``right-up``,
	//    ``left-down`` and ``left-up``. In all cases, the map is drawn
	//    row-by-row. (only supported for orthogonal maps at the moment)
	RenderOrder string `xml:"renderorder,attr"`

	// -  **compressionlevel:** The compression level to use for tile layer data
	//    (defaults to -1, which means to use the algorithm default).
	CompressionLevel string `xml:"compressionlevel,attr"`

	// -  **width:** The map width in tiles.
	Width int `xml:"width,attr"`

	// -  **height:** The map height in tiles.
	Height int `xml:"height,attr"`

	// -  **tilewidth:** The width of a tile.
	TileWidth int `xml:"tilewidth,attr"`

	// -  **tileheight:** The height of a tile.
	TileHeight int `xml:"tileheight,attr"`

	// -  **hexsidelength:** Only for hexagonal maps. Determines the width or
	//    height (depending on the staggered axis) of the tile's edge, in
	//    pixels.
	// -  **staggeraxis:** For staggered and hexagonal maps, determines which axis
	//    ("x" or "y") is staggered. (since 0.11)
	// -  **staggerindex:** For staggered and hexagonal maps, determines whether
	//    the "even" or "odd" indexes along the staggered axis are shifted.
	//    (since 0.11)
	// -  **backgroundcolor:** The background color of the map. (optional, may
	//    include alpha value since 0.15 in the form ``#AARRGGBB``. Defaults to
	//    fully transparent.)
	// -  **nextlayerid:** Stores the next available ID for new layers. This
	//    number is stored to prevent reuse of the same ID after layers have
	//    been removed. (since 1.2) (defaults to the highest layer id in the file
	//    + 1)
	NextLayerId int `xml:"nextlayerid,attr"`

	// -  **nextobjectid:** Stores the next available ID for new objects. This
	//    number is stored to prevent reuse of the same ID after objects have
	//    been removed. (since 0.11) (defaults to the highest object id in the file
	//    + 1)
	NextObjectId int `xml:"nextobjectid,attr"`

	// -  **infinite:** Whether this map is infinite. An infinite map has no fixed
	//    size and can grow in all directions. Its layer data is stored in chunks.
	//    (``0`` for false, ``1`` for true, defaults to 0)
	Infinite bool `xml:"infinite,attr"`

	// The ``tilewidth`` and ``tileheight`` properties determine the general
	// grid size of the map. The individual tiles may have different sizes.
	// Larger tiles will extend at the top and right (anchored to the bottom
	// left).

	// A map contains three different kinds of layers. Tile layers were once
	// the only type, and are simply called ``layer``, object layers have the
	// ``objectgroup`` tag and image layers use the ``imagelayer`` tag. The
	// order in which these layers appear is the order in which the layers are
	// rendered by Tiled.

	// The ``staggered`` orientation refers to an isometric map using staggered
	// axes.

	// Can contain at most one: :ref:`tmx-properties`
	Properties TMXProperties `xml:"properties"`

	// Can contain any number: :ref:`tmx-tileset`, :ref:`tmx-layer`,
	// :ref:`tmx-objectgroup`, :ref:`tmx-imagelayer`, :ref:`tmx-group` (since 1.0),
	// :ref:`tmx-editorsettings` (since 1.3)
	Tilesets []TMXTileset `xml:"tileset"`
	// LayerData string       `xml:",innerxml"`
	Layers []TMXAnyLayer
}

func (tm *TiledMap) getName() string {
	return tm.XMLName.Local
}

func UnmarshalLayersTilesetProperties(d *xml.Decoder, endTag string) (layers []TMXAnyLayer, tilesets []TMXTileset, properties TMXProperties, err error) {
	for {
		var token xml.Token
		token, err = d.Token()
		if err != nil {
			return nil, nil, properties, err
		}
		if token == nil {
			break
		}
		switch token.(type) {
		case xml.StartElement:
			start := token.(xml.StartElement)
			switch start.Name.Local {
			case "layer":
				layer := TMXLayer{}

				if err = d.DecodeElement(&layer, &start); err != nil {
					return nil, nil, properties, err
				}
				layers = append(layers, layer)

			case "group":
				layer := TMXGroup{}

				if err = d.DecodeElement(&layer, &start); err != nil {
					return nil, nil, properties, err
				}
				layers = append(layers, layer)

			case "objectgroup":
				layer := TMXObjectGroup{}
				if err = d.DecodeElement(&layer, &start); err != nil {
					return nil, nil, properties, err
				}
				layers = append(layers, layer)

			case "imagelayer":
				layer := TMXImageLayer{}

				if err = d.DecodeElement(&layer, &start); err != nil {
					return nil, nil, properties, err
				}
				layers = append(layers, layer)
			case "tileset":
				tileset := TMXTileset{}

				if err = d.DecodeElement(&tileset, &start); err != nil {
					return nil, nil, properties, err
				}
				tilesets = append(tilesets, tileset)
			case "properties":
				if err = d.DecodeElement(&properties, &start); err != nil {
					return nil, nil, properties, err
				}
			case "editorsettings":
				// ignore
			default:
				panic(fmt.Sprintf("Unknown tag name in map: %v", start.Name))
			}
		case xml.EndElement:
			if token.(xml.EndElement).Name.Local == endTag {
				// we are done parsing the map
				return layers, tilesets, properties, nil
			}
		case xml.CharData:
			//panic(fmt.Sprintf("Unexpected CharData map: %v", string(startToken.(xml.CharData))))
		case xml.Comment:
			//panic(fmt.Sprintf("Unexpected Comment map: %v", string(startToken.(xml.Comment))))
		case xml.ProcInst:
			panic(fmt.Errorf("unexpected ProcInst map: %v", token.(xml.ProcInst)))
		case xml.Directive:
			panic(fmt.Errorf("unexpected Directive map: %v", token.(xml.Directive)))
		default:
			panic(`Unexpected token type. Malformed Map XML`)
		}
	}
	return layers, tilesets, properties, fmt.Errorf(`no end tag found`)
}

func (tm *TiledMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {

	if start.Name.Local != "map" {
		return fmt.Errorf(`invalid xml name for TiledMap: "%v"`, start.Name)
	}
	tm.XMLName = start.Name

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		// -  **version:** The TMX format version. Was "1.0" so far, and will be
		//    incremented to match minor Tiled releases.
		case "version":
			tm.Version = attr.Value

			// -  **tiledversion:** The Tiled version used to save the file (since Tiled
			//    1.0.1). May be a date (for snapshot builds). (optional)
		case "tiledversion":
			tm.TiledVersion = attr.Value

			// -  **orientation:** Map orientation. Tiled supports "orthogonal",
			//    "isometric", "staggered" and "hexagonal" (since 0.11).
		case "orientation":
			tm.Orientation = attr.Value

			// -  **renderorder:** The order in which tiles on tile layers are rendered.
			//    Valid values are ``right-down`` (the default), ``right-up``,
			//    ``left-down`` and ``left-up``. In all cases, the map is drawn
			//    row-by-row. (only supported for orthogonal maps at the moment)
		case "renderorder":
			tm.RenderOrder = attr.Value

			// -  **compressionlevel:** The compression level to use for tile layer data
			//    (defaults to -1, which means to use the algorithm default).
		case "compressionlevel":
			tm.CompressionLevel = attr.Value

			// -  **width:** The map width in tiles.
		case "width":
			tm.Width, err = strconv.Atoi(attr.Value)
			if err != nil {
				return err
			}

			// -  **height:** The map height in tiles.
		case "height":
			tm.Height, err = strconv.Atoi(attr.Value)
			if err != nil {
				return err
			}

			// -  **tilewidth:** The width of a tile.
		case "tilewidth":
			tm.TileWidth, err = strconv.Atoi(attr.Value)
			if err != nil {
				return err
			}

			// -  **tileheight:** The height of a tile.
		case "tileheight":
			tm.TileHeight, err = strconv.Atoi(attr.Value)
			if err != nil {
				return err
			}

			// -  **hexsidelength:** Only for hexagonal maps. Determines the width or
			//    height (depending on the staggered axis) of the tile's edge, in
			//    pixels.
			// -  **staggeraxis:** For staggered and hexagonal maps, determines which axis
			//    ("x" or "y") is staggered. (since 0.11)
			// -  **staggerindex:** For staggered and hexagonal maps, determines whether
			//    the "even" or "odd" indexes along the staggered axis are shifted.
			//    (since 0.11)
			// -  **backgroundcolor:** The background color of the map. (optional, may
			//    include alpha value since 0.15 in the form ``#AARRGGBB``. Defaults to
			//    fully transparent.)
			// -  **nextlayerid:** Stores the next available ID for new layers. This
			//    number is stored to prevent reuse of the same ID after layers have
			//    been removed. (since 1.2) (defaults to the highest layer id in the file
			//    + 1)
		case "nextlayerid":
			tm.NextLayerId, err = strconv.Atoi(attr.Value)
			if err != nil {
				return err
			}

			// -  **nextobjectid:** Stores the next available ID for new objects. This
			//    number is stored to prevent reuse of the same ID after objects have
			//    been removed. (since 0.11) (defaults to the highest object id in the file
			//    + 1)
		case "nextobjectid":
			tm.NextObjectId, err = strconv.Atoi(attr.Value)
			if err != nil {
				return err
			}

			// -  **infinite:** Whether this map is infinite. An infinite map has no fixed
			//    size and can grow in all directions. Its layer data is stored in chunks.
			//    (``0`` for false, ``1`` for true, defaults to 0)
		case "infinite":
			tm.Infinite, err = strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
		default:
			// do nothing. some attributes are ignored.
		}
	}
	tm.Layers, tm.Tilesets, tm.Properties, err = UnmarshalLayersTilesetProperties(d, "map")
	if err != nil {
		return err
	}

	return nil
}

func (tm *TiledMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

type TMXEditorSettings struct {

	// 	<editorsettings>
	// ----------------
	XMLName xml.Name `xml:"editorsettings"`

	// This element contains various editor-specific settings, which are generally
	// not relevant when reading a map.

	// Can contain: :ref:`tmx-chunksize`, :ref:`tmx-export`

}

type TMXChunkSize struct {
	// 	<chunksize>
	// ~~~~~~~~~~~
	XMLName xml.Name `xml:"chunksize"`

	// -  **width:** The width of chunks used for infinite maps (default to 16).
	// -  **height:** The width of chunks used for infinite maps (default to 16).

}

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

type TMXWangSets struct {
	// 	<wangsets>
	// ~~~~~~~~~~
	XMLName xml.Name `xml:"wangsets"`

	// Contains the list of Wang sets defined for this tileset.

	// Can contain any number: :ref:`tmx-wangset`

}

type TMXWangset struct {
	// 	<wangset>
	// ^^^^^^^^^
	XMLName xml.Name `xml:"wangset"`

	// Defines a list of corner colors and a list of edge colors, and any
	// number of Wang tiles using these colors.

	// -  **name:** The name of the Wang set.
	// -  **tile:** The tile ID of the tile representing this Wang set.

	// Can contain at most one: :ref:`tmx-properties`

	// Can contain up to 255: :ref:`tmx-wangcolor` (since Tiled 1.5)

	// Can contain any number: :ref:`tmx-wangtile`

}

type TMXWangColor struct {
	// 	<wangcolor>
	// '''''''''''
	XMLName xml.Name `xml:"wangcolor"`

	// A color that can be used to define the corner and/or edge of a Wang tile.

	// -  **name:** The name of this color.
	// -  **color:** The color in ``#RRGGBB`` format (example: ``#c17d11``).
	// -  **tile:** The tile ID of the tile representing this color.
	// -  **probability:** The relative probability that this color is chosen
	//    over others in case of multiple options. (defaults to 0)

	// Can contain at most one: :ref:`tmx-properties`

}

type TMXWangTile struct {
	// 	<wangtile>
	// ''''''''''
	XMLName xml.Name `xml:"wangtile"`

	// Defines a Wang tile, by referring to a tile in the tileset and
	// associating it with a certain Wang ID.

	// -  **tileid:** The tile ID.
	// -  **wangid:** "The Wang ID, given by a comma-separated list of indexes
	//    (starting from 1, because 0 means _unset_) referring to the Wang colors in
	//    the Wang set in the following order: top, top right, right, bottom right,
	//    bottom, bottom left, left, top left (since Tiled 1.5). Before Tiled 1.5, the
	//    Wang ID was saved as a 32-bit unsigned integer stored in the format
	//    ``0xCECECECE`` (where each C is a corner color and each E is an edge color,
	//    in reverse order)."
	// -  *hflip:* Whether the tile is flipped horizontally (removed in Tiled 1.5).
	// -  *vflip:* Whether the tile is flipped vertically (removed in Tiled 1.5).
	// -  *dflip:* Whether the tile is flipped on its diagonal (removed in Tiled 1.5).

}

type TMXLayer struct {
	// 	<layer>
	// -------
	XMLName xml.Name `xml:"layer"`

	// All :ref:`tmx-tileset` tags shall occur before the first :ref:`tmx-layer` tag
	// so that parsers may rely on having the tilesets before needing to resolve
	// tiles.

	// -  **id:** Unique ID of the layer. Each layer that added to a map gets
	//    a unique id. Even if a layer is deleted, no layer ever gets the same
	//    ID. Can not be changed in Tiled. (since Tiled 1.2)
	Id int `xml:"id,attr"`
	// -  **name:** The name of the layer. (defaults to "")
	Name string `xml:"name,attr"`
	// -  *x:* The x coordinate of the layer in tiles. Defaults to 0 and can not be changed in Tiled.
	X int `xml:"x,attr"`
	// -  *y:* The y coordinate of the layer in tiles. Defaults to 0 and can not be changed in Tiled.
	Y int `xml:"y,attr"`
	// -  **width:** The width of the layer in tiles. Always the same as the map width for fixed-size maps.
	Width int `xml:"width,attr"`
	// -  **height:** The height of the layer in tiles. Always the same as the map height for fixed-size maps.
	Height int `xml:"height,attr"`
	// -  **opacity:** The opacity of the layer as a value from 0 to 1. Defaults to 1.
	Opacity float32 `xml:"opacity,attr"`
	// -  **visible:** Whether the layer is shown (1) or hidden (0). Defaults to 1.
	Visible bool `xml:"visible,attr"`
	// -  **tintcolor:** A :ref:`tint color <tint-color>` that is multiplied with any tiles drawn by this layer in ``#AARRGGBB`` or ``#RRGGBB`` format (optional).
	TintColor string `xml:"tintcolor,attr"`
	// -  **offsetx:** Horizontal offset for this layer in pixels. Defaults to 0.
	//    (since 0.14)
	OffsetX int `xml:"offsetx,attr"`
	// -  **offsety:** Vertical offset for this layer in pixels. Defaults to 0.
	//    (since 0.14)
	OffsetY int `xml:"offsety,attr"`
	// -  **parallaxx:** Horizontal :ref:`parallax factor <parallax-factor>` for this layer. Defaults to 1. (since 1.5)
	ParallaxX float32 `xml:"parallaxx,attr"`
	// -  **parallaxy:** Vertical :ref:`parallax factor <parallax-factor>` for this layer. Defaults to 1. (since 1.5)
	ParallaxY float32 `xml:"parallaxy,attr"`

	// Can contain at most one: :ref:`tmx-properties`, :ref:`tmx-data`
	Properties TMXProperties `xml:"properties"`
	Data       TMXData       `xml:"data"`
}

func (tm *TMXLayer) getName() string {
	return tm.XMLName.Local
}

type TMXData struct {
	// 	<data>
	// ~~~~~~
	XMLName xml.Name `xml:"data"`

	// -  **encoding:** The encoding used to encode the tile layer data. When used,
	//    it can be "base64" and "csv" at the moment. (optional)
	// -  **compression:** The compression used to compress the tile layer data.
	//    Tiled supports "gzip", "zlib" and (as a compile-time option since Tiled 1.3)
	//    "zstd".

	// When no encoding or compression is given, the tiles are stored as
	// individual XML ``tile`` elements. Next to that, the easiest format to
	// parse is the "csv" (comma separated values) format.

	// The base64-encoded and optionally compressed layer data is somewhat more
	// complicated to parse. First you need to base64-decode it, then you may
	// need to decompress it. Now you have an array of bytes, which should be
	// interpreted as an array of unsigned 32-bit integers using little-endian
	// byte ordering.

	// Whatever format you choose for your layer data, you will always end up
	// with so called "global tile IDs" (gids). They are global, since they may
	// refer to a tile from any of the tilesets used by the map. In order to
	// find out from which tileset the tile is you need to find the tileset
	// with the highest ``firstgid`` that is still lower or equal than the gid.
	// The tilesets are always stored with increasing ``firstgid``\ s.

	// Can contain any number: :ref:`tmx-tilelayer-tile`, :ref:`tmx-chunk`

}

type TMXChunk struct {
	// 	<chunk>
	// ~~~~~~~
	XMLName xml.Name `xml:"chunk"`

	// -  **x:** The x coordinate of the chunk in tiles.
	X int `xml:"x,attr"`

	// -  **y:** The y coordinate of the chunk in tiles.
	Y int `xml:"y,attr"`

	// -  **width:** The width of the chunk in tiles.
	Width int `xml:"width,attr"`

	// -  **height:** The height of the chunk in tiles.
	Height int `xml:"height,attr"`

	// This is currently added only for infinite maps. The contents of a chunk
	// element is same as that of the ``data`` element, except it stores the
	// data of the area specified in the attributes.

	// Can contain any number: :ref:`tmx-tilelayer-tile`
	Tiles []TMXTileLayerTile `xml:"tile"`

	Data string `xml:",innerxml"`
}

type TMXTileLayerTile struct {
	// 	<tile>
	// ~~~~~~
	XMLName xml.Name `xml:"tile"`

	// -  **gid:** The global tile ID (default: 0).
	Gid int `xml:"gid,attr"`

	// Not to be confused with the ``tile`` element inside a ``tileset``, this
	// element defines the value of a single tile on a tile layer. This is
	// however the most inefficient way of storing the tile layer data, and
	// should generally be avoided.

}

type TMXObjectGroup struct {
	// <objectgroup>
	// -------------
	XMLName xml.Name `xml:"objectgroup"`

	// -  **id:** Unique ID of the layer. Each layer that added to a map gets
	//    a unique id. Even if a layer is deleted, no layer ever gets the same
	//    ID. Can not be changed in Tiled. (since Tiled 1.2)
	// -  **name:** The name of the object group. (defaults to "")
	// -  **color:** The color used to display the objects in this group. (defaults
	//    to gray ("#a0a0a4"))
	// -  *x:* The x coordinate of the object group in tiles. Defaults to 0 and
	//    can no longer be changed in Tiled.
	// -  *y:* The y coordinate of the object group in tiles. Defaults to 0 and
	//    can no longer be changed in Tiled.
	// -  *width:* The width of the object group in tiles. Meaningless.
	// -  *height:* The height of the object group in tiles. Meaningless.
	// -  **opacity:** The opacity of the layer as a value from 0 to 1. (defaults to
	//    1)
	// -  **visible:** Whether the layer is shown (1) or hidden (0). (defaults to 1)
	// -  **tintcolor:** A color that is multiplied with any tile objects drawn by this layer, in ``#AARRGGBB`` or ``#RRGGBB`` format (optional).
	// -  **offsetx:** Horizontal offset for this object group in pixels. (defaults
	//    to 0) (since 0.14)
	// -  **offsety:** Vertical offset for this object group in pixels. (defaults
	//    to 0) (since 0.14)
	// -  **draworder:** Whether the objects are drawn according to the order of
	//    appearance ("index") or sorted by their y-coordinate ("topdown").
	//    (defaults to "topdown")

	// The object group is in fact a map layer, and is hence called "object
	// layer" in Tiled.

	// Can contain at most one: :ref:`tmx-properties`

	// Can contain any number: :ref:`tmx-object`

}

func (tm *TMXObjectGroup) getName() string {
	return tm.XMLName.Local
}

type TMXObject struct {
	// <object>
	// ~~~~~~~~
	XMLName xml.Name `xml:"object"`

	// -  **id:** Unique ID of the object. Each object that is placed on a map gets
	//    a unique id. Even if an object was deleted, no object gets the same
	//    ID. Can not be changed in Tiled. (since Tiled 0.11)
	// -  **name:** The name of the object. An arbitrary string. (defaults to "")
	// -  **type:** The type of the object. An arbitrary string. (defaults to "")
	// -  **x:** The x coordinate of the object in pixels. (defaults to 0)
	// -  **y:** The y coordinate of the object in pixels. (defaults to 0)
	// -  **width:** The width of the object in pixels. (defaults to 0)
	// -  **height:** The height of the object in pixels. (defaults to 0)
	// -  **rotation:** The rotation of the object in degrees clockwise around (x, y).
	//    (defaults to 0)
	// -  **gid:** A reference to a tile. (optional)
	// -  **visible:** Whether the object is shown (1) or hidden (0). (defaults to
	//    1)
	// -  **template:** A reference to a :ref:`template file <tmx-template-files>`. (optional)

	// While tile layers are very suitable for anything repetitive aligned to
	// the tile grid, sometimes you want to annotate your map with other
	// information, not necessarily aligned to the grid. Hence the objects have
	// their coordinates and size in pixels, but you can still easily align
	// that to the grid when you want to.

	// You generally use objects to add custom information to your tile map,
	// such as spawn points, warps, exits, etc.

	// When the object has a ``gid`` set, then it is represented by the image
	// of the tile with that global ID. The image alignment currently depends
	// on the map orientation. In orthogonal orientation it's aligned to the
	// bottom-left while in isometric it's aligned to the bottom-center. The
	// image will rotate around the bottom-left or bottom-center, respectively.

	// When the object has a ``template`` set, it will borrow all the
	// properties from the specified template, properties saved with the object
	// will have higher priority, i.e. they will override the template
	// properties.

	// Can contain at most one: :ref:`tmx-properties`, :ref:`tmx-ellipse` (since
	// 0.9), :ref:`tmx-point` (since 1.1), :ref:`tmx-polygon`, :ref:`tmx-polyline`,
	// :ref:`tmx-text` (since 1.0)

}

type TMXEllipse struct {
	// <ellipse>
	// ~~~~~~~~~
	XMLName xml.Name `xml:"ellipse"`

	// Used to mark an object as an ellipse. The existing ``x``, ``y``,
	// ``width`` and ``height`` attributes are used to determine the size of
	// the ellipse.
}

type TMXPoint struct {
	// <point>
	// ~~~~~~~~~
	XMLName xml.Name `xml:"point"`

	// Used to mark an object as a point. The existing ``x`` and ``y`` attributes
	// are used to determine the position of the point.

}

type TMXPolygon struct {
	// <polygon>
	// ~~~~~~~~~
	XMLName xml.Name `xml:"polygon"`

	// -  **points:** A list of x,y coordinates in pixels.

	// Each ``polygon`` object is made up of a space-delimited list of x,y
	// coordinates. The origin for these coordinates is the location of the
	// parent ``object``. By default, the first point is created as 0,0
	// denoting that the point will originate exactly where the ``object`` is
	// placed.

}

type TMXPolyLine struct {
	// <polyline>
	// ~~~~~~~~~~
	XMLName xml.Name `xml:"polyline"`

	// -  **points:** A list of x,y coordinates in pixels.

	// A ``polyline`` follows the same placement definition as a ``polygon``
	// object.

}

type TMXText struct {
	// <text>
	// ~~~~~~
	XMLName xml.Name `xml:"text"`

	// -  **fontfamily:** The font family used (defaults to "sans-serif")
	// -  **pixelsize:** The size of the font in pixels (not using points,
	//    because other sizes in the TMX format are also using pixels)
	//    (defaults to 16)
	// -  **wrap:** Whether word wrapping is enabled (1) or disabled (0).
	//    (defaults to 0)
	// -  **color:** Color of the text in ``#AARRGGBB`` or ``#RRGGBB`` format
	//    (defaults to #000000)
	// -  **bold:** Whether the font is bold (1) or not (0). (defaults to 0)
	// -  **italic:** Whether the font is italic (1) or not (0). (defaults to 0)
	// -  **underline:** Whether a line should be drawn below the text (1) or
	//    not (0). (defaults to 0)
	// -  **strikeout:** Whether a line should be drawn through the text (1) or
	//    not (0). (defaults to 0)
	// -  **kerning:** Whether kerning should be used while rendering the text
	//    (1) or not (0). (defaults to 1)
	// -  **halign:** Horizontal alignment of the text within the object
	//    (``left``, ``center``, ``right`` or ``justify``, defaults to ``left``)
	//    (since Tiled 1.2.1)
	// -  **valign:** Vertical alignment of the text within the object (``top``
	//    , ``center`` or ``bottom``, defaults to ``top``)

	// Used to mark an object as a text object. Contains the actual text as
	// character data.

	// For alignment purposes, the bottom of the text is the descender height of
	// the font, and the top of the text is the ascender height of the font. For
	// example, ``bottom`` alignment of the word "cat" will leave some space below
	// the text, even though it is unused for this word with most fonts. Similarly,
	// ``top`` alignment of the word "cat" will leave some space above the "t" with
	// most fonts, because this space is used for diacritics.

	// If the text is larger than the object's bounds, it is clipped to the bounds
	// of the object.
}

type TMXImageLayer struct {
	// <imagelayer>
	// ------------
	XMLName xml.Name `xml:"imagelayer"`

	// -  **id:** Unique ID of the layer. Each layer that added to a map gets
	// a unique id. Even if a layer is deleted, no layer ever gets the same
	// ID. Can not be changed in Tiled. (since Tiled 1.2)
	// -  **name:** The name of the image layer. (defaults to "")
	// -  **offsetx:** Horizontal offset of the image layer in pixels. (defaults to
	// 0) (since 0.15)
	// -  **offsety:** Vertical offset of the image layer in pixels. (defaults to
	// 0) (since 0.15)
	// -  *x:* The x position of the image layer in pixels. (defaults to 0, deprecated
	// since 0.15)
	// -  *y:* The y position of the image layer in pixels. (defaults to 0, deprecated
	// since 0.15)
	// -  **opacity:** The opacity of the layer as a value from 0 to 1. (defaults to
	// 1)
	// -  **visible:** Whether the layer is shown (1) or hidden (0). (defaults to 1)
	// -  **tintcolor:** A color that is multiplied with the image drawn by this layer in ``#AARRGGBB`` or ``#RRGGBB`` format (optional).

	// A layer consisting of a single image.

	// Can contain at most one: :ref:`tmx-properties`, :ref:`tmx-image`

}

func (tm *TMXImageLayer) getName() string {
	return tm.XMLName.Local
}

type TMXGroup struct {
	// <group>
	// -------
	XMLName xml.Name

	// -  **id:** Unique ID of the layer. Each layer that added to a map gets
	//    a unique id. Even if a layer is deleted, no layer ever gets the same
	//    ID. Can not be changed in Tiled. (since Tiled 1.2)
	Id int
	// -  **name:** The name of the group layer. (defaults to "")
	Name string
	// -  **offsetx:** Horizontal offset of the group layer in pixels. (defaults to
	//    0)
	OffsetX int
	// -  **offsety:** Vertical offset of the group layer in pixels. (defaults to
	//    0)
	OffsetY int
	// -  **opacity:** The opacity of the layer as a value from 0 to 1. (defaults to
	//    1)
	Opacity float64
	// -  **visible:** Whether the layer is shown (1) or hidden (0). (defaults to 1)
	Visible bool
	// -  **tintcolor:** A color that is multiplied with any graphics drawn by any child layers, in ``#AARRGGBB`` or ``#RRGGBB`` format (optional).
	TintColor string
	// A group layer, used to organize the layers of the map in a hierarchy.
	// Its attributes ``offsetx``, ``offsety``, ``opacity``, ``visible`` and
	// ``tintcolor`` recursively affect child layers.

	// Can contain at most one: :ref:`tmx-properties`
	Properties TMXProperties

	// Can contain any number: :ref:`tmx-layer`,
	// :ref:`tmx-objectgroup`, :ref:`tmx-imagelayer`, :ref:`tmx-group`
	Layers []TMXAnyLayer
}

func (tm *TMXGroup) getName() string {
	return tm.XMLName.Local
}

func (gr *TMXGroup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {

	// <group>
	// -------
	if start.Name.Local != "group" {
		return fmt.Errorf(`invalid xml name for TiledMap: "%v"`, start.Name)
	}
	gr.XMLName = start.Name

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		// -  **id:** Unique ID of the layer. Each layer that added to a map gets
		//    a unique id. Even if a layer is deleted, no layer ever gets the same
		//    ID. Can not be changed in Tiled. (since Tiled 1.2)
		case "id":
			gr.Id, err = strconv.Atoi(attr.Value)
			if err != nil {
				return err
			}

			// -  **name:** The name of the group layer. (defaults to "")
		case "name":
			gr.Name = attr.Value
			// -  **offsetx:** Horizontal offset of the group layer in pixels. (defaults to
			//    0)
		case "offsetx":
			gr.OffsetX, err = strconv.Atoi(attr.Value)
			if err != nil {
				return err
			}

			// -  **offsety:** Vertical offset of the group layer in pixels. (defaults to
			//    0)
		case "offsety":
			gr.OffsetY, err = strconv.Atoi(attr.Value)
			if err != nil {
				return err
			}

			// -  **opacity:** The opacity of the layer as a value from 0 to 1. (defaults to
			//    1)
		case "opacity":
			gr.Opacity, err = strconv.ParseFloat(attr.Value, 32)
			if err != nil {
				return err
			}

			// -  **visible:** Whether the layer is shown (1) or hidden (0). (defaults to 1)
		case "visible":
			gr.Visible, err = strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}

			// -  **tintcolor:** A color that is multiplied with any graphics drawn by any child layers, in ``#AARRGGBB`` or ``#RRGGBB`` format (optional).
		case "tintcolor":
			gr.TintColor = attr.Value

			// A group layer, used to organize the layers of the map in a hierarchy.
			// Its attributes ``offsetx``, ``offsety``, ``opacity``, ``visible`` and
			// ``tintcolor`` recursively affect child layers.
		default:
			// do nothing. some attributes are ignored.
		}
	}
	var tilesets []TMXTileset
	gr.Layers, tilesets, gr.Properties, err = UnmarshalLayersTilesetProperties(d, "map")
	if err != nil {
		return err
	}
	if tilesets != nil {
		return fmt.Errorf(`Tileset tag found in group layer.`)
	}

	return nil
}

func (tm *TMXGroup) MarshalXML(d *xml.Encoder, start xml.StartElement) (err error) {
	return nil
}

type TMXProperties struct {
	// <properties>
	// ------------
	XMLName xml.Name `xml:"properties"`

	// Wraps any number of custom properties. Can be used as a child of the
	// ``map``, ``tileset``, ``tile`` (when part of a ``tileset``),
	// ``terrain``, ``wangset``, ``wangcolor``, ``layer``, ``objectgroup``,
	// ``object``, ``imagelayer`` and ``group`` elements.

	// Can contain any number: :ref:`tmx-property`
	Properties []TMXProperty `xml:"property"`
}

func (tm *TMXProperties) getName() string {
	return tm.XMLName.Local
}

type TMXProperty struct {
	// <property>
	// ~~~~~~~~~~
	XMLName xml.Name `xml:"property"`

	// -  **name:** The name of the property.
	Name string `xml:"name,attr"`

	// -  **type:** The type of the property. Can be ``string`` (default), ``int``,
	//    ``float``, ``bool``, ``color``, ``file`` or ``object`` (since 0.16, with
	//    ``color`` and ``file`` added in 0.17, and ``object`` added in 1.4).
	Type string `xml:"type,attr"`

	// -  **value:** The value of the property. (default string is "", default
	//    number is 0, default boolean is "false", default color is #00000000, default
	//    file is "." (the current file's parent directory))
	Value      string `xml:"value,attr"`
	InnerValue string `xml:",innerxml"`

	// Boolean properties have a value of either "true" or "false".

	// Color properties are stored in the format ``#AARRGGBB``.

	// File properties are stored as paths relative from the location of the
	// map file.

	// Object properties can reference any object on the same map and are stored as an
	// integer (the ID of the referenced object, or 0 when no object is referenced).
	// When used on objects in the Tile Collision Editor, they can only refer to
	// other objects on the same tile.

	// When a string property contains newlines, the current version of Tiled
	// will write out the value as characters contained inside the ``property``
	// element rather than as the ``value`` attribute. It is possible that a
	// future version of the TMX format will switch to always saving property
	// values inside the element rather than as an attribute.

}

type TMXTemplate struct {
	// <template>
	// ~~~~~~~~~~

	// The template root element contains the saved :ref:`map object <tmx-object>`
	// and a :ref:`tileset <tmx-tileset>` element that points to an external
	// tileset, if the object is a tile object.

	// Example of a template file:

	//    .. code:: xml

	// 	<?xml version="1.0" encoding="UTF-8"?>
	// 	<template>
	// 	 <tileset firstgid="1" source="desert.tsx"/>
	// 	 <object name="cactus" gid="31" width="81" height="101"/>
	// 	</template>

	// Can contain at most one: :ref:`tmx-tileset`, :ref:`tmx-object`
}
