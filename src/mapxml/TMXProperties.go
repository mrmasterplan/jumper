package mapxml

import (
	"encoding/xml"
)

type tmxProperties struct {
	Properties TMXProperties `xml:"properties"`
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
