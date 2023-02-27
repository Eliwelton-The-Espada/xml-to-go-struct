package main

import (
	"bytes"
	"go/format"
	"log"
	"os"
	"strings"

	"github.com/miku/zek"
	splstr "github.com/rAndrade360/split-structs"
)

func main() {
	rootNode := new(zek.Node)
	rootNode.MaxExamples = 10

	var buf bytes.Buffer

	if _, err := rootNode.ReadFrom(strings.NewReader(xmlvar)); err != nil {
		log.Fatal("Err read: ", err.Error())
	}

	sw := zek.NewStructWriter(&buf)
	sw.WithComments = false
	sw.WithJSONTags = false
	sw.Strict = false
	sw.ExampleMaxChars = 25
	sw.Compact = true
	sw.UniqueExamples = false
	sw.OmitEmptyText = false
	sw.Banner = ""
	if err := sw.WriteNode(rootNode); err != nil {
		log.Fatal("Err read: ", err.Error())
	}

	f, err := os.Create("str.go")
	if err != nil {
		log.Fatal("Err read: ", err.Error())
	}

	d, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal("Err read: ", err.Error())
	}

	d = splstr.SplitStructs(d)

	f.Write(d)
	log.Printf("V: %#v", buf.String())
}
