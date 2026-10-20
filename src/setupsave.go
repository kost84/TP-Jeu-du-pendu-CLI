package src

type Sauvegarde struct {
	Pseudo  string
	Mot     string
	Lettres []string
	Essais  int
}

func (s *Sauvegarde) Save(pseudo string, mot string, lettres []string, essais int) {
	s.Pseudo = pseudo
	s.Mot = mot
	s.Lettres = lettres
	s.Essais = essais
}
