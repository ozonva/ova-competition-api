package utils

import (
	"github.com/google/go-cmp/cmp"
	"ozonva/ova-competition-api/internal/models"
	"testing"
	"time"
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
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8}, 0, nil, true},
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
		{
			[]int{1, 2, 3}, 10, [][]int{{1, 2, 3}}, false,
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

func TestCompetitionsToMap(t *testing.T) {
	competitions := []models.Competition{
		models.NewCompetition(1, "Name 1", time.Now()),
		models.NewCompetition(2, "Name 2", time.Now()),
	}

	mapped, err := CompetitionsToMap(competitions)
	if err != nil {
		t.Fatal("received error when mapping competitions to map", err)
	} else {
		if len(mapped) != len(competitions) {
			t.Fatalf("expected %d elements in map, found: %d", len(competitions), len(mapped))
		} else {
			for _, competition := range competitions {
				if !cmp.Equal(competition.Id, mapped[competition.Id].Id) {
					t.Fatalf("competition %d has different content in map or is absent", competition.Id)
				}
			}
		}
	}

	competitions = append(competitions, models.NewCompetition(1, "Duplicate", time.Now()))
	_, err1 := CompetitionsToMap(competitions)
	if err1 == nil {
		t.Fatal("no duplicate competitions found, but it should")
	}
}
