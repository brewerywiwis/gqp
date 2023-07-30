package gqp

import (
	"fmt"
	"strconv"
	"strings"
)

func parsePagination(input string) (PaginationElem, error) {
	sentence := strings.TrimSpace(input)
	if len(sentence) <= 0 {
		return PaginationElem{}, nil
	}

	tokens := strings.Split(sentence, ",")
	if len(tokens) != 2 {
		return PaginationElem{}, fmt.Errorf("pagination pattern is not valid example -> page=10,10 got: %v", input)
	}
	for i, t := range tokens {
		tokens[i] = strings.TrimSpace(t)
	}

	p, err := strconv.Atoi(tokens[0])
	if err != nil {
		return PaginationElem{}, err
	}
	s, err := strconv.Atoi(tokens[1])
	if err != nil {
		return PaginationElem{}, err
	}

	if p < 0 || s < 0 {
		return PaginationElem{}, fmt.Errorf("pagination value is not lower than 0 -> got (page: %v, size:%v)", p, s)
	}

	return PaginationElem{
		Page: uint(p),
		Size: uint(s),
	}, nil
}
