package romannumerals

import (
	"strings"
)

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"}, // Value, Symbol
}

func ConvertToRoman(arabicNum uint16) string {
	romanNumber := strings.Builder{}

	for _, v := range allRomanNumerals {
		for arabicNum >= v.Value {
			romanNumber.WriteString(v.Symbol)
			arabicNum -= v.Value
		}
	}

	return romanNumber.String()
}

func ConvertToArabic(romanNum string) uint16 {
	var arabicNum uint16 = 0

	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(romanNum, numeral.Symbol) {
			arabicNum += numeral.Value
			romanNum = strings.TrimPrefix(romanNum, numeral.Symbol)
		}
	}

	return arabicNum
}

func ConvertToArabicRec(romanNum string) uint16 {
	// f("XLVII") = 40 + f("VII")
	//           = 40 + 5 + f("II")
	//           = 40 + 5 + 1 + f("I")
	//           = 40 + 5 + 1 + 1 + f("")
	//           = 47

	for _, numeral := range allRomanNumerals {
		romanNum, ok := strings.CutPrefix(romanNum, numeral.Symbol)
		if ok {
			return numeral.Value + ConvertToArabicRec(romanNum)
		}
	}

	// f(XLVII) = f(X) + f(LVII) = 10 + ... WRONG
	// return ConvertToArabic(string(romanNum[0])) + ConvertToArabic(string(romanNum[1:]))
	// base case (f("") = 0)
	return 0

}
