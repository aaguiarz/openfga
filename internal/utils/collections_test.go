package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFilter(t *testing.T) {
	cases := map[string]struct {
		src          []int
		dst          []int
		predicate    func(int) bool
		countMatches int
		expectation  []int
	}{
		"equal_length_all_match": {
			src:          []int{1, 2, 3, 4, 5},
			dst:          make([]int, 5),
			predicate:    func(i int) bool { return i == i },
			countMatches: 5,
			expectation:  []int{1, 2, 3, 4, 5},
		},
		"equal_length_no_match": {
			src:          []int{1, 2, 3, 4, 5},
			dst:          make([]int, 5),
			predicate:    func(i int) bool { return i != i },
			countMatches: 0,
			expectation:  []int{0, 0, 0, 0, 0},
		},
		"smaller_dst_all_match": {
			src:          []int{1, 2, 3, 4, 5},
			dst:          make([]int, 2),
			predicate:    func(i int) bool { return i == i },
			countMatches: 2,
			expectation:  []int{1, 2},
		},
		"larger_dst_all_match": {
			src:          []int{1, 2, 3, 4, 5},
			dst:          make([]int, 8),
			predicate:    func(i int) bool { return i == i },
			countMatches: 5,
			expectation:  []int{1, 2, 3, 4, 5, 0, 0, 0},
		},
		"nil_src": {
			src:          nil,
			dst:          make([]int, 5),
			predicate:    func(i int) bool { return i == i },
			countMatches: 0,
			expectation:  []int{0, 0, 0, 0, 0},
		},
		"nil_dst": {
			src:          []int{1, 2, 3, 4, 5},
			dst:          nil,
			predicate:    func(i int) bool { return i == i },
			countMatches: 0,
			expectation:  nil,
		},
		"nil_dst_nil_src": {
			src:          nil,
			dst:          nil,
			predicate:    func(i int) bool { return i == i },
			countMatches: 0,
			expectation:  nil,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			count := Filter(tc.dst, tc.src, tc.predicate)
			require.Equal(t, tc.countMatches, count)
			require.Equal(t, tc.expectation, tc.dst)
		})
	}
}
