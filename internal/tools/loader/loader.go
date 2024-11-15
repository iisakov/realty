package loader

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"realty/internal/apartment"
	"realty/internal/developer"
	"realty/internal/residential"
	"realty/internal/tools/scorer"
	"strings"
)

type Loadable interface {
	Load()
}

type Dumpable interface {
	Dump()
}

func Load(fp string, l Loadable) {
	f, err := os.ReadFile(fp)
	if err != nil {
		log.Fatal("не смогли открыть файл: ", err)
	}

	err = json.Unmarshal([]byte(f), &l)
	if err != nil {
		log.Fatal("не удалось распарсить файл: ", err)
	}
}

func Dump(fp string, d fmt.Stringer) {
	f, err := os.Create(fp)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.WriteString(d.String())
	if err != nil {
		log.Fatal(err)
	}
}

func ToCSV(fp string, s scorer.Scorer) {

	f, err := os.Create(fp)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	h, v := s.FormatToCSV()
	writer.Write(h)
	writer.WriteAll(v)

}

func Init(ds developer.Developers) {
	residentialDir, err := filepath.Glob("source/residentials/*.json")
	if err != nil {
		log.Fatal(err)
	}

	for _, fp := range residentialDir {
		var rs residential.Residentials
		dl := strings.Split(filepath.Base(fp), ".")[0]

		Load(fp, &rs)
		ds.AddResidentials(dl, rs...)
	}

	apartmentDir, err := filepath.Glob("source/apartaments/*.json")
	if err != nil {
		log.Fatal(err)
	}

	for _, fp := range apartmentDir {
		var as apartment.Apartments
		rl := strings.Split(filepath.Base(fp), ".")[0]

		Load(fp, &as)
		ds.ResidentialByLabel(rl).AddApartments(as...)
	}
}
