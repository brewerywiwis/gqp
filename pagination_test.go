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

type newPaginationElemTestCaseInput struct {
	Page  uint64
	Size  uint64
	Total uint64
}

type newPaginationElemTestCase struct {
	Input  newPaginationElemTestCaseInput
	Output *PaginationElem
}

func TestNewPagination(t *testing.T) {
	testCases := []newPaginationElemTestCase{
		{
			Input: newPaginationElemTestCaseInput{
				Page:  0,
				Size:  0,
				Total: 0,
			},
			Output: &PaginationElem{
				Page:   0,
				Size:   0,
				Actual: 0,
				Total:  0,
			},
		},
		{
			Input: newPaginationElemTestCaseInput{
				Page:  0,
				Size:  5,
				Total: 10,
			},
			Output: &PaginationElem{
				Page:   0,
				Size:   5,
				Actual: 5,
				Total:  10,
			},
		},
		{
			Input: newPaginationElemTestCaseInput{
				Page:  0,
				Size:  5,
				Total: 4,
			},
			Output: &PaginationElem{
				Page:   0,
				Size:   5,
				Actual: 4,
				Total:  4,
			},
		},
		{
			Input: newPaginationElemTestCaseInput{
				Page:  1,
				Size:  5,
				Total: 4,
			},
			Output: &PaginationElem{
				Page:   0,
				Size:   5,
				Actual: 4,
				Total:  4,
			},
		},
		{
			Input: newPaginationElemTestCaseInput{
				Page:  3,
				Size:  1,
				Total: 4,
			},
			Output: &PaginationElem{
				Page:   3,
				Size:   1,
				Actual: 1,
				Total:  4,
			},
		},
		{
			Input: newPaginationElemTestCaseInput{
				Page:  3,
				Size:  0,
				Total: 4,
			},
			Output: &PaginationElem{
				Page:   0,
				Size:   0,
				Actual: 0,
				Total:  4,
			},
		},
		{
			Input: newPaginationElemTestCaseInput{
				Page:  10,
				Size:  1,
				Total: 6,
			},
			Output: &PaginationElem{
				Page:   5,
				Size:   1,
				Actual: 1,
				Total:  6,
			},
		},
	}

	for _, tc := range testCases {
		elem := NewPaginationElem(tc.Input.Page, tc.Input.Size, tc.Input.Total)
		assert.Equal(t, tc.Output, elem)
	}
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
