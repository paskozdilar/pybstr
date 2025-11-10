package pybstr_test

import (
	"log"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"testing"

	pybstr "github.com/paskozdilar/go-python-bstring.git"
)

const Dir = "testdata"

func TestEncode(t *testing.T) {
	examplesMap := make(map[string]struct{})
	files, err := os.ReadDir(Dir)
	if err != nil {
		t.Fatalf("Failed to read testdata directory: %v", err)
	}
	for _, file := range files {
		// Get filename without ".bstr" or ".data" extension.
		name := file.Name()[:len(file.Name())-5]
		examplesMap[name] = struct{}{}
	}
	examples := slices.Sorted(maps.Keys(examplesMap))
	log.Println("Found examples:", examples)

	for _, example := range examples {
		t.Run(example, func(t *testing.T) {
			data, err := os.ReadFile(filepath.Join(Dir, example+".data"))
			if err != nil {
				t.Fatalf("Failed to read .data file for %q: %v", example, err)
			}
			bstr, err := os.ReadFile(filepath.Join(Dir, example+".bstr"))
			if err != nil {
				t.Fatalf("Failed to read .bstr file for %q: %v", example, err)
			}
			input := data
			expected := string(bstr)

			output := pybstr.Encode([]byte(input))
			if output != expected {
				// Get exact index of first difference for better debugging.
				idx := 0
				for i := range min(len(output), len(expected)) {
					if output[i] != expected[i] {
						break
					}
					idx++
				}
				t.Fatalf("Encode(%q) = %s; want %s; first difference at index %d", input, output, expected, idx)
			}
		})
	}
}

// func FuzzEncode(f *testing.F) {
// 	// Seed corpus with known edge cases
// 	f.Add("Hello\nWorld")
// 	f.Add("C:\\path\\file")
// 	f.Add("")
// 	f.Add("😀")        // Unicode
// 	f.Add("\x00\xFF") // raw bytes
//
// 	f.Fuzz(func(t *testing.T, input string) {
// 		output := Encode(input)
//
// 		// Basic sanity checks
// 		if len(output) < 3 || output[0] != 'b' || output[1] != '\'' || output[len(output)-1] != '\'' {
// 			t.Errorf("Invalid format for input %q: got %q", input, output)
// 		}
//
// 		// Optional: check that output is ASCII-safe
// 		for i, b := range output {
// 			if b < 0x20 && b != '\n' && b != '\t' && b != '\r' {
// 				t.Logf("Suspicious byte at %d: %q", i, b)
// 			}
// 		}
// 	})
// }
