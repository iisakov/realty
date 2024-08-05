package main

import (
	"realty/internal/developer"
	"realty/internal/tools/loader"
)

var ds = developer.Developers{}

func init() {
	// Загружаем застройщиков
	loader.Load("source/developer.json", &ds)
	loader.Init(ds)

}

func main() {
	loader.Dump("out/realty.md", ds)
}
