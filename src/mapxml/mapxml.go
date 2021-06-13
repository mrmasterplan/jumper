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

type tmxTiledMapAttrs struct {
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

	// -  **height:** The map height in tiles.
	tmxWidthHeight

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
}

type TiledMap struct {
	// 	<map>
	// -----
	XMLName xml.Name `xml:"map"`

	tmxTiledMapAttrs

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
	// -  **y:** The y coordinate of the chunk in tiles.
	tmxXY

	// -  **width:** The width of the chunk in tiles.
	// -  **height:** The height of the chunk in tiles.
	tmxWidthHeight

	// This is currently added only for infinite maps. The contents of a chunk
	// element is same as that of the ``data`` element, except it stores the
	// data of the area specified in the attributes.

	// Can contain any number: :ref:`tmx-tilelayer-tile`
	Tiles []TMXTileLayerTile `xml:"tile"`

	Data string `xml:",chardata"`
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
