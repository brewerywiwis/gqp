package gqp

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type PaginationElem struct {
	Page   uint64 `json:"page"`
	Size   uint64 `json:"size"`
	Actual uint64 `json:"actual"`
	Total  uint64 `json:"total"`
}

func NewPaginationElem(page uint64, size uint64, total uint64) *PaginationElem {
	if size <= 0 {
		return &PaginationElem{Page: 0, Size: 0, Actual: 0, Total: total}
	}

	p := PaginationElem{
		Page: page,
		Size: size,
	}
	p = p.AdjustWithTotalData(total)

	return &p
}
func (p PaginationElem) AdjustWithTotalData(total uint64) PaginationElem {
	actualPage := p.Page
	actualSize := p.Size
	if (actualPage+1)*actualSize > total {
		actualPage = uint64(math.Ceil(float64(total)/float64(actualSize))) - 1
		actualSize = total - (actualPage * actualSize)
	}

	return PaginationElem{
		Page:   actualPage,
		Size:   p.Size,
		Actual: actualSize,
		Total:  total,
	}
}

func (p *PaginationElem) ToSQL() string {
	return fmt.Sprintf("limit %v offset %v", p.Size, p.Page*p.Size)
}

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
		Page: uint64(p),
		Size: uint64(s),
	}, nil
}
