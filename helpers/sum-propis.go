// Package num2word holds Number to words converter
// package num2word
package main

import (
	"fmt"
	"strings"
	"unicode"
)

var repl = [][]string{
	
	// t - тысячи; m - милионы; M - миллиарды;
	{",,,,,,", "eM"},
	{",,,,,", "em"},
	{",,,,", "et"},
	// e - единицы; d - десятки; c - сотни;
	{",,,", "e"},
	{",,", "d"},
	{",", "c"},
	{"0c0d0et", ""},
	{"0c0d0em", ""},
	{"0c0d0eM", ""},
	
	// --
	{"0c",   ""},
	{"1c",   "сто "},
    {"2c",   "двісті "},
    {"3c",   "триста "},
    {"4c",   "чотириста "},
    {"5c",   "п'ятсот "},
    {"6c",   "шістсот "},
    {"7c",   "сімсот "},
    {"8c",   "вісімсот "},
    {"9c",   "дев'ятсот "},
    {"1d0e", "десять "},
    {"1d1e", "одинадцять "},
    {"1d2e", "дванадцять "},
    {"1d3e", "тринадцять "},
    {"1d4e", "чотирнадцять "},
    {"1d5e", "п'ятнадцять "},
    {"1d6e", "шістнадцять "},
    {"1d7e", "сімнадцять "},
    {"1d8e", "вісімнадцять "},
    {"1d9e", "дев'ятнадцять "},
    	// --
	{"0d", ""},
	{"2d", "двадцять "},
    {"3d", "тридцять "},
    {"4d", "сорок "},
    {"5d", "п`ятдесят "},
    {"6d", "шістдесят "},
    {"7d", "сімдесят "},
    {"8d", "вісімдесят "},
    {"9d", "дев'яносто "},
	// --
	{ "0e", ""},
	{ "5e", "п'ять "},
    { "6e", "шість "},
    { "7e", "сім "},
    { "8e", "вісім "},
    { "9e", "дев`ять "},
    // -
    { "1e.", "одна гривна "},
    { "2e.", "дві гривни "},
    { "3e.", "три гривни "},
    { "4e.", "Чотири гривни "},
    { "1et", "одна тисяча "},
    { "2et", "дві тисячі "},
    { "3et", "три тисячі "},
    { "4et", "чотири тисячі "},
    { "1em", "один мільйон "},
    { "2em", "два мільйони "},
    { "3em", "три мільйони "},
    { "4em", "чотири мільйони "},
    { "1eM", "один мільярд "},
    { "2eM", "два мільярди "},
    { "3eM", "три мільярди "},
    { "4eM", "чотири мільярди "},
    
    // блок для написання копійок без скорочення "коп"
    { "11k", "11 копійок"},
    { "12k", "12 копійок"},
    { "13k", "13 копійок"},
    { "14k", "14 копійок"},
    { "1k",  "1 копійка"},
    { "2k",  "2 копійки"},
    { "3k",  "3 копійки"},
    { "4k",  "4 копійки"},
    { "K",   "копійок"},
    // -
    { ".", "гривнi "},
    { "t", "тисяч "},
    { "m", "мільйонів "},
    { "M", "мільярдів "},
}


var repl3 = [][]string{
	// t - тысячи; m - милионы; M - миллиарды;
	{",,,,,,", "eM"},
	{",,,,,", "em"},
	{",,,,", "et"},
	// e - единицы; d - десятки; c - сотни;
	{",,,", "e"},
	{",,", "d"},
	{",", "c"},
	{"0c0d0et", ""},
	{"0c0d0em", ""},
	{"0c0d0eM", ""},
	// --
	{"0c", ""},
	{"1c", "сто "},
	{"2c", "двести "},
	{"3c", "триста "},
	{"4c", "четыреста "},
	{"5c", "пятьсот "},
	{"6c", "шестьсот "},
	{"7c", "семьсот "},
	{"8c", "восемьсот "},
	{"9c", "девятьсот "},
	{"1d0e", "десять "},
	{"1d1e", "одиннадцать "},
	{"1d2e", "двенадцать "},
	{"1d3e", "тринадцать "},
	{"1d4e", "четырнадцать "},
	{"1d5e", "пятнадцать "},
	{"1d6e", "шестнадцать "},
	{"1d7e", "семнадцать "},
	{"1d8e", "восемнадцать "},
	{"1d9e", "девятнадцать "},
	// --
	{"0d", ""},
	{"2d", "двадцать "},
	{"3d", "тридцать "},
	{"4d", "сорок "},
	{"5d", "пятьдесят "},
	{"6d", "шестьдесят "},
	{"7d", "семьдесят "},
	{"8d", "восемьдесят "},
	{"9d", "девяносто "},
	// --
	{"0e", ""},
	{"5e", "пять "},
	{"6e", "шесть "},
	{"7e", "семь "},
	{"8e", "восемь "},
	{"9e", "девять "},
	// --
	{"1e.", "одна гривна "},
	{"2e.", "два гривны "},
	{"3e.", "три гривны "},
	{"4e.", "четыре гривны "},
	{"1et", "одна тысяча "},
	{"2et", "две тысячи "},
	{"3et", "три тысячи "},
	{"4et", "четыре тысячи "},
	{"1em", "один миллион "},
	{"2em", "два миллиона "},
	{"3em", "три миллиона "},
	{"4em", "четыре миллиона "},
	{"1eM", "один миллиард "},
	{"2eM", "два миллиарда "},
	{"3eM", "три миллиарда "},
	{"4eM", "четыре миллиарда "},
	//  блок для написания копеек без сокращения "коп"
	{"11k", "11 копеек"},
	{"12k", "12 копеек"},
	{"13k", "13 копеек"},
	{"14k", "14 копеек"},
	{"1k", "1 копейка"},
	{"2k", "2 копейки"},
	{"3k", "3 копейки"},
	{"4k", "4 копейки"},
	{"k", " копеек"},
	// --
	{".", "гривен "},
	{"t", "тысяч "},
	{"m", "миллионов "},
	{"M", "миллиардов "},
}

var mask = []string{",,,", ",,", ",", ",,,,", ",,", ",", ",,,,,", ",,", ",", ",,,,,,", ",,", ","}

// RuMoney - деньги прописью на русском
func RuMoney(number float64, upperFirstChar bool) string {

	s := fmt.Sprintf("%.2f", number)
	l := len(s)

	dest := s[l-3:l] + "k"
	s     = s[:l-3]
	l     = len(s)
	
	for i := l; i > 0; i-- {
		c   := string(s[i-1])
		dest = c + mask[l-i] + dest
	}

	for _, r := range repl {
		dest = strings.Replace(dest, r[0], r[1], -1)
	}

	if upperFirstChar {
		a    := []rune(dest)
		a[0]  = unicode.ToUpper(a[0])
		dest  = string(a)
	}

	return dest
}



func main(){
	f:=RuMoney(235456.23, false)
	fmt.Println(235456.23, f)
}
