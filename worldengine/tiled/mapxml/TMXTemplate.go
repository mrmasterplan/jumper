package mapxml

import (
	"encoding/xml"
)

type TMXTemplate struct {
	// <template>
	// ~~~~~~~~~~
	XMLName xml.Name `xml:"template"`

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