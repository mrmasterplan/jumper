package TmxObjectGroup

import (
	"encoding/xml"

	"github.com/mrmasterplan/jumper/worldengine/tiled/mapxml/TmxLayerBase"
	"github.com/mrmasterplan/jumper/worldengine/tiled/mapxml/TmxObject"
	"github.com/mrmasterplan/jumper/worldengine/tiled/mapxml/TmxProperties"
)

type EmbedTmxObjectGroup struct {
	Objectgroup *TMXObjectGroup `xml:"objectgroup"`
}

type TMXObjectGroup struct {
	// <objectgroup>
	// -------------
	XMLName xml.Name `xml:"objectgroup"`

	TmxLayerBase.EmbedTmxLayerBase

	// -  **color:** The color used to display the objects in this group. (defaults
	//    to gray ("#a0a0a4"))
	// -  **draworder:** Whether the objects are drawn according to the order of
	//    appearance ("index") or sorted by their y-coordinate ("topdown").
	//    (defaults to "topdown")

	// The object group is in fact a map layer, and is hence called "object
	// layer" in Tiled.

	// Can contain at most one: :ref:`tmx-properties`
	TmxProperties.EmbedTmxProperties

	// Can contain any number: :ref:`tmx-object`
	Objects []TmxObject.TmxObject `xml:"object"`
}
