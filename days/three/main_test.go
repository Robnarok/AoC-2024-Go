package three

import "testing"

func TestSecond(t *testing.T) {
	input := []string{
		"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
	}
	got := Second(input)
	want := 96
	if got != want {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestPrepareLine(t *testing.T) {
	want := "amul(1,2)527489mul(4,3)AAAAmul(123)"
	input := "amul(1,2)527489mul(4,3)don't()mul(123,4)do()mul(123)don't()123"
	got := prepareLine(input)

	if want != got {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestFilterLine(t *testing.T) {
	want := []string{"mul(1,2)", "mul(4,3)"}
	input := "amul(1,2)527489mul(4,3)"
	got := filterLine(input)

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

func TestParseCommand(t *testing.T) {
	want1 := 1
	want2 := 2
	input := "mul(1,2)"
	got1, got2, err := parseCommand(input)
	if err != nil {
		t.Fatal(err)
	}

	if want1 != got1 || want2 != got2 {
		t.Errorf("want1: %v; want2: %v; got1: %v got2: %v", want1, want2, got1, got2)
	}
}
