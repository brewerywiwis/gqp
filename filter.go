package gqp

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

type FilterElem struct {
	Op    string
	Value any
}

func parseFilterElem(fe any) ([]FilterElem, error) {
	result := make([]FilterElem, 0)
	switch c := fe.(type) {
	case map[string]any:
		for k, v := range c {
			switch c2 := v.(type) {
			case string:
				result = append(result, FilterElem{Op: k, Value: c2})
			case int, float64:
				result = append(result, FilterElem{Op: k, Value: c2})
			case map[string]any, []any:
				tmp, err := parseFilterElem(c2)
				if err != nil {
					return []FilterElem{}, fmt.Errorf("value is not support: %v, %v", k, v)
				}
				result = append(result, FilterElem{Op: k, Value: tmp})
			default:
				result = append(result, FilterElem{Op: k, Value: v})
			}
		}
	case []interface{}:
		for _, v := range c {
			tmp, err := parseFilterElem(v)
			if err != nil {
				return []FilterElem{}, err
			}
			result = append(result, tmp...)
		}
	default:
		return []FilterElem{}, fmt.Errorf("value is not support: %v", fe)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Op < result[j].Op
	})

	return result, nil
}

func parseFilter(input string) (map[string][]FilterElem, error) {
	if len(strings.TrimSpace(input)) <= 0 {
		return map[string][]FilterElem{}, nil
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		return map[string][]FilterElem{}, err
	}

	result := make(map[string][]FilterElem, len(data))
	for k, v := range data {
		switch c := v.(type) {
		case map[string]interface{}:
			tmp, err := parseFilterElem(c)
			if err != nil {
				return map[string][]FilterElem{}, err
			}

			result[k] = tmp
		default:
			return map[string][]FilterElem{}, fmt.Errorf("value is not a map")
		}

	}
	return result, nil
}
