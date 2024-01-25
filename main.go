package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	reader *bufio.Reader
)

// funcao para realizar a validacao da keyWord
func ValidateKeyWord(keyWord string) error {

	if len(keyWord) < 5 {
		return fmt.Errorf("the key must contain more than 5 letters")
	}

	seenLetters := make(map[byte]struct{})

	for _, char := range keyWord {

		if string(char) == " " || string(char) == "" {
			return fmt.Errorf("a blank character found in the keyword")
		}

		if _, found := seenLetters[byte(char)]; found {
			return fmt.Errorf("a duplicated letter found in the keyword")
		}

		seenLetters[byte(char)] = struct{}{}

	}

	return nil
}

// Funcao para encriptar o texto com base em uma chave
func Encrypt(keyWord, phrase string) {

	numRows := int((float64(len(phrase) / len(keyWord)))) + 1

	matrix := make([][]byte, numRows)
	encryptedMatrix := make([][]byte, numRows)
	for i := range matrix {
		matrix[i] = make([]byte, len(keyWord))
		encryptedMatrix[i] = make([]byte, len(keyWord))

	}

	// Ordena uma string
	strSlice := strings.Split(keyWord, "")

	sort.Strings(strSlice)

	index := 0
	for i := range matrix {
		for j := range strSlice {
			if index < len(phrase) {
				matrix[i][j] = phrase[index]
				index++
			} else {
				// Preencher com um caractere fictício se a frase acabar
				matrix[i][j] = ' '
			}
		}
	}

	for i := range matrix {
		for j := range keyWord {
			originalIndex := strings.Index(strings.Join(strSlice, ""), string(keyWord[j]))
			encryptedMatrix[i][j] = matrix[i][originalIndex]
		}
	}

	// Juntar toda a matriz em uma string
	var resultBuilder strings.Builder
	for i := range encryptedMatrix {
		resultBuilder.Write(encryptedMatrix[i])
	}

	result := resultBuilder.String()

	fmt.Println("Fase original:", phrase)
	fmt.Println("Resultado:", result)

}

func main() {

	reader = bufio.NewReader(os.Stdin)

	fmt.Print("Enter the password to perform encryption: ")
	keyWord, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("error on read string: %v\n", err.Error())
		return
	}

	if err = ValidateKeyWord(strings.ToLower(keyWord)); err != nil {
		fmt.Printf("error: %v\n", err.Error())
		return
	}

	fmt.Print("Enter the phrase to encrypt: ")

	phrase, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("error on read string: %v\n", err.Error())
		return
	}

	phrase = strings.Trim(phrase, "\n")

	Encrypt(keyWord, phrase)

}
