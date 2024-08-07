package scorer

import (
	"fmt"
	"math"
	"realty/internal/developer"
	"realty/internal/factor"
	"reflect"
	"slices"
	"sort"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type score struct {
	RowDescriptions []map[string]string
	Value           []map[string]float64
}

func (s score) FormatToCSV() (header []string, value [][]string) {

	for _, v := range s.Value {
		for vk := range v {
			if !slices.Contains(header, vk) && vk != "Total" {
				header = append(header, vk)
			}
		}
	}

	sort.Slice(header, func(i, j int) bool {
		return header[j] > header[i]
	})

	for _, rd := range s.RowDescriptions {
		value = append(value, []string{})
		for rdk := range rd {
			if !slices.Contains(header, rdk) {
				header = append(header, rdk)
			}
		}
	}

	header = append(header, "Total")

	for _, h := range header {
		i := 0
		for _, rd := range s.RowDescriptions {
			if v, ok := rd[h]; ok {
				value[i] = append(value[i], v)
			}
			i++
		}
		i = 0
		for _, v := range s.Value {
			if vv, ok := v[h]; ok {
				value[i] = append(value[i], fmt.Sprintf("%.2f", vv))
			}
			i++
		}
		i = 0
	}

	return
}

type Scorer struct {
	MaxCost           int     `json:"maxCost"`
	MaxInitialPayment int     `json:"maxInitialPayment"`
	AvgAreaCost       float64 `json:"avgAreaCost"`
	MinArea           float64 `json:"minArea"`

	score
}

type Scorable interface {
}

func (s Scorer) Load() {}

func (s *Scorer) Estimate(ds developer.Developers, fs factor.Factors) {

	maxAreaCost, maxArea, maxCost := 0.0, 0.0, 0

	for _, d := range ds {
		for _, r := range d.Residentials {
			for _, a := range r.Apartments {
				aAreaCost := float64(a.Cost) / a.Area
				if maxAreaCost < aAreaCost {
					maxAreaCost = aAreaCost
				}
				if maxArea < a.Area {
					maxArea = a.Area
				}
				if maxCost < a.Cost {
					maxCost = a.Cost
				}
			}
		}
	}

	rowNum := 0
	s.RowDescriptions = []map[string]string{}
	s.Value = []map[string]float64{}

	for _, d := range ds {
		for _, r := range d.Residentials {
			for _, a := range r.Apartments {
				s.RowDescriptions = append(s.RowDescriptions, map[string]string{})
				s.Value = append(s.Value, map[string]float64{})
				totalScore := d.Score

				aAreaCost := float64(a.Cost) / a.Area

				tag := " - "
				if len(a.Tags) >= 1 {
					tag = strings.Join(a.Tags, " ")
				}
				link := " - "
				if len(a.Links) > 0 {
					link = a.Links[0]
				}

				s.RowDescriptions[rowNum]["tag"] = tag
				s.RowDescriptions[rowNum]["link"] = link

				s.Value[rowNum]["AreaCostScore"] = roundFloat(float64(aAreaCost)/maxAreaCost, 2)
				s.Value[rowNum]["AreaScore"] = roundFloat(maxArea/float64(a.Area), 2)
				s.Value[rowNum]["CostScore"] = roundFloat(float64(a.Cost)/float64(maxCost), 2)

				s.Value[rowNum]["AreaCost"] = aAreaCost
				s.Value[rowNum]["Area"] = a.Area
				s.Value[rowNum]["Cost"] = float64(a.Cost)

				s.Value[rowNum]["divAreaCost"] = roundFloat(aAreaCost-s.AvgAreaCost, 2)
				s.Value[rowNum]["divArea"] = roundFloat(a.Area-s.MinArea, 2)
				s.Value[rowNum]["divCost"] = roundFloat(float64(s.MaxCost-a.Cost), 2)

				baseScore := roundFloat(float64(2-float64(aAreaCost)/maxAreaCost*maxArea/float64(a.Area)*float64(a.Cost)/float64(maxCost)), 2)
				s.Value[rowNum]["baseScore"] = baseScore

				totalScore *= baseScore

				for _, factor := range fs {

					// Для ЖК
					curentScore, probableTotalScore, ok := addFactor(r, factor, fs, totalScore)
					if ok {
						s.Value[rowNum][factor.Factor] = curentScore
						totalScore = probableTotalScore
					}

					// Для квартир
					curentScore, probableTotalScore, ok = addFactor(a, factor, fs, totalScore)
					if ok {
						s.Value[rowNum][factor.Factor] = curentScore
						totalScore = probableTotalScore
					}
					if factor.Factor == "floor" {
						score := 1.0
						switch {
						case a.Floor == 1:
							score, ok = fs.ByValue("Первый")

						case a.Floor == r.MaxFloor:
							score, ok = fs.ByValue("Верхний")

						case a.Floor < r.MaxFloor/2:
							score, ok = fs.ByValue("Ниже середины")

						case a.Floor >= r.MaxFloor/2:
							score, ok = fs.ByValue("Выше середины")
						}
						if ok {
							s.Value[rowNum]["floor"] = score
							totalScore *= score
						}
					}
				}

				windowsNumScore := float64(a.WindowsNum)/100.00 + 1
				s.Value[rowNum]["windowsNum"] = windowsNumScore
				totalScore *= windowsNumScore

				s.Value[rowNum]["Total"] = roundFloat(totalScore, 2)
				rowNum++
			}
		}
	}
}

func (s Scorer) Score() score {
	return s.score
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func addFactor(item Scorable, f factor.Factor, fs factor.Factors, totalScore float64) (float64, float64, bool) {

	itemV := reflect.ValueOf(item)
	itemFactorV := itemV.FieldByName(cases.Title(language.English).String(f.Factor))

	switch {
	case itemFactorV.Kind() == reflect.Slice:
		score := 1.0
		for i := 0; i < itemFactorV.Len(); i++ {
			if currentScore, ok := fs.ByValue(itemFactorV.Index(i).String()); ok {
				score *= currentScore
				totalScore *= score
			}
		}
		return roundFloat(score, 2), totalScore, true

	case itemFactorV.Kind() == reflect.String:
		if score, ok := fs.ByValue(itemFactorV.String()); ok {
			return roundFloat(score, 2), totalScore, true
		}
	}
	return 0, 0, false
}
