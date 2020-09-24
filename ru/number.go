package ru

import (
	"math"
)

var power = []struct {
	gender int
	one    string
	four   string
	many   string
}{
	{}, // add one item to match a loop counter
	{gender: 1, one: "тысяча ", four: "тысячи ", many: "тысяч "},
	{gender: 0, one: "миллион ", four: "миллиона ", many: "миллионов "},
	{gender: 0, one: "миллиард ", four: "миллиарда ", many: "миллиардов "},
	{gender: 0, one: "триллион ", four: "триллиона ", many: "триллионов "},
}

var unit = []struct {
	one []string
	two string
	dec string
	hun string
}{
	{one: []string{"", ""}, two: "десять ", dec: "", hun: ""},
	{one: []string{"один ", "одна "}, two: "одиннадцать ", dec: "десять ", hun: "сто "},
	{one: []string{"два ", "две "}, two: "двенадцать ", dec: "двадцать ", hun: "двести "},
	{one: []string{"три ", "три "}, two: "тринадцать ", dec: "тридать ", hun: "триста "},
	{one: []string{"четыре ", "четыре "}, two: "четырнадцать ", dec: "сорок ", hun: "четыреста "},
	{one: []string{"пять ", "пять "}, two: "пятнадцать ", dec: "пятьдесят ", hun: "пятьсот "},
	{one: []string{"шесть ", "шесть "}, two: "шестнадцать ", dec: "шестьдесят ", hun: "шестьсот "},
	{one: []string{"семь ", "семь "}, two: "семнадцать ", dec: "семьдесят ", hun: "семьсот "},
	{one: []string{"восемь ", "восемь "}, two: "восемнадцать ", dec: "восемьдесят ", hun: "восемьсот "},
	{one: []string{"девять ", "девять "}, two: "девятнадцать ", dec: "девяносто ", hun: "девятьсот "},
}

func Number2Text(value int) string {
	if value == 0 {
		return "ноль"
	}

	var result string
	if value < 0 {
		result = "минус "
		value = -value
	}

	maxPower := len(power) - 1
	divisor := int(math.Pow(1000, float64(maxPower)))

	for i := maxPower - 1; i >= 0; i-- {
		divisor /= 1000
		quotient := value / divisor
		value %= divisor

		if quotient == 0 {
			continue
		}

		var str string
		for quotient >= 20 {
			if quotient >= 100 {
				str += unit[quotient/100].hun
				quotient %= 100
			} else if quotient >= 20 {
				str += unit[quotient/10].dec
				quotient %= 10
			}
		}
		if quotient >= 10 {
			str += unit[quotient-10].two
		} else {
			str += unit[quotient].one[power[i].gender]
		}

		switch quotient {
		case 1:
			str += power[i].one
		case 2, 3, 4:
			str += power[i].four
		default:
			str += power[i].many
		}
		result += str
	}
	return result[:len(result)-1] // remove the last space
}
