package gqp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type parseFilterTestCase struct {
	Input     string
	Output    map[string][]FilterElem
	OutputErr bool
}

func TestParseFilter(t *testing.T) {
	testCases := []parseFilterTestCase{
		{
			Input:     "",
			Output:    map[string][]FilterElem{},
			OutputErr: false,
		},
		{
			Input: `{"price":{"lt":10.00,"gt":1.00}, "created_at":{"eq":"2023-07-28T01:55:58.207207+07:00"}}`,
			Output: map[string][]FilterElem{
				"price": {
					{Op: "gt", Value: 1.00},
					{Op: "lt", Value: 10.00},
				},
				"created_at": {
					{Op: "eq", Value: "2023-07-28T01:55:58.207207+07:00"},
				},
			},
			OutputErr: false,
		},
		{
			Input: `{"price":{"lt":10.00,"gt":1.00},"created_at":{"eq":"2023-07-28T01:55:58.207207+07:00"}}`,
			Output: map[string][]FilterElem{
				"price": {
					{Op: "gt", Value: 1.00},
					{Op: "lt", Value: 10.00},
				},
				"created_at": {
					{Op: "eq", Value: "2023-07-28T01:55:58.207207+07:00"},
				},
			},
			OutputErr: false,
		},
		{
			Input: `{"created_at":{"eq":"2023-07-28T01:55:58.207207+07:00"},"price":{"lt":10.00,"gt":1.00}}`,
			Output: map[string][]FilterElem{
				"created_at": {
					{Op: "eq", Value: "2023-07-28T01:55:58.207207+07:00"},
				},
				"price": {
					{Op: "gt", Value: 1.00},
					{Op: "lt", Value: 10.00},
				},
			},
			OutputErr: false,
		},
		{
			Input: `{"created_at":{"eq":"2023-07-28T01:55:58.207207+07:00"},"price":{"or":[{"lt":10.00},{"gt":1.00}]}}`,
			Output: map[string][]FilterElem{
				"created_at": {
					{Op: "eq", Value: "2023-07-28T01:55:58.207207+07:00"},
				},
				"price": {
					{Op: "or", Value: []FilterElem{
						{Op: "gt", Value: 1.00},
						{Op: "lt", Value: 10.00},
					}},
				},
			},
			OutputErr: false,
		},
		{
			Input: `{"created_at":{"not":{"eq":"2023-07-28T01:55:58.207207+07:00"}},"price":{"or":[{"lt":10.00},{"gt":1.00}]}}`,
			Output: map[string][]FilterElem{
				"created_at": {
					{Op: "not", Value: []FilterElem{
						{Op: "eq", Value: "2023-07-28T01:55:58.207207+07:00"}},
					},
				},
				"price": {
					{Op: "or", Value: []FilterElem{
						{Op: "gt", Value: 1.00},
						{Op: "lt", Value: 10.00},
					}},
				},
			},
			OutputErr: false,
		},
		{
			Input: `{"created_at":{"not":{"eq":"2023-07-28T01:55:58.207207+07:00"}},"price":{"or":[{"lt":10.00005},{"gt":1.01}]}}`,
			Output: map[string][]FilterElem{
				"created_at": {
					{Op: "not", Value: []FilterElem{
						{Op: "eq", Value: "2023-07-28T01:55:58.207207+07:00"}},
					},
				},
				"price": {
					{Op: "or", Value: []FilterElem{
						{Op: "gt", Value: 1.01},
						{Op: "lt", Value: 10.00005},
					}},
				},
			},
			OutputErr: false,
		},
		{
			Input: `{"created_at":{"eq":"2023-07-28T01:55:58.207207+07:00"},"price":{"or":[{"lt":10.00},{"and":[{"lt":6.00},{"gt":1.00}]}]}}`,
			Output: map[string][]FilterElem{
				"created_at": {
					{Op: "eq", Value: "2023-07-28T01:55:58.207207+07:00"},
				},
				"price": {
					{Op: "or", Value: []FilterElem{
						{Op: "and", Value: []FilterElem{
							{Op: "gt", Value: 1.00},
							{Op: "lt", Value: 6.00},
						}},
						{Op: "lt", Value: 10.00},
					}},
				},
			},
			OutputErr: false,
		},
	}

	i := 0
	for _, tc := range testCases {
		res, err := parseFilter(tc.Input)
		if tc.OutputErr != (err != nil) {
			assert.Fail(t, "parse filter error", err)
		} else {
			i++
			assert.Equal(t, tc.Output, res, "case: %v", i)
		}

	}
}

func BenchmarkParseFilterSmallCase(b *testing.B) {
	t := parseFilterTestCase{
		Input: `{"price":{"lt":10.00,"gt":1.00}, "created_at":{"eq":"2023-07-28T01:55:58.207207+07:00"}}`,
		Output: map[string][]FilterElem{
			"price": {
				{Op: "gt", Value: 1.00},
				{Op: "lt", Value: 10.00},
			},
			"created_at": {
				{Op: "eq", Value: "2023-07-28T01:55:58.207207+07:00"},
			},
		},
		OutputErr: false,
	}
	for n := 0; n < b.N; n++ {
		_, err := parseFilter(t.Input)
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkParseFilterMediumCase(b *testing.B) {
	t := parseFilterTestCase{
		Input: `{"created_at":{"not":{"eq":"2023-07-28T01:55:58.207207+07:00"}},"price":{"or":[{"lt":10.00005},{"gt":1.01}]}}`,
		Output: map[string][]FilterElem{
			"created_at": {
				{Op: "not", Value: []FilterElem{
					{Op: "eq", Value: "2023-07-28T01:55:58.207207+07:00"}},
				},
			},
			"price": {
				{Op: "or", Value: []FilterElem{
					{Op: "gt", Value: 1.01},
					{Op: "lt", Value: 10.00005},
				}},
			},
		},
		OutputErr: false,
	}
	for n := 0; n < b.N; n++ {
		_, err := parseFilter(t.Input)
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkParseFilterLargeCase(b *testing.B) {
	t := parseFilterTestCase{
		Input: `{"created_at":{"eq":"2023-07-28T01:55:58.207207+07:00"},"price":{"or":[{"lt":10.00},{"and":[{"lt":6.00},{"gt":1.00}]}]}}`,
		Output: map[string][]FilterElem{
			"created_at": {
				{Op: "eq", Value: "2023-07-28T01:55:58.207207+07:00"},
			},
			"price": {
				{Op: "or", Value: []FilterElem{
					{Op: "and", Value: []FilterElem{
						{Op: "gt", Value: 1.00},
						{Op: "lt", Value: 6.00},
					}},
					{Op: "lt", Value: 10.00},
				}},
			},
		},
		OutputErr: false,
	}
	for n := 0; n < b.N; n++ {
		_, err := parseFilter(t.Input)
		if err != nil {
			b.FailNow()
		}
	}
}
