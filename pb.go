package main

import (
	"fmt"

	"github.com/mrmasterplan/jumper/protobuf/level"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	p := level.PropertyString{
		Name:  "hello",
		Type:  0,
		Value: "world",
	}

	opt := protojson.MarshalOptions{
		Multiline:       true,
		Indent:          "  ",
		EmitUnpopulated: true,
	}
	b, err := opt.Marshal(&p)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	inb := []byte(`{
		"name":  "hello",
		"type":  "int",
		"value":  "world"
	  }`)

	uopt := protojson.UnmarshalOptions{
		AllowPartial: true,
	}
	if err := uopt.Unmarshal(inb, &p); err != nil {
		panic(err)
	}

	b, err = opt.Marshal(&p)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	
}
