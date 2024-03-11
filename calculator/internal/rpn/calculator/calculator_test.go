package calculator

import (
	"os"
	"testing"
)

func TestAll(t *testing.T) {
	type test struct {
		name string
		test func(t *testing.T)
	}

	tests := []test{
		{name: "1k", test: Test_1k},
		{name: "10k", test: Test_10k},
		{name: "100k", test: Test_100k},
		{name: "1000k", test: Test_1000k},
		{name: "10000k", test: Test_10000k},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.test)
	}

}
func Test_1k(t *testing.T) {
	equation := loadTestdata_t(t, "../testdata/src/halimath/1k")
	got, err := Calculate(equation)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	want := float64(248253190541.03891)
	if got != want {
		t.Fatalf("got: %f, want %f", got, want)
	}
}

func Test_10k(t *testing.T) {
	equation := loadTestdata_t(t, "../testdata/src/halimath/10k")
	got, err := Calculate(equation)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	want := float64(214512826488689376)
	if got != want {
		t.Fatalf("got: %f, want %f", got, want)
	}
}

func Test_100k(t *testing.T) {
	equation := loadTestdata_t(t, "../testdata/src/halimath/100k")
	got, err := Calculate(equation)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	want := float64(481585283158357967896576)
	if got != want {
		t.Fatalf("got: %f, want %f", got, want)
	}
}

func Test_1000k(t *testing.T) {
	equation := loadTestdata_t(t, "../testdata/src/halimath/1m")
	got, err := Calculate(equation)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	want := float64(830415166156152287340265472)
	if got != want {
		t.Fatalf("got: %f, want %f", got, want)
	}
}

func Test_10000k(t *testing.T) {
	equation := loadTestdata_t(t, "../testdata/src/halimath/10m")
	got, err := Calculate(equation)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	want := float64(6780466519056739908798922614519103488)
	if got != want {
		t.Fatalf("got: %f, want %f", got, want)
	}
}

func loadTestdata_t(t *testing.T, file string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("error loading file %s: %s", file, err)
	}
	return string(content)
}
