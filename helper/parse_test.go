package helper_test

import (
	"testing"

	"github.com/Robnarok/AoC-2024-Go/helper"
)

func TestParseLength(t *testing.T) {
	filename := "../input/test-1.txt"
	content := helper.Parse(filename)
	want := []string{"1", "2", "3"}
	if len(content) != len(want) {
		t.Fatalf("helper.Parse(\"test-1\").len() != %d, got %d", len(want), len(content))
	}
}

func TestParseContent(t *testing.T) {
	filename := "../input/test-1.txt"
	content := helper.Parse(filename)

	want := []string{"1", "2", "3"}
	if len(content) != len(want) {
		t.Fatalf("helper.Parse(\"test-1\").len() != %d, got %d", len(want), len(content))
	}
	for i := range content {
		if want[i] != content[i] {
			t.Fatalf("content[%d] != %v, got %v. want = %v, got = %v", i, want[i], content[i], want, content)
		}
	}
}
