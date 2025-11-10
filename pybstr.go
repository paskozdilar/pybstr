// Package pybstr provides functions to encode and decode Python-style byte
// strings in Go.
package pybstr

import (
	"strconv"
	"strings"
)

// Encode encodes a byte slice into a Python-style byte string representation.
func Encode(input []byte) string {
	builder := strings.Builder{}

	for _, b := range input {
		switch b {
		case '\n':
			builder.WriteString(`\n`)
		case '\r':
			builder.WriteString(`\r`)
		case '\t':
			builder.WriteString(`\t`)
		case '\\':
			builder.WriteString(`\\`)
		case '\'':
			builder.WriteString(`\'`)
		default:
			// Printable ASCII range
			if b >= 0x20 && b <= 0x7E {
				builder.WriteByte(b)
			} else {
				builder.WriteString("\\x")
				// Encode two-digit hex
				hex := strconv.FormatInt(int64(b), 16)
				if len(hex) == 1 {
					builder.WriteByte('0')
				}
				builder.WriteString(hex)
			}
		}
	}

	return "b'" + builder.String() + "'"
}

// func Decode(input string) ([]byte, error) {
// }
