package stnchelper

import (
	"fmt"
	"regexp"
	"strings"
)

func findAllNumber(str string) string {
	// Regex Pattern:
	// -? -> Optional minus sign (for negative numbers)
	// \d+        -> One or more digits (0-9)
	// (\.\d+)? -> Optional decimal part (a dot followed by digits)
	re := regexp.MustCompile(`-?\d+(\.\d+)?`)

	// The FindAllString function finds all matches.
	// The -1 parameter indicates no limit, meaning it should return all matches.
	bulunanSayilar := re.FindAllString(str, -1)
	// result, _ := strconv.ParseInt(bulunanSayilar[0], 6, 12)
	//  fmt.Println("Bulunan Sayılar:", result)
	// fmt.Println("Bulunan Sayılar:", bulunanSayilar)
	return bulunanSayilar[0]
}

func findAllString(str string) string {
	var sb strings.Builder //https://www.perplexity.ai/search/bana-golang-struct-slice-icine-9DOZYn4MRFqz8bpg7pU.PQ#2e7b4d51-eeec-4c79-9487-c60ead9ef435

	// 1. YÖNTEM: Rakam olmayan (non-digit) her şeyi bul
	// [^0-9]+  -> 0-9 arası rakamlar HARİÇ her şey
	// \D+      -> (Alternatif kısa yazım) Digit olmayan her şey
	reNoDigits := regexp.MustCompile(`[^0-9]+`)

	matches := reNoDigits.FindAllString(str, -1)

	fmt.Println("--- Rakam Hariç Parçalar ---")
	for _, m := range matches {
		fmt.Printf("'%s'\n", m)
		sb.WriteString(m)
	}
	sonuc := strings.TrimSpace(sb.String())

	return sonuc
}
