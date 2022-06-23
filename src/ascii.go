package asciiart

import (
	"fmt"
	"strings"
)

func StringToAsciiArt(str string, template map[rune]string) (string, error) {
	result := ""

	strs := strings.Split(str, "\r\n")
	for _, l := range strs {
		if l == "" {
			result += "\n"
		}
		lines := make([]string, 8)
		if len(l) == 0 {
			continue
		}
		for i := 0; i < len(l); i++ {
			if !(l[i] > 31 && l[i] < 127 || l[i] == 10) {
				return "", fmt.Errorf("only ascii chars, please")
			}
			v := 8
			for j := 0; j < v; j++ {
				lines[j] += strings.Split(template[rune(l[i])], "\n")[j]
			}
		}
		for _, k := range lines {
			result += k + "\n"
		}
	}
	return result, nil
}

func ParseTemplateToMap(path string) (map[rune]string, error) {
	text, err := ReadFile(path)
	if err != nil {
		return nil, err
	}
	text = text[1:]
	charsList := strings.Split(text, "\n\n")
	charsMap := make(map[rune]string)
	for i := 0; i < len(charsList); i++ {
		charsMap[' '+rune(i)] = charsList[i]
	}
	if len(charsMap) != 95 {
		for i := 0; i < 95; i++ {
			charsMap[' '+rune(i)] = " "
		}
		return nil, fmt.Errorf("incorrect template, please fix this")
	}
	return charsMap, nil
}
