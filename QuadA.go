package main

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	if x < 0 || y < 0 {
		return
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			// Will determine the carachter to print based on the position
			var char rune
			if i == 0 || i == y-1 {
				char = '-'
			} else if j == 0 || j == x-1 {
				char = '|'
			} else {
				char = ' '
			}
			// Prints the character
			z01.PrintRune(char)

		}
		// Prints a new line after the end of each row
		z01.PrintRune('\n')
	}

}
func main() {
	QuadA(5, 3)
}
