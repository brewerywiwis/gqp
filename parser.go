package gqp

type QueryInput struct {
	Filter     string
	Sort       string
	Pagination string
}

type FilterElem struct {
	Op    string
	Value any
}

type SortElem struct {
	Key  string
	Desc bool
}

type PaginationElem struct {
	Page uint
	Size uint
}

type QueryOutput struct {
	Filter     map[string][]FilterElem
	Sort       []SortElem
	Pagination PaginationElem
}

type QueryParser struct {
}

func NewQueryParser() *QueryParser {
	return &QueryParser{}
}

func (q *QueryParser) Parse(input QueryInput) (QueryOutput, error) {
	f, err := parseFilter(input.Filter)
	if err != nil {
		return QueryOutput{}, err
	}
	s, err := parseSort(input.Sort)
	if err != nil {
		return QueryOutput{}, err
	}
	p, err := parsePagination(input.Pagination)
	if err != nil {
		return QueryOutput{}, err
	}

	return QueryOutput{
		Filter:     f,
		Sort:       s,
		Pagination: p,
	}, nil
}
