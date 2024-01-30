package hangweb

import (
	"math/rand"
	"os"
	"bufio"
	"time"
)

// Fonction qui va créer le mot avec quelques lettres afficher
func CreateWord(word string) string {
	// On crée une liste de string
	wordtofind := []string{}
	// On boucle pour remplir cette liste de _
	for k := 0; k < len([]rune(word)); k++ {
		if word[k] == '-' {
			wordtofind = append(wordtofind, "-")
		} else {
			wordtofind = append(wordtofind, "_")
		}
	}
	// On boucle quelques fois pour afficher des lettres random
	for i := 0; i < (len([]rune(word))/2 - 1); i++ {
		// Génère un nombre random
		tempr := rand.Intn(len([]rune(word)))
		// Si à l'emplacement random se trouve un _ on rentre dan la boucle
		if wordtofind[tempr] == "_" {
			// On remplace le _ par la lettre correspondant à l'emplacement
			wordtofind[tempr] = string([]rune(word)[tempr])
		} else {
			i--
		}
	}
	myrep := ""
	// On boucle dans notre liste de string, pour la transformer en string
	for _, letter := range wordtofind {
		myrep += letter
	}
	// On retourne notre string
	return myrep
}

func RecupWord(fichier string)string{
	f, _ := os.OpenFile(fichier, os.O_RDWR, 0644)
	scanner := bufio.NewScanner(f)
	var wordlist []string
	for scanner.Scan() {
		wordlist = append(wordlist, scanner.Text())
	}
	rand.Seed(time.Now().UnixNano())
	randomword := wordlist[rand.Intn(len(wordlist))]
	return randomword
}


/*func Jouer(){
	for true {
		// Le joueur choisit une lettre ou un mot
		fmt.Print("Choose: ")
		fmt.Scanln(&choose)
		// On vérifie si ce qu'il a marqué est valide
		randomwordhide, state = hangweb.IsInputOk(choose, randomword, randomwordhide, &usedletter)
		// On clear le terminal
		//hangman_classic.Clear()
		// Si le joueur a déjà fait une erreur
		if essaie != 10 {
			// On affiche l'état du pendu
			fmt.Println(hang[9-essaie])
		}
		// Si ce que le joueur a marqué n'est pas valide
		if state == "fail" {
			// On diminue le nombre d'essai restant du joueur
			essaie--
			// On affiche l'état de la partie
			fmt.Println(hang[9-essaie])
			fmt.Printf("La lettre %v n'est pas comprise dans le mot, il ne te reste plus que : %v essais\n", choose, essaie)
			// Si la lettre a déjà été utilisé
		} else if state == "usedletter" {
			// On affiche le message correspondant
			fmt.Printf("Lettre déjà utilisée\n")
			// Si la lettre est valide
		} else if state == "good" {
			// On affiche le message correspondant
			fmt.Printf("La lettre %v est bien comprise dans le mot\n", choose)
			// Si le mot rentré n'est pas de la bonne taille
		} else if state == "wordinvalid" {
			// On affiche le message correspondant
			fmt.Printf("Le format n'est pas valide, veuillez rentrer une lettre ou un mot de bonne taille\n")
			// Si le mot est le bon
		} else if state == "wordgood" {
			// On affiche le message correspondant
			fmt.Printf("Tu as trouvé, il te restait %v essai(s), le mot est : %v", essaie, randomword)
			// On arrête le programme
			os.Exit(0)
			// Si l'input n'est pas une lettre
		} else if state == "error" {
			// On affiche le message correspondant
			fmt.Println("La lettre est invalide, veuillez recommencer")
			// Si la mot n'est pas le bon
		} else if state == "wordwrong" {
			// On retire 2 essais au lieu de 1
			essaie -= 2
			// On affiche le message correspondant
			fmt.Printf("Le mot proposé n'est pas le bon, il te reste %v essais\n", essaie)
		}
		// Si le joueur n'a plus d'essais
		if essaie <= 0 {
			// On clear le terminal
			hangweb.Clear()
			// On affiche l'état de la partie
			fmt.Print(hang[9])
			fmt.Printf("Tu as perdu, le mot était : %v", randomword)
			// On arrête le programme
			os.Exit(0)
		}
		fmt.Println(randomwordhide)
		// Si le mot a été totalement découvert
		if randomwordhide == randomword {
			// On clear le terminal
			hangweb.Clear()
			// On affiche le message correspondant
		}
	}
}*/