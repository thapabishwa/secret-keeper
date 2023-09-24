package helpers

import (
	"reflect"
	"testing"
)

var globTests = []struct {
	pattern string
	result  []string
	wantErr bool
}{
	{"testdata/*", nil, false},
	{"testdata/a", nil, false},
	{"match.go", nil, false},
	{"mat?h.go", nil, false},
	{"*", []string{"helpers.go", "helpers_test.go"}, false},
	{"*.go", []string{"helpers.go", "helpers_test.go"}, false},
	// bad pattern
	{"[", nil, false},
}

func TestFileList(t *testing.T) {
	for _, test := range globTests {
		got, err := FileList(test.pattern)
		if (err != nil) != test.wantErr {
			t.Errorf("FileList(%q) error = %v, wantErr %v", test.pattern, err, test.wantErr)
		}
		if !reflect.DeepEqual(got, test.result) {
			t.Errorf("FileList(%q) = %+#v, want %+#v", test.pattern, got, test.result)
		}
	}
}
