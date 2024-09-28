package utils

import "fmt"

func SpaceAround(strings []string) []string {
	spaced := make([]string, len(strings))
	for i, s := range strings {
		spaced[i] = fmt.Sprintf(" %s ", s)
	}
	return spaced
}
