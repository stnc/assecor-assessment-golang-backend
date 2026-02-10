package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Person struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	ZipCode  string `json:"zipcode"`
	City     string `json:"city"`
	Color    string `json:"color"`
}

func main() {

	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// DebugType(reader)
	var persons []Person
	var newData []string
	// var mydata []string
	lineNum := 0
	var headerCols int

	record, err := reader.Read()

	headerCols = len(record)

	fmt.Printf("Header kolon sayısı: %d (ilk satır: %v)\n", headerCols, record)

	for {
		wholeData, err := reader.Read()
		// if err != nil {
		// 	log.Fatal("Header okunamadı:", err)
		// }
		// DebugType(record)
		wholeData = cleanEmptySlice(wholeData)
		// fmt.Println(len(record))

		for i := 0; i < len(wholeData); i++ {
			newData = append(newData, wholeData[i])
		}

		// fmt.Println(wholeData)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Satır %d okuma hatası: %v", lineNum+1, err)
			continue
		}
	}
	// fmt.Println(newData)
	// fmt.Println(len(newData))

	stnc := arrayChunkWithgo(newData, headerCols)
	csvTOjson(&persons, stnc)
	// fmt.Printf("%+v\n", persons)
	// makeSlice()

	guzelJson, err := json.MarshalIndent(persons, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(guzelJson))

}

func colorConventor(color string) string {

	color = strings.TrimSpace(color)
	color2, _ := strconv.Atoi(color)

	switch color2 {
	case 1:
		return "blau"
	case 2:
		return "grün"
	case 3:
		return "violett"
	case 4:
		return "rot"
	case 5:
		return "gelb"
	case 6:
		return "türkis"
	case 7:
		return "weiß"
	default:
		return "violett"
	}
}
func csvTOjson(persons *[]Person, recordData [][]string) {
	emptyColsCheck := len(recordData)
	fmt.Println(emptyColsCheck)
	// var row int = 10
	for i, row := range recordData {
		extractZipCode := row[2]
		extractZipCode = findAllNumber(row[2])
		city := findAllString(row[2])
		i++
		p := Person{
			Name:     strings.TrimSpace(row[0]),
			City:     city,
			LastName: strings.TrimSpace(row[1]),
			ZipCode:  extractZipCode,
			Color:    colorConventor(row[3]),
			ID:       i,
		}
		*persons = append(*persons, p)
	}
	// kullanicilar := make([]Kullanici, 5) // 5 adet boş struct oluşturur
	// kullanicilar[0] = Kullanici{Ad: "Ece", Yas: 33}
}

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

func DebugType(v any) {
	t := reflect.TypeOf(v)
	fmt.Println("Type:", t.String())
	fmt.Println("Kind:", t.Kind())
}
