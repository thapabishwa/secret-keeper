package helpers

import (
	"reflect"
	"testing"

	log "github.com/sirupsen/logrus"
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
	{"../*/*.go", []string{"../commander/commander.go", "../commander/commander_test.go", "../config/config.go", "../config/config_test.go", "../helpers/helpers.go", "../helpers/helpers_test.go", "../vault/vault.go", "../vault/vault_test.go"}, false},
	// bad pattern
	{"[", nil, true},
}

func TestFileList(t *testing.T) {

	for _, test := range globTests {
		got, err := FileList(test.pattern, log.DebugLevel)
		if (err != nil) != test.wantErr {
			t.Errorf("FileList(%q) error = %v, wantErr %v", test.pattern, err, test.wantErr)
		}
		if !reflect.DeepEqual(got, test.result) {
			t.Errorf("FileList(%q) = %+#v, want %+#v", test.pattern, got, test.result)
		}
	}
}
