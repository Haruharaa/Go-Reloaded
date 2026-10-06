# Go-Reloaded

Outil en ligne de commande, écrit en Go, qui corrige et met en forme un texte.
Il lit un fichier d'entrée, applique une série de règles, puis écrit le résultat dans un fichier de sortie.

## Utilisation

```bash
go run . <fichier_entree> <fichier_sortie>
```

Exemple :

```bash
$ cat sample.txt
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
$ go run . sample.txt result.txt
$ cat result.txt
Simply add 66 and 2 and you will see the result is 68.
```

## Règles appliquées

| Règle | Exemple d'entrée | Résultat |
|---|---|---|
| `(hex)` : hexadécimal → décimal | `1E (hex) files` | `30 files` |
| `(bin)` : binaire → décimal | `10 (bin) years` | `2 years` |
| `(up)` : majuscules | `go (up)` | `GO` |
| `(low)` : minuscules | `SHOUTING (low)` | `shouting` |
| `(cap)` : majuscule initiale | `bridge (cap)` | `Bridge` |
| `(up, n)`, `(low, n)`, `(cap, n)` : sur les n mots précédents | `so exciting (up, 2)` | `SO EXCITING` |
| Ponctuation `. , ! ? : ;` collée au mot précédent | `there ,and` | `there, and` |
| Groupes de ponctuation (`...`, `!?`) | `thinking ... You` | `thinking... You` |
| Apostrophes `' '` collées aux mots qu'elles entourent | `' awesome '` | `'awesome'` |
| `a` → `an` devant une voyelle ou un `h` | `a untold story` | `an untold story` |

## Structure du projet

```
Go-Reloaded/
├── go.mod              # déclaration du module Go
├── main.go             # lecture des arguments et des fichiers, écriture du résultat
├── transform.go        # toutes les transformations du texte
└── transform_test.go   # tests unitaires
```

`Process` (dans `transform.go`) applique les étapes dans cet ordre :

1. découpage du texte en mots ;
2. marqueurs `(up)`, `(hex)`, `(cap, 2)`… ;
3. ponctuation ;
4. apostrophes ;
5. `a` → `an`.

L'ordre compte : les marqueurs sont traités avant la ponctuation pour qu'un signe comme `!` ne se colle pas à un marqueur (`(up)!`).

## Tests

```bash
go test -v
```

Les tests couvrent tous les exemples du sujet ainsi que des cas limites (marqueur en début de texte, nombre trop grand, faux nombre hexadécimal, texte vide).

## Packages utilisés

Uniquement la bibliothèque standard de Go : `os`, `fmt`, `strings`, `strconv`, `regexp`, `testing`.
