//go:build ignore
// +build ignore

package ent

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	if err := entc.Generate("./schema", &gen.Config{}); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
