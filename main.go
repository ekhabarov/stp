// Command main represents a Protobuf message generator out of Golang structures.
package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/ekhabarov/sts"
	"github.com/iancoleman/strcase"
)

var (
	filePath   = flag.String("f", "", "/path/to/file.go with structures")
	structName = flag.String("s", "", "Structure name to parse")
)

func main() {
	flag.Parse()

	// TODO(ekhabarov): Use field tag as destination field name.
	f, err := sts.Parse(*filePath, []string{"proto"})
	must(err)

	fields := f.Structs[*structName]

	s := make(sort.StringSlice, len(fields))
	for k, v := range fields {
		fld := strcase.ToSnake(k)
		s[v.Ord] = fmt.Sprintf("%s %s = %d;", v.Type, fld, v.Ord+1)
	}

	fmt.Printf("message %s {\n", *structName)
	for _, v := range s {
		fmt.Printf("\t%s\n", v)
	}

	fmt.Printf("}\n")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
