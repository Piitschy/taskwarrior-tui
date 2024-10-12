package tw

import (
	"errors"
	"strings"
)

type Filter struct {
	key   string
	value string
}

func NewFilter(key, value string) (Filter, error) {
	if key == "" {
		return Filter{}, errors.New("Key cannot be empty")
	}
	return Filter{key, value}, nil
}

func NewFilterFromString(filterString string) (Filter, error) {
	kv := strings.Split(filterString, ":")
	if len(kv) != 2 {
		return Filter{}, errors.New("Invalid filter string")
	}
	return NewFilter(kv[0], kv[1])
}

func (f Filter) String() string {
	return f.key + ":" + f.value
}

type Filters []Filter

func NewFilters() *Filters {
	return &Filters{}
}

func (f *Filters) AddFilter(key, value string) error {
	filter, err := NewFilter(key, value)
	if err != nil {
		return err
	}
	*f = append(*f, filter)
	return nil
}

func (f *Filters) AddFilterFromString(filterString string) error {
	filter, err := NewFilterFromString(filterString)
	if err != nil {
		return err
	}
	*f = append(*f, filter)
	return nil
}

func (f *Filters) String() string {
	filterStrings := make([]string, len(*f))
	for _, filter := range *f {
		filterStrings = append(filterStrings, filter.key+":"+filter.value)
	}
	filterString := "(" + filterStrings[0]
	for _, filter := range filterStrings[1:] {
		filterString += " and " + filter
	}
	filterString += ")"
	return filterString
}
