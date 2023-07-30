package gqp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type parseSortTestCase struct {
	Input     string
	Output    []SortElem
	OutputErr bool
}

func TestParseSort(t *testing.T) {
	testCases := []parseSortTestCase{
		{
			Input:     "",
			Output:    []SortElem{},
			OutputErr: false,
		},
		{
			Input:     ",",
			Output:    []SortElem{},
			OutputErr: true,
		},
		{
			Input: "-price, created_at",
			Output: []SortElem{
				{Key: "price", Desc: true},
				{Key: "created_at", Desc: false},
			},
			OutputErr: false,
		},
		{
			Input: "-price, created_at",
			Output: []SortElem{
				{Key: "price", Desc: true},
				{Key: "created_at", Desc: false},
			},
			OutputErr: false,
		},
		{
			Input: "-created_at, price",
			Output: []SortElem{
				{Key: "created_at", Desc: true},
				{Key: "price", Desc: false},
			},
			OutputErr: false,
		},
		{
			Input: "-created_at,price",
			Output: []SortElem{
				{Key: "created_at", Desc: true},
				{Key: "price", Desc: false},
			},
			OutputErr: false,
		},
		{
			Input:     "-,price",
			Output:    []SortElem{},
			OutputErr: true,
		},
		{
			Input:     ",price",
			Output:    []SortElem{},
			OutputErr: true,
		},
		{
			Input:     "price,",
			Output:    []SortElem{},
			OutputErr: true,
		},
		{
			Input:     "price,-",
			Output:    []SortElem{},
			OutputErr: true,
		},
	}

	for _, tc := range testCases {
		res, err := parseSort(tc.Input)
		if tc.OutputErr != (err != nil) {
			assert.Fail(t, "parse sort error: ", err)
		} else {
			assert.Equal(t, tc.Output, res)
		}

	}
}
