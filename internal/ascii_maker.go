package internal

import (
	"os"
	"strings"
)

type AsciiArt struct {
	asciiChars map[rune][8]string
	banner     string
}

func NewAsciiArt(banner string) *AsciiArt {
	return &AsciiArt{
		asciiChars: make(map[rune][8]string),
		banner:     "internal/assets/fonts/" + banner + ".txt",
	}
}

func (a *AsciiArt) FillMap() error {
	text, err := os.ReadFile(a.banner)
	if err != nil {
		return err
	}
	textStr := strings.ReplaceAll(string(text), "\r", "")
	data := strings.Split(textStr, "\n")

	var currentChar rune = 32
	var currentCharRepresentation [8]string
	var count int
	for i, line := range data {
		if i == 0 || i == len(data)-1 {
			continue
		}
		if len(line) == 0 {
			count = 0
			currentChar++
			currentCharRepresentation = [8]string{}
			continue
		}
		currentCharRepresentation[count] = line
		count++
		if count == 8 {
			a.asciiChars[currentChar] = currentCharRepresentation
		}
	}
	return nil
}

func (a *AsciiArt) AsciiStringBuilder(s string) string {
	var builder strings.Builder
	inputStrInRune := []rune(s)
	for i := 0; i <= 7; i++ {
		for _, char := range inputStrInRune {
			if charRepresentation, found := a.asciiChars[char]; found {
				builder.WriteString(charRepresentation[i])
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (a *AsciiArt) AsciiMake(inputStr string) (string, int) {
	if !a.stringIsValid(inputStr) {
		return "", 400
	}
	if err := a.FillMap(); err != nil {
		return "", 500
	}

	var builder strings.Builder
	arr := strings.Split(inputStr, "\n")
	for _, word := range arr {
		if word == "" && len(arr) > 1 {
			builder.WriteString("\n")
			continue
		}
		builder.WriteString(a.AsciiStringBuilder(word))
	}
	return builder.String(), 0
}

func (a *AsciiArt) stringIsValid(str string) bool {
	strInRune := []rune(str)
	for _, char := range strInRune {
		if char < '\n' || char > '~' {
			return false
		}
	}
	return true
}
