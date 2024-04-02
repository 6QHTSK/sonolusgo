package sonolusgo

import (
	"fmt"
	"strconv"
)

type ServerOptionSection struct {
	Type    string         `json:"string"`
	Title   string         `json:"title"`
	Icon    string         `json:"icon"`
	Options []ServerOption `json:"options"`
}

type ServerOption interface {
	GetQuery() string
	GetValueStr(queryResult string) string
}

type SearchTextOption struct {
	Query       string `json:"query"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	PlaceHolder string `json:"placeHolder"`
}

func (o SearchTextOption) GetQuery() string {
	return o.Query
}

func (o SearchTextOption) GetValueStr(queryResult string) string {
	return queryResult
}

func NewSearchTextOption(query string, name string, placeHolder string) SearchTextOption {
	return SearchTextOption{Query: query, Name: name, Type: "text", PlaceHolder: placeHolder}
}

type SearchSliderOption struct {
	Query string  `json:"query"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Def   float64 `json:"def"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Step  float64 `json:"step"`
	Unit  string  `json:"unit"`
}

func (o SearchSliderOption) GetQuery() string {
	return o.Query
}

func (o SearchSliderOption) GetValueStr(queryResult string) string {
	value, err := strconv.ParseFloat(queryResult, 64)
	if err != nil {
		return fmt.Sprintf("%f", o.Def)
	}
	if value < o.Min || value > o.Max {
		return fmt.Sprintf("%f", o.Def)
	}
	return queryResult
}

func NewSearchSliderOption(query string, name string, def float64, min float64, max float64, step float64, unit string) SearchSliderOption {
	return SearchSliderOption{Query: query, Name: name, Type: "slider", Def: def, Min: min, Max: max, Step: step, Unit: unit}
}

type SearchToggleOption struct {
	Query string `json:"query"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Def   int    `json:"def"`
}

func (o SearchToggleOption) GetQuery() string {
	return o.Query
}

func (o SearchToggleOption) GetValueStr(queryResult string) string {
	value, err := strconv.ParseFloat(queryResult, 64)
	if err != nil {
		return strconv.Itoa(o.Def)
	}
	if value != 0 && value != 1 {
		return strconv.Itoa(o.Def)
	}
	return queryResult
}

func NewSearchToggleOption(query string, name string, def int) SearchToggleOption {
	if def != 0 && def != 1 {
		panic("Toggle should be 0 or 1")
	}
	return SearchToggleOption{Query: query, Name: name, Type: "toggle", Def: def}
}

type SearchSelectOption struct {
	Query  string   `json:"query"`
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Def    int      `json:"def"`
	Values []string `json:"values,omitempty"`
}

func (o SearchSelectOption) GetQuery() string {
	return o.Query
}

func (o SearchSelectOption) GetValueStr(queryResult string) string {
	index, err := strconv.Atoi(queryResult)
	if err != nil || index < 0 {
		index = 0
	} else if index >= len(o.Values) {
		index = len(o.Values) - 1
	}
	return o.Values[index]
}

func NewSearchSelectOption(query string, name string, def int, values []string) SearchSelectOption {
	if def >= len(values) {
		panic("Def exceed values list")
	}
	return SearchSelectOption{Query: query, Name: name, Type: "select", Def: def, Values: values}
}
