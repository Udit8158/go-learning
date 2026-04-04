package romannumerals

import (
	"fmt"
	"log"
	"testing"
	"testing/quick"
)

type TestCase struct {
	name           string
	input          uint16
	expectedOutput string
}

var testCases = []TestCase{
	{name: "1 converts to I", input: 1, expectedOutput: "I"},
	{name: "2 converts to II", input: 2, expectedOutput: "II"},
	{name: "3 converts to III", input: 3, expectedOutput: "III"},
	{name: "4 converts to IV", input: 4, expectedOutput: "IV"},
	{name: "5 converts to V", input: 5, expectedOutput: "V"},
	{name: "6 converts to VI", input: 6, expectedOutput: "VI"},
	{name: "7 converts to VII", input: 7, expectedOutput: "VII"},
	{name: "8 converts to VIII", input: 8, expectedOutput: "VIII"},
	{name: "9 converts to IX", input: 9, expectedOutput: "IX"},
	{name: "10 converts to X", input: 10, expectedOutput: "X"},
	{name: "14 converts to XIV", input: 14, expectedOutput: "XIV"},
	{name: "39 converts to XXIX", input: 39, expectedOutput: "XXXIX"},
	{name: "40 converts to XXIX", input: 40, expectedOutput: "XL"},
	{name: "47 converts to XLVII", input: 47, expectedOutput: "XLVII"},
	{name: "57 converts to LVII", input: 57, expectedOutput: "LVII"},
	{name: "89 converts to LXXIX", input: 89, expectedOutput: "LXXXIX"},
	{name: "94 converts to XCIV", input: 94, expectedOutput: "XCIV"},
	{name: "400 converts to CD", input: 400, expectedOutput: "CD"},
	{name: "1984 converts to MCMLXXXIV", input: 1984, expectedOutput: "MCMLXXXIV"},
	{name: "3999 converts to MMMCMXCIX", input: 3999, expectedOutput: "MMMCMXCIX"},
}

func TestConvertToRoman(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := ConvertToRoman(test.input)
			if got != test.expectedOutput {
				t.Errorf("expected %q got %q", test.expectedOutput, got)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, test := range testCases {
		t.Run(fmt.Sprintf("%q should converts into %d", test.expectedOutput, test.input), func(t *testing.T) {
			got := ConvertToArabic(test.expectedOutput)
			if got != test.input {
				t.Errorf("expected %d got %d", test.input, got)
			}
		})
	}
}

func TestConvertToArabicRec(t *testing.T) {
	for _, test := range testCases {
		t.Run(fmt.Sprintf("%q should converts into %d", test.expectedOutput, test.input), func(t *testing.T) {
			got := ConvertToArabicRec(test.expectedOutput)
			if got != test.input {
				t.Errorf("expected %d got %d", test.input, got)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			// log.Println(arabic)
			return true
		}
		log.Println(arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 10000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
