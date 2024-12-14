package one

import (
	"testing"
)

func TestReadValues(t *testing.T) {
	input := "3   4"
	left, right, err := readValues(input)
	if err != nil {
		t.Errorf("err != nil; got %v", err)
	}
	if left != 3 {
		t.Errorf("left = %d, expected %d", left, 3)
	}
	if right != 4 {
		t.Errorf("left = %d, expected %d", left, 4)
	}
}

// This is kind of stupid, i am only wrapping a library function...
func TestSortNumbers(t *testing.T) {
	lefts := []int{2, 4, 3}
	rights := []int{2, 4, 3}
	gotLefts, gotRights := sortNumbers(lefts, rights)
	wantLefts := []int{2, 3, 4}
	wantRights := []int{2, 3, 4}

	missmatch := false
	for i := range gotLefts {
		if gotLefts[i] != wantLefts[i] {
			missmatch = true
		}
	}

	if missmatch {
		t.Errorf("left = %v, expected %v", gotLefts, wantLefts)
	}

	missmatch = false
	for i := range gotRights {
		if gotRights[i] != wantRights[i] {
			missmatch = true
		}
	}
	if missmatch {
		t.Errorf("left = %v, expected %v", gotRights, wantRights)
	}
}

// This is kind of stupid, i am only wrapping a library function...
func TestGetDifference(t *testing.T) {
	left := 3
	right := 5
	want := 2

	got := getDifference(left, right)

	if got != want {
		t.Errorf("missmatch of difference, Want: %d, got %d", want, got)
	}
}
