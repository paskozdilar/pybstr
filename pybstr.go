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

func Decode(input string) ([]byte, error) {
	if len(input) < 3 || input[0] != 'b' || input[1] != '\'' || input[len(input)-1] != '\'' {
		return nil, strconv.ErrSyntax
	}

	content := input[2 : len(input)-1]
	var result []byte

	for i := 0; i < len(content); i++ {
		if content[i] == '\\' {
			if i+1 >= len(content) {
				return nil, strconv.ErrSyntax
			}
			i++
			switch content[i] {
			case 'n':
				result = append(result, '\n')
			case 'r':
				result = append(result, '\r')
			case 't':
				result = append(result, '\t')
			case '\\':
				result = append(result, '\\')
			case '\'':
				result = append(result, '\'')
			case 'x':
				if i+2 >= len(content) {
					return nil, strconv.ErrSyntax
				}
				hexStr := content[i+1 : i+3]
				b, err := strconv.ParseUint(hexStr, 16, 8)
				if err != nil {
					return nil, err
				}
				result = append(result, byte(b))
				i += 2
			default:
				return nil, strconv.ErrSyntax
			}
		} else {
			result = append(result, content[i])
		}
	}

	return result, nil
}
