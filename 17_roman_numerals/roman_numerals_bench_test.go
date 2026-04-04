package romannumerals

import "testing"

func BenchmarkConvertToRoman(b *testing.B) {
	for b.Loop() {
		ConvertToRoman(3999)
	}
}

func BenchmarkConvertToArabic(b *testing.B) {
	for b.Loop() {
		ConvertToArabic("MMMCMXCIX")
	}
}

func BenchmarkConvertToArabicRec(b *testing.B) {
	for b.Loop() {
		ConvertToArabicRec("MMMCMXCIX")
	}
}
