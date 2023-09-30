package sort

import (
	"testing"

	sorts "github.com/ayoubzulfiqar/TheAlgorithms/algo/sort"
)

func TestBubble(t *testing.T) {

	got := sorts.BubbleSort()
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q\n", got, want)
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sorts.BubbleSort()
	}
}