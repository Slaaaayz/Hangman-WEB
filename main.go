package main

import (
	"encoding/csv"
	"fmt"
	hw "hangweb/projets"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	// Handler de la page d'accueil
	game := Game{ //On initialise l'objet game qui va nous permettre de gérer le jeu
		Guess:   "",
		Vie:     10,
		Message: "",
		Pseudo:  "",
	}

	// HANDLER DE TOUTES LES PAGES
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		game = homeHandler(w, r, game)
	})
	http.HandleFunc("/recreate", func(w http.ResponseWriter, r *http.Request) {
		game = recreateHandler(w, r, game)
	})
	http.HandleFunc("/jouer", func(w http.ResponseWriter, r *http.Request) {
		game = jouerHandler(w, r, game)
	})
	http.HandleFunc("/echec", func(w http.ResponseWriter, r *http.Request) {
		echecHandler(w, r, game)
	})
	http.HandleFunc("/victoire", func(w http.ResponseWriter, r *http.Request) {
		victoireHandler(w, r, game)
	})

	// Pour ces 2 handlers on a pas besoin de l'objet game
	http.HandleFunc("/rules", rulesHandler)
	http.HandleFunc("/scoreboard", scoreboardHandler)

	// Démarrer le serveur sur le port 8080
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}
}

type Game struct { // Notre objet Game
	Wordc      string //le mot caché
	Word       string
	Guess      string // le mot/la lettre tenté par le joueur
	Vie        int
	Tabletters []string //Tableau des lettres déjà utilisées
	Message    string
	Pseudo     string
	Level      string // fichier de mots utilisé
	NameLevel  string //Nom du niveau (easy,medium,hard)
}

type scoreData struct { //Objet qui stock les données du scoreboard
	Name  string
	Level string
	Score int
}

func homeHandler(w http.ResponseWriter, r *http.Request, game Game) Game { // PAGE D'ACCUEIL
	game.Level = r.FormValue("level")
	tmpl, err := template.ParseFiles("./templates/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Erreur de chargement oupsi")
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return game
}

func rulesHandler(w http.ResponseWriter, r *http.Request) { // PAGE DES REGLES DU JEU
	tmpl, err := template.ParseFiles("./templates/rules.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("erreur de chargement")
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func jouerHandler(w http.ResponseWriter, r *http.Request, game Game) Game { //PAGE DU JEU
	game.Guess = r.FormValue("entry") // On récupère l'essai du joueur
	etat := ""
	game.Wordc, etat, game.Tabletters = hw.IsInputOk(game.Guess, game.Word, game.Wordc, game.Tabletters) // On le vérifie
	if game.Guess == "mentors" {
		http.Redirect(w, r, "https://www.youtube.com/watch?v=dQw4w9WgXcQ", http.StatusSeeOther)
	} else if game.Guess == "" {
		game.Message = ""
	} else {
		switch etat { // et en fonction du résultats on affiche le message correspondant
		case "wordgood":
			game.Message = "You won"
			http.Redirect(w, r, "/victoire", http.StatusSeeOther)
		case "wordwrong":
			game.Message = "Nice try but no"
			game.Vie -= 2 //si on tente un mot mais que c'est faux -2 vies
		case "wordinvalid":
			game.Message = "Invalid intput" // Si on a tenté un mot dont la longueur est différente que celle du mot à trouver on considère que c'est une faute de frappe donc on ne retire pas de vie
		case "usedletter":
			game.Message = "Already used !"
		case "fail":
			game.Message = "Wrong letter"
			game.Vie -= 1 //Si on tente une mauvaise lettre -1 vie
		case "good":
			game.Message = "You found a letter !"
		case "error":
			game.Message = "character error"
		}

		if game.Wordc == game.Word { // Si le joueur a gagné
			game.Message = "win"
			http.Redirect(w, r, "/victoire", http.StatusSeeOther) //On le redirige sur la page de victoire
		}
		if game.Vie <= 0 { // Si il a perdu
			game.Message = "loose"
			game.Vie = 0
			http.Redirect(w, r, "/echec", http.StatusSeeOther) // On le redirige sur la page d'echec
		}
	}

	tmpl, err := template.ParseFiles("./templates/jouer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = tmpl.Execute(w, game) //on prend game au moment de l'execution pour pouvoir afficher ses attributs
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return game // On renvoie notre objet game après l'avoir update
}

func echecHandler(w http.ResponseWriter, r *http.Request, game Game) { // PAGE DE GAMEOVER
	tmpl, err := template.ParseFiles("./templates/echec.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func victoireHandler(w http.ResponseWriter, r *http.Request, game Game) { // PAGE DE VICTOIRE
	tmpl, err := template.ParseFiles("./templates/victoire.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// On enregistre le joueur dans le scoreboard

	file, err := os.OpenFile("scoreboard.csv", os.O_CREATE|os.O_APPEND, 0600)
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	wr := csv.NewWriter(file)
	defer wr.Flush()
	row := []string{game.Pseudo, game.NameLevel, strconv.Itoa(game.Vie + 10)}
	if err := wr.Write(row); err != nil {
		log.Fatalln("error writing record to file", err)
	}
}

func recreateHandler(w http.ResponseWriter, r *http.Request, game Game) Game { // Page qui reset la game
	game.Level = r.FormValue("level")
	switch game.Level {
	case "words.txt":
		game.NameLevel = "Easy"
	case "words2.txt":
		game.NameLevel = "Medium"
	case "words3.txt":
		game.NameLevel = "Hard"
	}
	game.Pseudo = r.FormValue("pseudo")
	randomword := hw.RecupWord("projets/text/" + game.Level)
	game.Wordc = hw.CreateWord(randomword)
	game.Word = randomword
	game.Tabletters = nil
	game.Vie = 10
	jouerHandler(w, r, game)
	return game
}

func scoreboardHandler(w http.ResponseWriter, r *http.Request) { // PAGE DU SCOREBOARD
	tmpl, err := template.ParseFiles("./templates/scoreboard.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, err := os.Open("scoreboard.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatalln("error reading the file", err)
	}

	// On vient trier les données du scoreboard par niveau puis on stock tout ces objets dans un tableau

	var ScoresEasy []scoreData
	var ScoresMedium []scoreData
	var ScoresHard []scoreData
	for _, row := range rows {
		datascore, _ := strconv.Atoi(row[2])
		data := scoreData{
			Name:  row[0],
			Level: row[1],
			Score: datascore,
		}
		switch data.Level {
		case "Easy":
			ScoresEasy = append(ScoresEasy, data)
		case "Medium":
			ScoresMedium = append(ScoresMedium, data)
		case "Hard":
			ScoresHard = append(ScoresHard, data)
		}
	}
	ScoresEasy = TriScore(ScoresEasy)
	ScoresMedium = TriScore(ScoresMedium)
	ScoresHard = TriScore(ScoresHard)
	var Scores [][]scoreData
	Scores = append(Scores, ScoresEasy)
	Scores = append(Scores, ScoresMedium)
	Scores = append(Scores, ScoresHard)
	err = tmpl.Execute(w, Scores)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func TriScore(Scores []scoreData) []scoreData { // On met les meilleurs scores en haut du scoreboard
	i := 1
	for i < len(Scores) {
		if Scores[i].Score > Scores[i-1].Score {
			safet := Scores[i]
			Scores[i] = Scores[i-1]
			Scores[i-1] = safet
			i = 1
		} else {
			i++
		}
	}
	return Scores
}
