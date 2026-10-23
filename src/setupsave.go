package src

type Sauvegarde struct {
	Pseudo  string   `json:"pseudo"`
	Mot     string   `json:"mot"`
	Lettres []string `json:"lettres"`
	Essais  int      `json:"essais"`
}

func (s *Sauvegarde) Save(pseudo string, mot string, lettres []string, essais int) {
	s.Pseudo = pseudo
	s.Mot = mot
	s.Lettres = lettres
	s.Essais = essais
}
