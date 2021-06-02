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