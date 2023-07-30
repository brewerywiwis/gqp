package gqp

type QueryInput struct {
	Filter     string
	Sort       string
	Pagination string
}

type QueryOutput struct {
	Filter     map[string][]FilterElem
	Sort       []SortElem
	Pagination PaginationElem
}

func (o *QueryOutput) ToSQL() string {
	return ""
}
