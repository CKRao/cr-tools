package word

import (
	"strings"
	"unicode"
)

//ToUpper make str to upper
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

//ToLower make str to lower
func ToLower(s string) string {
	return strings.ToLower(s)
}

//UnderscoreToUpperCamelCase make underscore str to upperCamelCase str
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

//UnderscoreToLowerCamelCase make underscore str to lowerCamelCase str
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

//CamelCaseToUnderscore make CamelCase str to underscore str
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}

		if unicode.IsUpper(r) {
			output = append(output, '_')
		}

		output = append(output, unicode.ToLower(r))
	}

	return string(output)
}
