package main

import "testing"

func TestProcess(t *testing.T) {
	tests := []struct {
		entree  string
		attendu string
	}{
		// (hex) et (bin)
		{"1E (hex) files were added", "30 files were added"},
		{"It has been 10 (bin) years", "It has been 2 years"},
		{"Simply add 42 (hex) and 10 (bin) and you will see the result is 68.", "Simply add 66 and 2 and you will see the result is 68."},

		// (up), (low), (cap)
		{"Ready, set, go (up) !", "Ready, set, GO!"},
		{"I should stop SHOUTING (low)", "I should stop shouting"},
		{"Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},

		// (up, n), (low, n), (cap, n)
		{"This is so exciting (up, 2)", "This is SO EXCITING"},
		{"IT WAS THE (low, 3) winter", "it was the winter"},
		{"the age of foolishness (cap, 3)", "the Age Of Foolishness"},

		// Ponctuation
		{"I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"I was thinking ... You were right", "I was thinking... You were right"},
		{"Punctuation tests are ... kinda boring ,what do you think ?", "Punctuation tests are... kinda boring, what do you think?"},

		// Apostrophes
		{"I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},

		// a -> an
		{"There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"There is no greater agony than bearing a untold story inside you.", "There is no greater agony than bearing an untold story inside you."},

		// Exemple complet du sujet
		{"it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			"It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair."},

		// Cas limites : le programme ne doit pas planter
		{"(up) hello", "hello"},
		{"hi (up, 10)", "HI"},
		{"hello (hex)", "hello"},
		{"a", "a"},
		{"", ""},
	}

	for _, test := range tests {
		obtenu := Process(test.entree)
		if obtenu != test.attendu {
			t.Errorf("\nentrée  : %q\nobtenu  : %q\nattendu : %q", test.entree, obtenu, test.attendu)
		}
	}
}
