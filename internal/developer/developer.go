package developer

import (
	"fmt"
	"realty/internal/residential"
	"strings"
)

type Developer struct {
	Name  string   `json:"name"`
	Label string   `json:"label"`
	Score float64  `json:"score"`
	Links []string `json:"links"`
	residential.Residentials
}

func (d Developer) String() string {
	rs := []string{}
	for _, r := range d.Residentials {
		rs = append(rs, "\n### "+r.String())
	}

	return fmt.Sprintf(
		"# %s\n\n- **Базовый коэффициент:** %.2f\n- **Ссылки:** %s\n## Жилые комплексы застройщика %s\n%s",
		d.Name,
		d.Score,
		"<"+strings.Join(d.Links, ">, <")+">\n",
		d.Name,
		strings.Join(rs, "\n"),
	)
}

func (d *Developer) AddResidentials(rs ...residential.Residential) {
	d.Residentials = append(d.Residentials, rs...)
}

type Developers []Developer

func (ds Developers) String() string {
	s := []string{}
	for _, d := range ds {
		s = append(s, "\n"+d.String())
	}

	return strings.Join(s, "\n---\n")
}

func (ds Developers) AddResidentials(l string, rs ...residential.Residential) {
	for i, d := range ds {
		if d.Label == l {
			ds[i].AddResidentials(rs...)
		}
	}
}

func (ds Developers) ByName(n string) *Developer {
	for i, d := range ds {
		if d.Name == n {
			return &ds[i]
		}
	}
	return nil
}

func (ds Developers) ByLabel(l string) *Developer {
	for i, d := range ds {
		if d.Label == l {
			return &ds[i]
		}
	}
	return nil
}

func (ds Developers) ResidentialByLabel(l string) *residential.Residential {
	for _, d := range ds {
		if r := d.Residentials.ByLabel(l); r != nil {
			return r
		}
	}
	return nil
}

func (ds Developers) CountApartaments() (result int) {
	for _, d := range ds {
		for _, r := range d.Residentials {
			for range r.Apartments {
				result++
			}
		}
	}

	return result
}

func (ds Developers) Load() {}
