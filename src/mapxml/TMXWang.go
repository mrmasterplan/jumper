package mapxml

import (
	"encoding/xml"
)


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
