package main

import (
	"strings"
)

func Process(input string) string {
	mots := strings.Fields(input) // découpe le texte en liste de mots

	mots = applyModifiers(mots)

	return strings.Join(mots, " ") // recolle les mots avec un espace
}

func applyModifiers(mots []string) []string {
	var result []string
	for _, mot := range mots {
		if mot == "(up)" {
			if len(result) > 0 {
				result[len(result)-1] = strings.ToUpper(result[len(result)-1])
			}
		} else if mot == "(low)" {
			if len(result) > 0 {
				result[len(result)-1] = strings.ToLower(result[len(result)-1])
			}
		} else {
			result = append(result, mot)
		}
	}
	return result
}