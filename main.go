package main

import (
	"realty/internal/developer"
	"realty/internal/factor"
	"realty/internal/tools/loader"
	"realty/internal/tools/scorer"
)

var ds = developer.Developers{}
var fs = factor.Factors{}
var s = scorer.Scorer{}

func init() {
	// Загружаем счетовода
	loader.Load("source/config.json", &s)

	// Загружаем факторы
	loader.Load("source/factors.json", &fs)

	// Загружаем застройщиков
	loader.Load("source/developer.json", &ds)
	loader.Init(ds)
}

func main() {
	s.Estimate(ds, fs)

	loader.ToCSV("out/realty.csv", s)
	loader.Dump("out/realty.md", ds)
}
