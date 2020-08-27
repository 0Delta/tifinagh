package main

import (
	"fmt"
	"os"
	"strings"
)

func concat(b rune, a rune) string {
	return string(b) + string(a)
}

func main() {
	a := []rune("ⴰⵉⵓⴻⵓ")
	b := []rune("ⴽⵛⵜⵏⵀⵎⵢⵔⵡ")
	c := []rune("ⴳⵣⴷⴱ")
	d := []rune("ⵒ")

	m1 := map[string]string{
		"あ": string(a[0]), "い": string(a[1]), "う": string(a[2]), "え": string(a[3]), "お": string(a[4]),
	}
	m2 := map[string]string{
		"か": concat(b[0], a[0]), "き": concat(b[0], a[1]), "く": concat(b[0], a[2]), "け": concat(b[0], a[3]), "こ": concat(b[0], a[4]),
		"さ": concat(b[1], a[0]), "し": concat(b[1], a[1]), "す": concat(b[1], a[2]), "せ": concat(b[1], a[3]), "そ": concat(b[1], a[4]),
		"た": concat(b[2], a[0]), "ち": concat(b[2], a[1]), "つ": concat(b[2], a[2]), "て": concat(b[2], a[3]), "と": concat(b[2], a[4]),
		"な": concat(b[3], a[0]), "に": concat(b[3], a[1]), "ぬ": concat(b[3], a[2]), "ね": concat(b[3], a[3]), "の": concat(b[3], a[4]),
		"は": concat(b[4], a[0]), "ひ": concat(b[4], a[1]), "ふ": concat(b[4], a[2]), "へ": concat(b[4], a[3]), "ほ": concat(b[5], a[4]),
		"ま": concat(b[5], a[0]), "み": concat(b[5], a[1]), "む": concat(b[5], a[2]), "め": concat(b[5], a[3]), "も": concat(b[6], a[4]),
		"や": concat(b[6], a[0]), "ゆ": concat(b[6], a[2]), "よ": concat(b[6], a[4]),
		"ら": concat(b[7], a[0]), "り": concat(b[7], a[1]), "る": concat(b[7], a[2]), "れ": concat(b[7], a[3]), "ろ": concat(b[8], a[4]),
		"わ": concat(b[8], a[0]), "を": concat(b[8], a[2]), "ん": concat(b[8], a[4]),
		"が": concat(c[0], a[0]), "ぎ": concat(c[0], a[1]), "ぐ": concat(c[0], a[2]), "げ": concat(c[0], a[3]), "ご": concat(c[0], a[4]),
		"ざ": concat(c[1], a[0]), "じ": concat(c[1], a[1]), "ず": concat(c[1], a[2]), "ぜ": concat(c[1], a[3]), "ぞ": concat(c[1], a[4]),
		"だ": concat(c[2], a[0]), "ぢ": concat(c[2], a[1]), "づ": concat(c[2], a[2]), "で": concat(c[2], a[3]), "ど": concat(c[2], a[4]),
		"ば": concat(c[3], a[0]), "び": concat(c[3], a[1]), "ぶ": concat(c[3], a[2]), "べ": concat(c[3], a[3]), "ぼ": concat(c[3], a[4]),
		"ぱ": concat(d[0], a[0]), "ぴ": concat(d[0], a[1]), "ぷ": concat(d[0], a[2]), "ぺ": concat(d[0], a[3]), "ぽ": concat(d[0], a[4]),
	}

	// fmt.Printf("%q\n", m)
	ret := ""
	reverse := false
	for i, v := range os.Args {
		if i == 0 {
			continue
		}
		if i == 1 && v == "-r" {
			reverse = true
			continue
		}
		for mk, mv := range m2 {
			if !reverse {
				v = strings.Replace(v, mk, mv, -1)
			} else {
				v = strings.Replace(v, mv, mk, -1)
			}
		}
		for mk, mv := range m1 {
			if !reverse {
				v = strings.Replace(v, mk, mv, -1)
			} else {
				v = strings.Replace(v, mv, mk, -1)
			}
		}
		ret += v
	}
	fmt.Println(ret)
}
