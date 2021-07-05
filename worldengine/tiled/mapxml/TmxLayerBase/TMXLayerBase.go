package TmxLayerBase

type EmbedTmxLayerBase struct {
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

}