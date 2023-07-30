package gqp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type parsePaginationTestCase struct {
	Input     string
	Output    PaginationElem
	OutputErr bool
}

func TestParsePagination(t *testing.T) {
	testCases := []parsePaginationTestCase{
		{
			Input: "1,1",
			Output: PaginationElem{
				Page: 1,
				Size: 1,
			},
			OutputErr: false,
		},
		{
			Input: "1,2",
			Output: PaginationElem{
				Page: 1,
				Size: 2,
			},
			OutputErr: false,
		},
		{
			Input:     "-1,1",
			Output:    PaginationElem{},
			OutputErr: true,
		},
		{
			Input:     "1,-1",
			Output:    PaginationElem{},
			OutputErr: true,
		},
		{
			Input:     "0,0",
			Output:    PaginationElem{},
			OutputErr: false,
		},
		{
			Input:     "",
			Output:    PaginationElem{},
			OutputErr: false,
		},
		{
			Input:     ",",
			Output:    PaginationElem{},
			OutputErr: true,
		},
		{
			Input:     "a,b",
			Output:    PaginationElem{},
			OutputErr: true,
		},
		{
			Input:     "a,b,c",
			Output:    PaginationElem{},
			OutputErr: true,
		},
	}

	for _, tc := range testCases {
		res, err := parsePagination(tc.Input)
		if tc.OutputErr != (err != nil) {
			assert.Fail(t, "parse pagination error: ", err)
		} else {
			assert.Equal(t, tc.Output, res)
		}

	}
}
