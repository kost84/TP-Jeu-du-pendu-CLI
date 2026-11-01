package main

import (
	"TP-Jeu-du-pendu-CLI/src"
	"os"
)

func main() {
	if len(os.Args)>1 {
		if os.Args[1] == "jouer"{
			src.StartGame(false)
		}else if os.Args[1] == "reprendre"{
			src.StartGame(true)
		}else if os.Args[1] == "score"{
			src.StartGame(false)
		}
	}else {
		src.StartGame(false)
	}
}
