package utils

import (
	"bytes"
	"regexp"
	"strings"
	"unicode"
)

const (
	space = " "
)

// IsEmpty returns true if the string is empty
func IsEmpty(text string) bool {
	return len(text) == 0
}

// IsNotEmpty returns true if the string is not empty
func IsNotEmpty(text string) bool {
	return !IsEmpty(text)
}

// IsBlank returns true if the string is blank (all whitespace)
func IsBlank(text string) bool {
	return len(strings.TrimSpace(text)) == 0
}

// IsNotBlank returns true if the string is not blank
func IsNotBlank(text string) bool {
	return !IsBlank(text)
}

func Left(src string, size int) string {
	return src[:size]
}

func Right(src string, size int) string {
	return src[len(src)-size:]
}

// Left justifies the text to the left
func LeftJustin(text string, size int) string {
	spaces := size - Length(text)
	if spaces <= 0 {
		return text
	}

	var buffer bytes.Buffer
	buffer.WriteString(text)

	for i := 0; i < spaces; i++ {
		buffer.WriteString(space)
	}
	return buffer.String()
}

// Right justifies the text to the right
func RightJustin(text string, size int) string {
	spaces := size - Length(text)
	if spaces <= 0 {
		return text
	}

	var buffer bytes.Buffer
	for i := 0; i < spaces; i++ {
		buffer.WriteString(space)
	}

	buffer.WriteString(text)
	return buffer.String()
}

// Center justifies the text in the center
func CenterJustin(text string, size int) string {
	left := RightJustin(text, (Length(text)+size)/2)
	return LeftJustin(left, size)
}

// Length counts the input while respecting UTF8 encoding and combined characters
func Length(text string) int {
	textRunes := []rune(text)
	textRunesLength := len(textRunes)

	sum, i, j := 0, 0, 0
	for i < textRunesLength && j < textRunesLength {
		j = i + 1
		for j < textRunesLength && IsMark(textRunes[j]) {
			j++
		}
		sum++
		i = j
	}
	return sum
}

// IsMark determines whether the rune is a marker
func IsMark(r rune) bool {
	return unicode.Is(unicode.Mn, r) || unicode.Is(unicode.Me, r) || unicode.Is(unicode.Mc, r)
}

// AddStringBytes 拼接字符串, 返回 bytes from bytes.Join()
func AddStringBytes(s ...string) []byte {
	switch len(s) {
	case 0:
		return []byte{}
	case 1:
		return []byte(s[0])
	}

	n := 0
	for _, v := range s {
		n += len(v)
	}

	b := make([]byte, n)
	bp := copy(b, s[0])
	for _, v := range s[1:] {
		bp += copy(b[bp:], v)
	}

	return b
}

// AddString 拼接字符串
func AddString(s ...string) string {
	return string(AddStringBytes(s...))
}

func ReplacePunctuationWithSpace(src string) string {
	reg1 := regexp.MustCompile(`[\f\t\n\r\v\-\^\$\.\*+\?{}()\/\[\]\|]`)
	reg2 := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	txt := reg2.ReplaceAll(reg1.ReplaceAll([]byte(src), []byte(" ")), []byte(" "))
	return string(txt)
}

func AddSpaceBetweenCharsAndNumbers(src string) string {
	var result []rune
	isNumber := false
	isChinese := false
	reg, _ := regexp.Compile("[^\\w\u4e00-\u9fa5]")
	src = reg.ReplaceAllString(src, " ")
	for i, s := range []rune(src) {
		if s >= '0' && s <= '9' {
			if !isNumber && i > 0 {
				result = append(result, ' ')
			}
			isNumber = true
		} else if isNumber {
			result = append(result, ' ')
			isNumber = false
		}
		if (unicode.Is(unicode.Han, s) && !isChinese) || (!unicode.Is(unicode.Han, s) && isChinese) {
			result = append(result, ' ')
			isChinese = !isChinese
		}
		result = append(result, s)
	}
	reg = regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	return strings.TrimSpace(string(reg.ReplaceAll([]byte(string(result)), []byte(" "))))
}


func ReplacePunctuation(src, replaceWith string) string {
	reg, _ := regexp.Compile("[^\\w\u4e00-\u9fa5]*")
	return reg.ReplaceAllString(src, replaceWith)
}
