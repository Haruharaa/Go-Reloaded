package main

import (
	"regexp"
	"strconv"
	"strings"
)

var espaceAvant = regexp.MustCompile(`\s+([.,!?:;]+)`)
var espaceApres = regexp.MustCompile(`([.,!?:;]+)([^\s.,!?:;])`)

func Process(input string) string {
	mots := strings.Fields(input) // découpe le texte en liste de mots

	mots = applyModifiers(mots)

	texte := strings.Join(mots, " ") // recolle les mots avec un espace

	texte = fixPunctuation(texte) // corrige la ponctuation

	return texte
}

func applyModifiers(mots []string) []string {
	var result []string
	for i := 0; i < len(mots); i++ {
		mot := mots[i]
		if mot == "(up)" {
			if len(result) > 0 {
				result[len(result)-1] = strings.ToUpper(result[len(result)-1])
			}
		} else if mot == "(low)" {
			if len(result) > 0 {
				result[len(result)-1] = strings.ToLower(result[len(result)-1])
			}
		} else if mot == "(cap)" {
			if len(result) > 0 {
				result[len(result)-1] = capitalize(result[len(result)-1])
			}
		} else if mot == "(up," && i+1 < len(mots) {
			texteNombre := strings.TrimSuffix(mots[i+1], ")") // "2)" → "2"
			n, err := strconv.Atoi(texteNombre)               // "2"  → 2
			if err == nil {
				if n > len(result) {
					n = len(result) // protection : pas plus de mots qu'il n'y en a
				}
				for j := len(result) - n; j < len(result); j++ {
					result[j] = strings.ToUpper(result[j])
				}
			}
			i++ // saute le mot "2)"
		} else if mot == "(low," && i+1 < len(mots) {
			texteNombre := strings.TrimSuffix(mots[i+1], ")")
			n, err := strconv.Atoi(texteNombre)
			if err == nil {
				if n > len(result) {
					n = len(result)
				}
				for j := len(result) - n; j < len(result); j++ {
					result[j] = strings.ToLower(result[j])
				}
			}
			i++
		} else if mot == "(cap," && i+1 < len(mots) {
			texteNombre := strings.TrimSuffix(mots[i+1], ")")
			n, err := strconv.Atoi(texteNombre)
			if err == nil {
				if n > len(result) {
					n = len(result)
				}
				for j := len(result) - n; j < len(result); j++ {
					result[j] = capitalize(result[j])
				}
			}
			i++
		} else if mot == "(hex)" {
			if len(result) > 0 {
				n, err := strconv.ParseInt(result[len(result)-1], 16, 64)
				if err == nil {
					result[len(result)-1] = strconv.FormatInt(n, 10)
				}
			}
		} else if mot == "(bin)" {
			if len(result) > 0 {
				n, err := strconv.ParseInt(result[len(result)-1], 2, 64)
				if err == nil {
					result[len(result)-1] = strconv.FormatInt(n, 10)
				}
			}
		} else {
			result = append(result, mot)
		}
	}
	return result
}

func capitalize(mot string) string {
	if len(mot) == 0 {
		return mot
	}
	return strings.ToUpper(mot[:1]) + strings.ToLower(mot[1:])
}

func fixPunctuation(texte string) string {
	// Supprime les espaces avant la ponctuation 
	texte = espaceAvant.ReplaceAllString(texte, "$1")
	// Ajoute un espace après la ponctuation si besoin
	texte = espaceApres.ReplaceAllString(texte, "$1 $2")
	return texte
}
