package tw

import (
	"errors"
	"strings"
)

type Filter struct {
	key      string
	value    string
	Disabled bool
}

func NewFilter(key, value string) (Filter, error) {
	if key == "" {
		return Filter{}, errors.New("Key cannot be empty")
	}
	return Filter{key, value, false}, nil
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

func NewFiltersFromString(filterString string) (*Filters, error) {
	filters := strings.Split(filterString, " and ")
	var f Filters
	for _, filter := range filters {
		filter, err := NewFilterFromString(filter)
		if err != nil {
			continue
		}
		f = append(f, filter)
	}
	return &f, nil
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

func (f Filters) String() string {
	filterStrings := make([]string, len(f))
	for i, filter := range f {
		filterStrings[i] = filter.String()
	}
	filterString := "(" + filterStrings[0]
	for _, filter := range filterStrings[1:] {
		filterString += " and " + filter
	}
	filterString += ")"
	return filterString
}
