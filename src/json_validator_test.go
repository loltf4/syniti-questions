package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validate(t *testing.T) {
	type testCase struct {
		name string
		inputData []Data

		expectedOutput []string
	}
	cases := []testCase{
		{
			name: "Should print nothing when all data is valid",
			inputData: []Data{{Name: "abc", Address: "def", Zip: "12345", ID: "abc123"}},
			expectedOutput: []string(nil),
		},
		{
			name: "Should print invalid ID when name is invalid",
			inputData: []Data{{Name: "", Address: "def", Zip: "12345", ID: "abc123"}},
			expectedOutput: []string{"abc123"},
		},
		{
			name: "Should print invalid ID when address is invalid",
			inputData: []Data{{Name: "abc", Address: "", Zip: "12345", ID: "abc123"}},
			expectedOutput: []string{"abc123"},
		},
		{
			name: "Should print all invalid IDs when zip codes are invalid",
			inputData: []Data{{Name: "abc", Address: "def", Zip: "1234", ID: "abc123"},{Name: "abc", Address: "def", Zip: "99951", ID: "def456"},{Name: "abc", Address: "def", Zip: "ABCDE", ID: "ghi789"},{Name: "zzz", Address: "zzz", Zip: "54321", ID: "zzz000"}},
			expectedOutput: []string{"abc123", "def456", "ghi789"},
		},
		{
			name: "Should print all dupe IDs when dupes are present",
			inputData: []Data{{Name: "abc", Address: "def", Zip: "12345", ID: "abc123"},{Name: "abc", Address: "def", Zip: "12345", ID: "def456"},{Name: "abc", Address: "def", Zip: "12345", ID: "ghi789"},{Name: "zzz", Address: "zzz", Zip: "54321", ID: "zzz000"}},
			expectedOutput: []string{"abc123", "def456", "ghi789"},
		},
		{
			name: "Should print invalid ids of dupes when all data is duplicate (excluding id)",
			inputData: []Data{{Name: "abc", Address: "abc", Zip: "12345", ID: "abc123"},{Name: "abc", Address: "def", Zip: "12345", ID: "def456"},{Name: "abc", Address: "abc", Zip: "54321", ID: "ghi789"},{Name: "abc", Address: "abc", Zip: "12345", ID: "zzz000"}},
			expectedOutput: []string{"abc123", "zzz000"},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := validate(c.inputData)

			assert.Equal(t, c.expectedOutput, out)
		})
	}
}