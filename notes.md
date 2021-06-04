to compile pb:

    protoc -I=protobuf\ --go_out=./ .\protobuf\*.proto


Dev tasks:
- ability to unmarshal tiled xml with tests
- protobuf defs roughly similar for saving level again
- conversion xml-pb-xml 
    - with errors
    - with tests
    - validation of xml should be here
        - animation consistency
        - special properties
            - type player
            - type collectible
            - collectibe function
            - etc.
    - assume on pb load that file is valid.
        - panic on parse error
        - check consistency (anmation etc.)

Links:
- https://tutorialedge.net/golang/parsing-xml-with-golang/
- https://golang.org/pkg/encoding/xml/#Unmarshal
- https://developers.google.com/protocol-buffers/docs/proto3
- https://github.com/mapeditor/tiled/blob/master/docs/reference/tmx-map-format.rst