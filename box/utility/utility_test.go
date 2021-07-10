package utility

import (
	"math/rand"
	"testing"
)

func TestTrimExtension(t *testing.T) {
	want := "filename"

	filepath := "filename.ext"
	got := TrimExtension(filepath)
	if got != want {
		t.Errorf("input: (%s)\ngot: %s\nwanted: %s\n",
			filepath, got, want,
		)
	}
}

func TestAnyLessThanZero(t *testing.T) {
	var numbers []float64
	for i := 0; i < 100; i++ {
		numbers = append(numbers, rand.Float64()*100-50)
	}

	flag := false
	for _, v := range numbers {
		want := v < 0
		got := AnyLessThanZero(v)
		if got != want {
			t.Errorf("input: (%f)\ngot: %t,\nwanted: %t\n",
				v, got, want,
			)
		}
		if got == true {
			flag = true
		}
	}

	want := flag
	got := AnyLessThanZero(numbers...)
	if got != want {
		t.Errorf("input: (%f)\ngot: %t\nwanted: %t\n",
			numbers, got, want,
		)
	}
}

func TestAnyNotPositive(t *testing.T) {
	var numbers []float64
	for i := 0; i < 100; i++ {
		numbers = append(numbers, rand.Float64()*100-50)
	}

	flag := false
	for _, v := range numbers {
		want := v <= 0
		got := AnyNotPositive(v)
		if got != want {
			t.Errorf("input: (%f)\ngot: %t\nwanted: %t\n",
				v, got, want,
			)
		}
		if got == true {
			flag = true
		}
	}

	want := flag
	got := AnyNotPositive(numbers...)
	if want != got {
		t.Errorf("input: (%f)\ngot: %t\nwant %t\n",
			numbers, got, want,
		)
	}
}
