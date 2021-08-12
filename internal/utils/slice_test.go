package utils

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

type ToBatchCase struct {
	sliceToSplit  []int
	batchSize     int
	expected      [][]int
	expectedError bool
}

type FilterWordsCase struct {
	sliceToFilter []string
	wordsToFilter []string
	expected      []string
}

func TestIntSliceToBatches(t *testing.T) {
	testCases := []ToBatchCase{
		{
			nil, 10, nil, true,
		},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 0, nil, true},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8}, 7, [][]int{{1, 2, 3, 4, 5, 6, 7}, {8}}, false,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8}, 3, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8}}, false,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8}, 2, [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, false,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8}, 1, [][]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}}, false,
		},
	}

	for _, testCase := range testCases {
		batches, err := IntSliceToBatches(testCase.sliceToSplit, testCase.batchSize)
		if err != nil {
			if !testCase.expectedError {
				t.Fatal("slicing error", err)
			}
		}

		if !cmp.Equal(batches, testCase.expected) {
			t.Fatal("slices do not match", batches, testCase.expected)
		}
	}
}

func TestFilterWords(t *testing.T) {
	testCases := []FilterWordsCase{
		{
			[]string{}, []string{}, []string{},
		},
		{
			[]string{"a", "b", "c", "d"}, []string{}, []string{"a", "b", "c", "d"},
		},
		{
			[]string{"a", "b", "c", "d"}, []string{"a", "c"}, []string{"b", "d"},
		},
		{
			[]string{"a", "b", "c", "d"}, []string{"a", "b", "c", "d"}, []string{},
		},
		{
			[]string{"a", "b", "c", "d"}, []string{"a", "a", "a", "a"}, []string{"b", "c", "d"},
		},
	}

	for _, testCase := range testCases {
		filtered := FilterWords(testCase.sliceToFilter, testCase.wordsToFilter)

		if !cmp.Equal(filtered, testCase.expected) {
			t.Fatal("slices do not match", filtered, testCase.expected)
		}
	}
}
