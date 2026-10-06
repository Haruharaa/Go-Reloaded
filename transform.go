package main

import (
	"regexp"
	"strconv"
	"strings"
)

var espaceAvant = regexp.MustCompile(`\s+([.,!?:;]+)`)
var espaceApres = regexp.MustCompile(`([.,!?:;]+)([^\s.,!?:;])`)

func Process(input string) string {
	mots := strings.Fields(input)    // découpe le texte en liste de mots
	mots = applyModifiers(mots)      // applique (up), (hex), (cap, 2)...
	texte := strings.Join(mots, " ") // recolle les mots avec un espace
	texte = fixPunctuation(texte)    // corrige la ponctuation
	mots = strings.Fields(texte)     // redécoupe le texte corrigé
	mots = fixQuotes(mots)           // corrige les apostrophes
	mots = fixArticles(mots)         // corrige les articles "a" et "an"
	texte = strings.Join(mots, " ")  // recolle
	return texte
}

// applyModifiers parcourt les mots et applique les marqueurs
// (up), (low), (cap), (hex), (bin), (up, n), (low, n), (cap, n)
// aux mots qui les précèdent. Les marqueurs sont retirés du texte.
func applyModifiers(mots []string) []string {
	var result []string
	for i := 0; i < len(mots); i++ {
		mot := mots[i]

		switch {
		// Marqueur simple : (up), (low), (cap), (hex), (bin)
		case mot == "(up)" || mot == "(low)" || mot == "(cap)" || mot == "(hex)" || mot == "(bin)":
			nom := strings.Trim(mot, "()") // "(up)" → "up"
			applyToLastWords(result, nom, 1)

		// Marqueur avec un nombre : "(up," suivi de "2)"
		case (mot == "(up," || mot == "(low," || mot == "(cap,") && i+1 < len(mots):
			nom := strings.Trim(mot, "(,")                             // "(up," → "up"
			n, err := strconv.Atoi(strings.TrimSuffix(mots[i+1], ")")) // "2)" → 2
			if err == nil {
				applyToLastWords(result, nom, n)
			}
			i++ // saute le mot "2)"

		// Mot normal
		default:
			result = append(result, mot)
		}
	}
	return result
}

// applyToLastWords applique la transformation "nom" aux n derniers mots.
func applyToLastWords(mots []string, nom string, n int) {
	if n > len(mots) {
		n = len(mots) // protection : pas plus de mots qu'il n'y en a
	}
	for j := len(mots) - n; j < len(mots); j++ {
		mots[j] = transformWord(mots[j], nom)
	}
}

// transformWord applique une seule transformation à un mot.
func transformWord(mot string, nom string) string {
	switch nom {
	case "up":
		return strings.ToUpper(mot)
	case "low":
		return strings.ToLower(mot)
	case "cap":
		return capitalize(mot)
	case "hex":
		return convertBase(mot, 16)
	case "bin":
		return convertBase(mot, 2)
	}
	return mot
}

// convertBase convertit un nombre écrit dans une base (16 ou 2) en décimal.
// Si le mot n'est pas un nombre valide, il est renvoyé tel quel.
func convertBase(mot string, base int) string {
	n, err := strconv.ParseInt(mot, base, 64)
	if err != nil {
		return mot
	}
	return strconv.FormatInt(n, 10)
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

func fixQuotes(mots []string) []string {
	var result []string
	ouvert := false
	for i := 0; i < len(mots); i++ {
		mot := mots[i]
		if mot == "'" && !ouvert && i+1 < len(mots) {
			// apostrophe OUVRANTE : on la colle au mot suivant
			mots[i+1] = "'" + mots[i+1]
			ouvert = true
		} else if mot == "'" && ouvert && len(result) > 0 {
			// apostrophe FERMANTE : on la colle au mot précédent
			result[len(result)-1] = result[len(result)-1] + "'"
			ouvert = false
		} else {
			result = append(result, mot)
		}
	}
	return result
}

func fixArticles(mots []string) []string {
	for i := 0; i < len(mots)-1; i++ {
		if mots[i] == "a" || mots[i] == "A" {
			premiere := strings.ToLower(mots[i+1][:1]) // 1re lettre du mot suivant, en minuscule
			if strings.Contains("aeiouh", premiere) {
				mots[i] = mots[i] + "n"
			}
		}
	}
	return mots
}
