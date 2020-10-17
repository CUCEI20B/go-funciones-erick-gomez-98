package main

func Burbuja(s []int64) {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				t := s[j]
				s[j] = s[j+1]
				s[j+1] = t
			}
		}
	}
}

func main() {
}