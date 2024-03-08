package main

import (
	"embed"
	"fmt"

	"github.com/bjartek/overflow"
)

//go:embed flow.json
//go:embed scripts/*
var filesystem embed.FS

func main() {
	fmt.Println("foo")
	o := overflow.Overflow(
		overflow.WithNetwork("mainnet"),
		overflow.WithBasePath(""),
		overflow.WithEmbedFS(filesystem),
	)

	o.Script("test").Print()
}
