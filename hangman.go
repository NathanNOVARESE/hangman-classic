package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Sert a nétoyer le terminal
func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	// Charger les mots depuis le fichier "words.txt"
	words, err := loadWords("words.txt")
	if err != nil {
		fmt.Printf("Erreur lors du chargement des mots: %v\n", err)
		return
	}

	rand.Seed(time.Now().UnixNano())
	wordToGuess := selectRandomWord(words)
	attemptsLeft := 10
	revealedLetters := make([]bool, len(wordToGuess))
	// Calculer le nombre de lettres à révéler
	lettersToReveal := len(wordToGuess)/2 - 1
	// Révéler des lettres aléatoires du mot
	for i := 0; i < lettersToReveal; i++ {
		randomIndex := rand.Intn(len(wordToGuess))
		if !revealedLetters[randomIndex] {
			revealedLetters[randomIndex] = true
		}
	}
	_error := 0
	guesses := make([]string, 0)

	// Affichage du pendue
	for attemptsLeft > 0 {
		displayWord(wordToGuess, revealedLetters)
		printJose(_error)
		fmt.Printf("Tentatives restantes: %d\n", attemptsLeft)
		fmt.Printf("Lettres donnée: %s\n", strings.Join(guesses, ", "))
		fmt.Print("Donner une lettre: ")
		var letter string
		fmt.Scanln(&letter)
		clearScreen()

		// Vérifie qu'une seule lettre est entrée
		if len(letter) != 1 || !isLetter(letter) {
			fmt.Println("Entrée non valide. Veuillez entrer une seule letre.")
			continue
		}

		if strings.Contains(wordToGuess, letter) {
			for i, char := range wordToGuess {
				if string(char) == letter {
					revealedLetters[i] = true
				}
			}

			// Affichage du pendue dans le terminal
		} else {
			fmt.Printf("La lettre '%s' n'est pas dans le mot.\n", letter)
			attemptsLeft--
			_error++
		}

		guesses = append(guesses, letter)

		if isWordGuessed(revealedLetters) {
			fmt.Printf("Bravo, vous avez devinés le mot suivant: %s", wordToGuess)
			break
		}
	}

	if attemptsLeft == 0 {
		fmt.Printf("Vous avez épuisé toute vos tentative. Le mot était: %s\n", wordToGuess)
		printJose(_error)
	}
}

// Charge les mots du fichier dans une slice, puis supprime les espaces inutiles
func loadWords(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	words := strings.Fields(strings.TrimSpace(string(content))) // Utilise la commande "strings.Fields" pour diviser les mots
	return words, nil
}

func loadhangman(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	words := strings.Fields(strings.TrimSpace(string(content))) // Utilise la commande "strings.Fields" pour diviser les mots
	return words, nil
}

func printJose(_error int) {
	f, err := os.Open("hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	cpt := 0
	cptEnd := 0
	switch _error {
	case 0:
		cpt = 0
		cptEnd = 0
	case 1:
		cpt = 0
		cptEnd = 8
	case 2:
		cpt = 9
		cptEnd = 16
	case 3:
		cpt = 16
		cptEnd = 24
	case 4:
		cpt = 24
		cptEnd = 32
	case 5:
		cpt = 32
		cptEnd = 40
	case 6:
		cpt = 40
		cptEnd = 48
	case 7:
		cpt = 48
		cptEnd = 56
	case 8:
		cpt = 56
		cptEnd = 64
	case 9:
		cpt = 64
		cptEnd = 72
	case 10:
		cpt = 72
		cptEnd = 80
	}
	i := 0
	for scanner.Scan() {
		if i >= cpt && i < cptEnd {
			fmt.Println(scanner.Text())
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// Choisis un mot aléatoire présent dans le fichier "words.txt"
func selectRandomWord(words []string) string {
	randIndex := rand.Intn(len(words))
	return words[randIndex]
}

// Sert à remplacer les lettres par "_"
func displayWord(word string, revealed []bool) {
	display := ""
	for i, char := range word {
		if revealed[i] {
			display += string(char)
		} else {
			display += "_"
		}
	}
	fmt.Println(display)
}

// Vérifie que c'est bien une lettre qui est entrée
func isLetter(s string) bool {
	return len(s) == 1 && ((s[0] >= 'a' && s[0] <= 'z') || (s[0] >= 'A' && s[0] <= 'Z'))
}

// Vérifie si le mot est deviné et permet de mettre fin a la partie
func isWordGuessed(revealed []bool) bool {
	for _, revealedLetter := range revealed {
		if !revealedLetter {
			return false
		}
	}
	return true
}
