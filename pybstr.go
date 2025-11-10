// Package pybstr provides functions to encode and decode Python-style byte
// strings.
package pybstr

import (
	"io"
	logg "log"
	"strconv"
	"strings"
)

var log = logg.New(io.Discard, "", 0)

// Encode encodes a byte slice into a Python-style byte string representation.
func Encode(input []byte) string {
	builder := strings.Builder{}

	for _, b := range input {
		switch b {
		case '\n':
			log.Println("Encoding newline", builder.String())
			builder.WriteString(`\n`)
		case '\r':
			log.Println("Encoding carriage return", builder.String())
			builder.WriteString(`\r`)
		case '\t':
			log.Println("Encoding tab", builder.String())
			builder.WriteString(`\t`)
		case '\\':
			log.Println("Encoding backslash", builder.String())
			builder.WriteString(`\\`)
		case '\'':
			log.Println("Encoding single quote", builder.String())
			builder.WriteString(`\'`)
		default:
			// Printable ASCII range
			if b >= 0x20 && b <= 0x7E {
				builder.WriteByte(b)
			} else {
				builder.WriteString("\\x")
				log.Println("Encoding prefix for hex escape", builder.String())
				// Encode two-digit hex
				hex := strconv.FormatInt(int64(b), 16)
				if len(hex) == 1 {
					builder.WriteByte('0')
					log.Println("Padding hex with leading zero: 0", builder.String())
				}
				builder.WriteString(hex)
				log.Println("Encoding byte as hex:", hex, builder.String())
			}
		}
	}

	return "b'" + builder.String() + "'"
}

// func Decode(input string) ([]byte, error) {
// }
