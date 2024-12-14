package two

import "testing"

func TestFirst(t *testing.T) {
	input := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}
	got := First(input)
	want := 2
	if got != want {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestSecond(t *testing.T) {
	input := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}
	got := Second(input)
	want := 4
	if got != want {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

// Test the code stolen from Stackoverflwo to remove an Element
func TestRemove(t *testing.T) {
	want := []int{1, 2, 4, 5}
	input := []int{1, 2, 3, 4, 5}
	got := remove(input, 2)

	mismatch := false

	for i := range got {
		if want[i] != got[i] {
			mismatch = true
		}
	}

	if mismatch {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestParseLine(t *testing.T) {
	want := []int{1, 2, 3, 4, 5}
	input := "1 2 3 4 5"
	got, err := parseLine(input)
	if err != nil {
		t.Errorf("Error %v", err)
	}

	mismatch := false

	for i := range got {
		if want[i] != got[i] {
			mismatch = true
		}
	}

	if mismatch {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

// Block of Tests which should be Table-Tests ...
func TestCheckSafetyIncrease(t *testing.T) {
	want := true
	input := []int{1, 2, 3, 4, 5}
	got := checkSafety(input)

	if got != want {
		t.Errorf("mismatch, want %v got %v", want, got)
	}
}

func TestCheckSafetySame(t *testing.T) {
	want := false
	input := []int{1, 2, 2, 4, 5}
	got := checkSafety(input)

	if got != want {
		t.Errorf("mismatch, want %v got %v", want, got)
	}
}

func TestCheckSafetyDecrese(t *testing.T) {
	want := true
	input := []int{5, 4, 3, 2, 1}
	got := checkSafety(input)

	if got != want {
		t.Errorf("mismatch, want %v got %v", want, got)
	}
}

func TestCheckSafetySwap(t *testing.T) {
	want := false
	input := []int{5, 4, 3, 4, 1}
	got := checkSafety(input)

	if got != want {
		t.Errorf("mismatch, want %v got %v", want, got)
	}
}

func TestCheckSafetyUnsafeIncrease(t *testing.T) {
	want := false
	input := []int{1, 2, 7, 8, 9}
	got := checkSafety(input)

	if got != want {
		t.Errorf("mismatch, want %v got %v", want, got)
	}
}

func TestCheckSafetyUnsafeDecrease(t *testing.T) {
	want := false
	input := []int{9, 7, 6, 2, 1}
	got := checkSafety(input)

	if got != want {
		t.Errorf("mismatch, want %v got %v", want, got)
	}
}
