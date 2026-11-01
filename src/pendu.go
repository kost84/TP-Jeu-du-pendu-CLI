package src

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

var Saisipseudo string

func SaisiPseudo() {
	fmt.Print("Veuillez saisir votre Pseudonyme : ")
	fmt.Scanln(&Saisipseudo)
}

var MotSecret string
var ListeMot []string

func MotSelec() {
	listemottempo, _ := os.ReadFile("mot.txt")
	var stringmottempo string = string(listemottempo)
	ListeMot = strings.Split(stringmottempo, "\n")
	indexrandom := rand.IntN(len(ListeMot))
	MotSecret = strings.ToLower(strings.TrimSpace(ListeMot[indexrandom]))
}

var LongMot string

func TailleMot(mot string) {
	var longMotTemp string = ""
	for range mot {
		longMotTemp += "_"
	}
	LongMot = longMotTemp
}

var LettrePropos string

func ProposLettre() bool {
	var lettreproposs string
	fmt.Print("Lettre proposé : ")
	fmt.Scanln(&lettreproposs)
	if len(ListeTester) == 0 {
		LettrePropos = lettreproposs
		ListeTester = append(ListeTester, lettreproposs)
		return true
	}
	for a := range ListeTester {
		if lettreproposs == ListeTester[a] {
			fmt.Println("Lettre déja proposé")
			lettreproposs = ""
			return false
		}
	}
	LettrePropos = lettreproposs
	ListeTester = append(ListeTester, lettreproposs)
	return true
}

func VerifLettre(lettre string) bool {
	if len(lettre) > 1 || len(lettre) < 1 {
		return false
	}
	return true
}

func BonneLettre(lettre string, mot string) bool {
	lettre = strings.ToLower(lettre)
	mot = strings.ToLower(mot)
	for indexMot := range mot {
		if mot[indexMot] == lettre[0] {
			return true
		}
	}
	return false
}

func PlaceLettre(lettre string, mot string) {
	var listemotProvi []string
	var longMotProvi string
	for c := range LongMot {
		listemotProvi = append(listemotProvi, string(LongMot[c]))
	}
	for a := range mot {
		if mot[a] == strings.ToLower(lettre)[0] {
			listemotProvi[a] = strings.ToLower(lettre)
		}
	}
	for b := range listemotProvi {
		longMotProvi += listemotProvi[b]
	}
	LongMot = longMotProvi
}

var NbEssais int = 5
var Loose bool = false

func CheckEssais() {
	if BonneLettre(LettrePropos, MotSecret) == false {
		NbEssais -= 1
	}
	if NbEssais == 0 {
		fmt.Print("\n\n")
		fmt.Println("Plus aucun essai restant.")
		fmt.Println("\nPerdu")
		Loose = true
	}
}

var ListeTester []string

func VerifWin() bool {
	if MotSecret == LongMot {
		return true
	}
	return false
}

func StartGame() {
	SaisiPseudo()
	MotSelec()
	TailleMot(MotSecret)
	for Loose == false {
		if VerifWin() {
			fmt.Print("\n\n")
			fmt.Println("Vous avez gagné")
			fmt.Println("\nLe mot secret était : ", MotSecret)
			fmt.Print("\n")
			break
		}
		fmt.Print("\n")
		fmt.Println(LongMot)
		fmt.Println("Nombre d'essaie restant :", NbEssais)
		fmt.Println("Lettre proposé : ", ListeTester)
		ProposLettre()
		c := VerifLettre(LettrePropos)
		if c == false {
			fmt.Println("Erreur : Vous avez saisi plusieurs lettres")
			continue
		} else {
			if BonneLettre(LettrePropos, MotSecret) == true {
				PlaceLettre(LettrePropos, MotSecret)
			} else {
				CheckEssais()
				if Loose == true {
					fmt.Print("\n")
				}
			}
		}
	}
}
