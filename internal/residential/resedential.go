package residential

import (
	"fmt"
	"realty/internal/apartment"
	"strings"
)

type Residential struct {
	Name           string   `json:"name"`
	Label          string   `json:"label"`
	Location       string   `json:"location"`
	Yard           []string `json:"yard"`
	Parking        []string `json:"parking"`
	Finishing      string   `json:"finishing"`
	MaxFloor       int      `json:"maxFloor"`
	Infrastructure []string `json:"infrastructure"`
	Address        string   `json:"address"`
	CompletionDate string   `json:"completionDate"`
	Links          []string `json:"links"`
	apartment.Apartments
}

func (r Residential) String() string {
	as := []string{}
	for _, a := range r.Apartments {
		as = append(as, "\n"+a.String())
	}

	return fmt.Sprintf(

		"%s\n\n- **Расположение:** %s\n- **Окончание строительства:** %s\n- **Двор:** %s\n- **Парковка:** %s\n- **Отделка:** %s\n- **Максимальный этаж:** %d\n- **Инфраструктура:** %s\n- **Адрес:** %s\n- **Ссылки:** %s\n#### Квартиры в %s\n%s",
		r.Name,
		r.Location,
		r.CompletionDate,
		strings.Join(r.Yard, ", "),
		strings.Join(r.Parking, ", "),
		r.Finishing,
		r.MaxFloor,
		strings.Join(r.Infrastructure, ", "),
		r.Address,
		strings.Join(r.Links, ", ")+"\n",
		r.Name,
		strings.Join(as, ",\n"),
	)
}

func (r *Residential) AddApartments(as ...apartment.Apartment) {
	r.Apartments = append(r.Apartments, as...)
}

type Residentials []Residential

func (rs Residentials) String() string {
	s := []string{}
	for _, r := range rs {
		s = append(s, r.String())
	}

	return strings.Join(s, "\n")
}

func (rs Residentials) ByName(n string) *Residential {
	for i, r := range rs {
		if r.Name == n {
			return &rs[i]
		}
	}
	return nil
}

func (rs Residentials) ByLabel(l string) *Residential {
	for i, r := range rs {
		if r.Label == l {
			return &rs[i]
		}
	}
	return nil
}

func (rs Residentials) Load() {}
