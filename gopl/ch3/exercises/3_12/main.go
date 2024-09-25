package main

func main() {

}

func IsIsomerization(s1, s2 string) bool {
	if len(s1) != len(s2) || s1 == s2 {
		return false
	}
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)
	for _, r := range s1 {
		m1[r]++
	}
	for _, r := range s2 {
		m2[r]++
	}

}
