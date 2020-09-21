package ru

import (
	"math"
)

type power struct {
	sex  int
	one  string
	four string
	many string
}

var powers = []power{
	{}, // add one item to match a loop counter
	{sex: 1, one: "тысяча ", four: "тысячи ", many: "тысяч "},
	{sex: 0, one: "миллион ", four: "миллиона ", many: "миллионов "},
	{sex: 0, one: "миллиард ", four: "миллиарда ", many: "миллиардов "},
	{sex: 0, one: "триллион ", four: "триллиона ", many: "триллионов "},
}

type unit struct {
	one []string
	two string
	dec string
	hun string
}

var units = []unit{
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
		return "ноль "
	}

	var result string
	if value < 0 {
		result = "минус "
		value = -value
	}

	maxPower := len(powers) - 1
	divisor := int(math.Pow(1000, float64(maxPower)))

	for i := maxPower - 1; i >= 0; i-- {
		divisor /= 1000
		carry := value / divisor
		value %= divisor
		var str string

		if carry == 0 {
			continue
		}

		for carry >= 20 {
			if carry >= 100 {
				str += units[carry/100].hun
				carry %= 100
			} else if carry >= 20 {
				str += units[carry/10].dec
				carry %= 10
			}
		}
		if carry >= 10 {
			str += units[carry-10].two
		} else {
			str += units[carry].one[powers[i].sex]
		}

		switch carry {
		case 1:
			str += powers[i].one
		case 2, 3, 4:
			str += powers[i].four
		default:
			str += powers[i].many
		}
		result += str
	}
	return result[:len(result)-1] // remove the last space
}
