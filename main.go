package main

import (
	"TP-Jeu-du-pendu-CLI/src"
	"fmt"
)

func main() {
	var saisipseudo string

	fmt.Println("Veuillez saisir votre Pseudonyme")
	fmt.Scanln(&saisipseudo)
	src.Save(saisipseudo, "test", []string{"a", "z"}, 5)
}
