package src

import (
	"fmt"
	"strings"
)

var Saisipseudo string

func SaisiPseudo() {
	fmt.Print("Veuillez saisir votre Pseudonyme : ")
	fmt.Scanln(&Saisipseudo)
}

var MotSecret string

func MotSelec() {
	MotSecret = strings.ToLower("Essaie")
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

func ProposLettre() {
	fmt.Print("Lettre proposé : ")
	fmt.Scanln(&LettrePropos)
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
	for range LongMot {
		listemotProvi = append(listemotProvi, "_")
	}
	for a := range mot {
		if mot[a] == lettre[0] {
			listemotProvi[a] = string(lettre[0])
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
		fmt.Println("Perdu")
		Loose = true
	}
}

func StartGame() {
	SaisiPseudo()
	MotSelec()
	TailleMot(MotSecret)
	for Loose == false {
		fmt.Println(LongMot)
		ProposLettre()
		c := VerifLettre(LettrePropos)
		if c == false {
			fmt.Println("Erreur : Vous avez saisi plusieurs lettres ou aucune")
			continue
		} else {
			if BonneLettre(LettrePropos, MotSecret) == true {
				PlaceLettre(LettrePropos, MotSecret)
			} else {
				CheckEssais()
				if NbEssais != 0 {
					fmt.Println("Nombre d'essaie restant :", NbEssais)
				}
			}
		}
	}
}
