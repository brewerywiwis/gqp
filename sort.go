package gqp

import (
	"fmt"
	"strings"
)

func parseSort(input string) ([]SortElem, error) {
	sentence := strings.TrimSpace(input)
	if len(sentence) <= 0 {
		return []SortElem{}, nil
	}

	tokens := strings.Split(sentence, ",")
	output := make([]SortElem, 0, len(tokens))
	for _, t := range tokens {
		desc := false
		t = strings.TrimSpace(t)
		if len(t) <= 0 {
			return []SortElem{}, fmt.Errorf("sort elem is empty -> got input: %v", input)
		}
		if t[0] == '-' {
			desc = true
			t = t[1:]
		}
		if len(t) <= 0 {
			return []SortElem{}, fmt.Errorf("sort elem is empty -> got input: %v", input)
		}
		output = append(output, SortElem{Key: t, Desc: desc})
	}

	return output, nil
}
