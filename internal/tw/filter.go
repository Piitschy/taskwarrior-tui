package tw

import "strings"

type Filter struct {
	key   string
	value string
}

func NewFilter(key, value string) Filter {
	return Filter{key, value}
}

func NewFilterFromString(filterString string) Filter {
	kv := strings.Split(filterString, ":")
	return Filter{kv[0], kv[1]}
}

func (f Filter) String() string {
	return f.key + ":" + f.value
}

type Filters []Filter

func NewFilters() *Filters {
	return &Filters{}
}

func (f *Filters) AddFilter(key, value string) {
	*f = append(*f, NewFilter(key, value))
}

func (f *Filters) AddFilterFromString(filterString string) {
	*f = append(*f, NewFilterFromString(filterString))
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
