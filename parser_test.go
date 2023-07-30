package gqp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type queryParserParseTestCase struct {
	Input     QueryInput
	Output    QueryOutput
	OutputErr bool
}

func TestQueryParser_Parse(t *testing.T) {
	testCases := []queryParserParseTestCase{
		{
			Input: QueryInput{
				Filter:     `{"price":{"lt":10.00,"gt":1.00}, "created_at":{"eq":"2023-07-28T01:55:58.207207+07:00"}}`,
				Sort:       "",
				Pagination: "",
			},
			Output: QueryOutput{
				Filter: map[string][]FilterElem{
					"price": {
						{Op: "gt", Value: 1.00},
						{Op: "lt", Value: 10.00},
					},
					"created_at": {
						{Op: "eq", Value: "2023-07-28T01:55:58.207207+07:00"},
					},
				},
				Sort:       []SortElem{},
				Pagination: PaginationElem{},
			},
			OutputErr: false,
		},
		{
			Input: QueryInput{
				Filter:     "",
				Sort:       "-price, created_at",
				Pagination: "",
			},
			Output: QueryOutput{
				Filter: map[string][]FilterElem{},
				Sort: []SortElem{
					{Key: "price", Desc: true},
					{Key: "created_at", Desc: false},
				},
				Pagination: PaginationElem{},
			},
			OutputErr: false,
		},
		{
			Input: QueryInput{
				Filter:     "",
				Sort:       "",
				Pagination: "1,2",
			},
			Output: QueryOutput{
				Filter: map[string][]FilterElem{},
				Sort:   []SortElem{},
				Pagination: PaginationElem{
					Page: 1,
					Size: 2,
				},
			},
			OutputErr: false,
		},
		{
			Input: QueryInput{
				Filter:     `{"price":{"lt":10.00,"gt":1.00}, "created_at":{"eq":"2023-07-28T01:55:58.207207+07:00"}}`,
				Sort:       "-price, created_at",
				Pagination: "1,2",
			},
			Output: QueryOutput{
				Filter: map[string][]FilterElem{
					"price": {
						{Op: "gt", Value: 1.00},
						{Op: "lt", Value: 10.00},
					},
					"created_at": {
						{Op: "eq", Value: "2023-07-28T01:55:58.207207+07:00"},
					},
				},
				Sort: []SortElem{
					{Key: "price", Desc: true},
					{Key: "created_at", Desc: false},
				},
				Pagination: PaginationElem{
					Page: 1,
					Size: 2,
				},
			},
			OutputErr: false,
		},
	}

	q := NewQueryParser()

	for i, tc := range testCases {
		res, err := q.Parse(tc.Input)
		if tc.OutputErr != (err != nil) {
			assert.Fail(t, "query parser -> parse error", err)
		} else {
			assert.Equal(t, tc.Output, res, "case: %v", i)
		}

	}
}
