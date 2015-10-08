package main

import (
	"github.com/Qs-F/PoiBlog"
)

func main() {
	var b poiblog.Blog
	b.ReadConf("/Users/QsF/tmp/rtest.json")
	b.NewServer()
}
