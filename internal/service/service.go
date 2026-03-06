package service

import (
	"strings"

	"6sprint/pkg/morse" // Импортируйте пакет morse
)

// AutoConvert определяет, в каком виде передан текст (код Морзе или обычный текст),
// и конвертирует его в противоположный формат.
func AutoConvert(input string) (string, error) {
	// Убираем все пробелы в начале и в конце строки
	input = strings.TrimSpace(input)

	// Если строка состоит только из точек, тире и пробелов, вероятно, это код Морзе
	if isMorse(input) {
		// Преобразуем код Морзе в текст
		return morse.ToText(input), nil
	}

	// Если строка обычный текст, преобразуем в код Морзе
	return morse.ToMorse(input), nil
}

// isMorse проверяет, является ли строка кодом Морзе
func isMorse(input string) bool {
	for _, c := range input {
		if !(c == '.' || c == '-' || c == ' ' || c == '\n') {
			return false
		}
	}
	return true
}
