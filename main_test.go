// halo kak aku buat random kata sandi otomatis ya , tujuanya biar ga bingung milih kata sandi gitu
package main

import "testing"

type TestCase struct { //struct untuk menyimpan data testing
	length   int
	expected int
}

func TestGeneratePassword(t *testing.T) {
	var testCases = []TestCase{
		{length: 8, expected: 8},
		{length: 12, expected: 12},
		{length: 16, expected: 16},
	}

	for _, test := range testCases {
		result := randomPassword(test.length)
		if len(result) != test.expected {
			t.Errorf("Expected password length %d, but got %d", test.expected, len(result))
		}
	}
}
