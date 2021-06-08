package mapxml

import (
	"encoding/xml"
	"fmt"
	"strings"
)

func decodeLayerStack(data string) (layers []TMXAnyLayer, properties *TMXProperties, error interface{}) {

	layerdecode := xml.NewDecoder(strings.NewReader(data))
	for {
		startToken, err := layerdecode.Token()
		if err != nil {
			return layers, properties, err
		}
		if startToken == nil {
			break
		}
		switch startToken.(type) {
		case xml.StartElement:
			startElement := startToken.(xml.StartElement)
			switch startToken.(xml.StartElement).Name.Local {
			case "layer":
				layer := TMXLayer{}

				if error = layerdecode.DecodeElement(&layer, &startElement); error != nil {
					return nil,nil,error
				}
				layers = append(layers, layer)

			case "group":
				layer := TMXGroup{}

				if error = layerdecode.DecodeElement(&layer, &startElement); error != nil {
					return nil,nil,error
				}
				var layerprops *TMXProperties
				layer.Layers, layerprops, error = decodeLayerStack(layer.LayerData)
				if error != nil {
					return nil,nil,error
				}

				layer.Properties = *layerprops

				layers = append(layers, layer)

			case "objectgroup":
				layer := TMXObjectGroup{}

				if error = layerdecode.DecodeElement(&layer, &startElement); error != nil {
					return nil,nil,error
				}
				layers = append(layers, layer)

			case "imagelayer":
				layer := TMXImageLayer{}

				if error = layerdecode.DecodeElement(&layer, &startElement); error != nil {
					return nil,nil,error
				}
				layers = append(layers, layer)
			case "properties":
				properties= &TMXProperties{}
				if error = layerdecode.DecodeElement(&properties, &startElement); error != nil {
					return nil,nil,error
				}
			case "editorsettings":
				// ignore
			default:
				panic(fmt.Sprintf("Unknown tag name in map: %v", startToken.(xml.StartElement).Name))
			}
		case xml.EndElement:
			panic(fmt.Sprintf("Unexpected EndElement map: %v", startToken.(xml.EndElement).Name))
		case xml.CharData:
			//panic(fmt.Sprintf("Unexpected CharData map: %v", string(startToken.(xml.CharData))))
		case xml.Comment:
			//panic(fmt.Sprintf("Unexpected Comment map: %v", string(startToken.(xml.Comment))))
		case xml.ProcInst:
			panic(fmt.Sprintf("Unexpected ProcInst map: %v", startToken.(xml.ProcInst)))
		case xml.Directive:
			panic(fmt.Sprintf("Unexpected Directive map: %v", startToken.(xml.Directive)))
		default:
			panic(`Unexpected token type. Malformed Map XML`)
		}
	}
	return layers, properties, nil
}

func ReadTiledMap(ba []byte) TiledMap {
	// m := fullMap{}
	var tiledMap TiledMap

	if err := xml.Unmarshal(ba, &tiledMap); err != nil {
		panic(err)
	}

	layers, properties, err := decodeLayerStack(tiledMap.LayerData)
	if err != nil {
		panic(err)
	}
	tiledMap.Layers = layers
	tiledMap.Properties = *properties

	return tiledMap
}
