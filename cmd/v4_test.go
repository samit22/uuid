package cmd

import (
	"reflect"
	"testing"

	"github.com/samit22/uuid/generator"
)

type mockUUID struct {
}

func (m mockUUID) V4() (string, error) {
	return "0000000-0000-0000-0000-000000000000", nil
}

func TestParseArgs(t *testing.T) {
	tests := map[string]struct {
		args        []string
		expectedNum int
		expectErr   bool
	}{
		"no arguments": {
			args:        nil,
			expectedNum: 1,
			expectErr:   false,
		},
		"one argument 3": {
			args:        []string{"3"},
			expectedNum: 3,
			expectErr:   false,
		},
		"one argument 0": {
			args:        []string{"0"},
			expectedNum: 0,
			expectErr:   true,
		},
		"non-integer argument": {
			args:        []string{"foobar"},
			expectedNum: 0,
			expectErr:   true,
		},
		"superfluous arguments": {
			args:        []string{"23", "foo", "bar", "baz"},
			expectedNum: 23,
			expectErr:   false,
		},
	}

	for testName, testData := range tests {
		t.Run(testName, func(t *testing.T) {
			n, err := parseArgs(testData.args)
			if testData.expectErr != (err != nil) {
				t.Fatalf("Expected error doesn't match: expected error = %t, error returned = %v", testData.expectErr, err)
			}
			if testData.expectedNum != n {
				t.Fatalf("Expected number doesn't match: expected %d, got %d", testData.expectedNum, n)
			}
		})
	}
}

func TestGenerateUUIDv4(t *testing.T) {
	tests := map[string]struct {
		n             int
		gen           generator.GeneratorV4
		expectedUUIDs []string
	}{
		"Single UUID": {
			n:             1,
			gen:           mockUUID{},
			expectedUUIDs: []string{"0000000-0000-0000-0000-000000000000"},
		},
		"Multiple UUIDs": {
			n:   3,
			gen: mockUUID{},
			expectedUUIDs: []string{
				"0000000-0000-0000-0000-000000000000",
				"0000000-0000-0000-0000-000000000000",
				"0000000-0000-0000-0000-000000000000",
			},
		},
	}

	for testName, testData := range tests {
		t.Run(testName, func(t *testing.T) {
			uuids, err := generateUUIDV4(testData.n, testData.gen)
			if err != nil {
				t.Fatalf("generateUUIDV4 returned an error where none was expected: %v", err)
			}
			if !reflect.DeepEqual(testData.expectedUUIDs, uuids) {
				t.Fatalf("expected %v UUIDs to be returned, but generateUUIDV4 returned %v instead", testData.expectedUUIDs, uuids)
			}
		})
	}
}
