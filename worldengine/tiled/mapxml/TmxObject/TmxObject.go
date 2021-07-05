package TmxObject

import (
	"encoding/xml"

	"github.com/mrmasterplan/jumper/worldengine/tiled/mapxml/TmxProperties"
	"github.com/mrmasterplan/jumper/worldengine/tiled/mapxml/TmxUtils"
)

type TmxObject struct {
	// <object>
	// ~~~~~~~~
	XMLName xml.Name `xml:"object"`

	// -  **id:** Unique ID of the object. Each object that is placed on a map gets
	//    a unique id. Even if an object was deleted, no object gets the same
	//    ID. Can not be changed in Tiled. (since Tiled 0.11)
	Id int `xml:"id,attr"`

	// -  **name:** The name of the object. An arbitrary string. (defaults to "")
	Name string `xml:"name,attr"`
	// -  **type:** The type of the object. An arbitrary string. (defaults to "")
	Type string `xml:"type,attr"`
	// -  **x:** The x coordinate of the object in pixels. (defaults to 0)
	// -  **y:** The y coordinate of the object in pixels. (defaults to 0)
	TmxUtils.EmbedXY

	// -  **width:** The width of the object in pixels. (defaults to 0)
	// -  **height:** The height of the object in pixels. (defaults to 0)
	TmxUtils.EmbedWidthHeight

	// -  **rotation:** The rotation of the object in degrees clockwise around (x, y).
	//    (defaults to 0)
	Rotation float64 `xml:"rotation,attr"`
	// -  **gid:** A reference to a tile. (optional)
	Gid int `xml:"gid,attr"`
	// -  **visible:** Whether the object is shown (1) or hidden (0). (defaults to
	//    1)
	Visible int8 `xml:"visible,attr"`
	// -  **template:** A reference to a :ref:`template file <tmx-template-files>`. (optional)
	Template string `xml:"template,attr"`
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
		 
	TmxProperties.EmbedTmxProperties
	Ellipse *TmxEllipse
	Point *TmxPoint
	Polygon *TmxPolygon
	Polyline *TmxPolyLine
	Text *TmxText
}

type TmxEllipse struct {
	// <ellipse>
	// ~~~~~~~~~
	XMLName xml.Name `xml:"ellipse"`

	// Used to mark an object as an ellipse. The existing ``x``, ``y``,
	// ``width`` and ``height`` attributes are used to determine the size of
	// the ellipse.
}

type TmxPoint struct {
	// <point>
	// ~~~~~~~~~
	XMLName xml.Name `xml:"point"`

	// Used to mark an object as a point. The existing ``x`` and ``y`` attributes
	// are used to determine the position of the point.

}

type TmxPolygon struct {
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

type TmxPolyLine struct {
	// <polyline>
	// ~~~~~~~~~~
	XMLName xml.Name `xml:"polyline"`

	// -  **points:** A list of x,y coordinates in pixels.

	// A ``polyline`` follows the same placement definition as a ``polygon``
	// object.

}

type TmxText struct {
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
