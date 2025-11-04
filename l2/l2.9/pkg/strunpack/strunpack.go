package strunpack

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	// ErrorStartsWithDigit строка начинается с цифры
	ErrorStartsWithDigit  = errors.New("error: starts with digit")
	// ErrorInvalidCharacter в строке присутствует знак не соответсвующий буквам, цифрам, и \
	ErrorInvalidCharacter = errors.New("error: invalid character in string")
	// ErrorEndsWithSlash строка оканчивается на \
	ErrorEndsWithSlash    = errors.New("error: ends with slash")
)

// StrUnpack осуществляет примитивную распаковку строки, содержащей повторяющиеся символы/руны.
func StrUnpack(s string) (string, error) {
	if s == "" {
		return s, nil
	}

	// если строка начинается с цифры - возвращаем ошибку
	if unicode.IsDigit(rune(s[0])) {
		return "", ErrorStartsWithDigit
	}

	// если оканчивается слешем - возвращаем ошибку
	if rune(s[len(s)-1]) == rune('\\') {
		return "", ErrorEndsWithSlash
	}

	length := utf8.RuneCountInString(s)

	// letter хранит символы, которые будут добавляться в итоговую строку
	var letter rune

	// в digit будут собираться цифры для распаковки
	var digit strings.Builder

	// count - сколько раз надо добавить символ в итоговую строку
	// st, end - начало и конец диапазона с цифрами
	count := 0

	res := make([]string, 0, length)
	for ind, val := range s {
		// если текущая руна символ, сохраняем ее в letter
		if unicode.IsLetter(val) {
			letter = val
			// если встречаем \ - пропускаем
		} else if val == rune('\\') {
			continue

		} else if unicode.IsDigit(val) {
			// если встречаем цифру и до нее встречали \ - считаем ее символом
			if rune(s[ind-1]) == rune('\\') {
				letter = val
			} else {
				digit.WriteRune(val)
			}
		} else {
			return "", ErrorInvalidCharacter
		}

		if digit.Len() == 0 {
			count = 1
		} else if ind < len(s) - 1 && !unicode.IsDigit(rune(s[ind + 1])) || ind == len(s) - 1{
			count, _ = strconv.Atoi(digit.String())
			if count >= 1 {
				count--
			} else if count == 0 {
				res = res[:len(res)-1]
			}
			digit.Reset()
		}
		for range count {
			res = append(res, string(letter))
		}
		count = 0
	}
	return strings.Join(res, ""), nil
}
