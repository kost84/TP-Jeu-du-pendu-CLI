package src

import (
	"encoding/json"
	"os"
)

type Sauvegarde struct {
	Pseudo  string   `json:"pseudo"`
	Mot     string   `json:"mot"`
	ListeLettres []string `json:"lettres"`
	Essais  int      `json:"essais"`
}

func (s *Sauvegarde) Save(pseudo string, mot string, listelettres []string, essais int) {
	s.Pseudo = pseudo
	s.Mot = mot
	s.ListeLettres = listelettres
	s.Essais = essais
	data,_ := json.Marshal(s)
	os.WriteFile("src/save-score/save.json", data, 0644)
}

func ReprendreSave() (*Sauvegarde){
	var s Sauvegarde
	data,_ := os.ReadFile("src/save-score/save.json")
	json.Unmarshal(data, &s)
	return &s
}
