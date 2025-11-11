package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

func matchAnagrams(array []string) map[string][]string {
	tempMap := make(map[string][]string, len(array))

	for _, word := range array {
		lower := strings.ToLower(word)
		sorted := sortString(lower)
		tempMap[sorted] = append(tempMap[sorted], lower)
	}

	result := make(map[string][]string, len(array))

	for _, array := range tempMap {
		if len(array) > 1 {
			sort.Strings(array)
			key := array[0]
			result[key] = array
		}
	}

	return result
}

func sortString(s string) string {
	r := []rune(s)
	slices.Sort(r)
	return string(r)
}

func main() {
	sl := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	fmt.Println(matchAnagrams(sl))

}
