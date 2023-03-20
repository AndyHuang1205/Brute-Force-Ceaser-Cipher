/**
Andy Huang and , Samiul Hassan
Brute Force Ceaser Cipher using golang
3/19/2023
*/

package main

import (
	"fmt"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
	ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func encrypt(n int, plaintext string) string {
	var result string
	for _, l := range plaintext {
		if strings.ContainsRune(ALPHABET, l) {
			index := strings.IndexRune(ALPHABET, l)
			newIndex := (index + n) % 26
			result += string(ALPHABET[newIndex])
		} else if strings.ContainsRune(alphabet, l) {
			index := strings.IndexRune(alphabet, l)
			newIndex := (index + n) % 26
			result += string(alphabet[newIndex])
		} else {
			result += string(l)
		}
	}
	return result
}
func bruteForce(ciphertext string) []string {
	result := make([]string, 26)
	for _, l := range ciphertext {
		if strings.ContainsRune(ALPHABET, l) {
			index := strings.IndexRune(ALPHABET, l)
			for i := 0; i < len(result); i++ {
				newIndex := (index - i + 26) % 26
				result[i] += string(ALPHABET[newIndex])
			}
		} else if strings.ContainsRune(alphabet, l) {
			index := strings.IndexRune(alphabet, l)
			for i := 0; i < len(result); i++ {
				newIndex := (index - i + 26) % 26
				result[i] += string(alphabet[newIndex])
			}
		} else {
			for i := 0; i < len(result); i++ {
				result[i] += string(l)
			}
		}
	}
	return result
}
func main() {
	message := "Come here Watson"
	key := 3
	cipherText := encrypt(key, message)
	for i, j := range bruteForce(cipherText) {
		fmt.Printf("%2d: %s\n", i, j)
	}
}
