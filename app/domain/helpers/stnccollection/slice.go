package stnccollection

import (
	"math"
	"slices"
	"strings"
)

// https://www.perplexity.ai/search/golang-slice-icindeki-boslukla-AiCVtEHcR1G5y1yQXXFOaw#0
func cleanEmptySlice(liste []string) []string {

	// liste := []string{"elma", "   ", "armut", "", "muz"}

	// Hem tamamen boş ("") hem de sadece boşluk içerenleri ("   ") siler
	liste = slices.DeleteFunc(liste, func(s string) bool {
		return strings.TrimSpace(s) == ""
	})

	// fmt.Printf("%q\n", liste) // Çıktı: ["elma" "armut" "muz"]
	return liste
}

// https://www.perplexity.ai/search/php-deki-array-chunk-fonksiyon-4QFCCf_nQOSPSn3qkuoOKQ
// TODO : versionan gore refactor edilmesi gerekiyor
func arrayChunkWithgo(items []string, slicesCount int) [][]string {

	// slices.Chunk bir "iterator" döndürür.
	// slices.Collect ile bunu doğrudan [][]int haline getirebilirsiniz.
	chunks := slices.Collect(slices.Chunk(items, slicesCount))

	// fmt.Println(chunks)
	return chunks
}

// https://golangcode.com/check-if-row-exists-in-slice/
// FindSlice elemnt
func FindSlice(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func FindSliceTypes(slice []string, val string) bool {

	for _, n := range slice {
		if val == n {
			return true
		}
	}
	return false

}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

// link https://stackoverflow.com/questions/18390266/how-can-we-truncate-float64-type-to-a-particular-precision
// ToFixedDecimal decimal format
func ToFixedDecimal(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
