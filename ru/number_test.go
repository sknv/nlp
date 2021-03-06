package ru

import (
	"testing"
)

func TestNumber2Text(t *testing.T) {
	numbers := map[int]string{
		4:       "четыре",
		13:      "тринадцать",
		25:      "двадцать пять",
		148:     "сто сорок восемь",
		3407:    "три тысячи четыреста семь",
		6000:    "шесть тысяч",
		9112:    "девять тысяч сто двенадцать",
		17010:   "семнадцать тысяч десять",
		63414:   "шестьдесят три тысячи четыреста четырнадцать",
		877002:  "восемьсот семьдесят семь тысяч два",
		3146018: "три миллиона сто сорок шесть тысяч восемнадцать",
		9643120: "девять миллионов шестьсот сорок три тысячи сто двадцать",
	}
	for num := range numbers {
		text := Number2Text(num)
		if text != numbers[num] {
			t.Errorf("Expected: %s, got instead: %s", numbers[num], text)
		}
	}
}
