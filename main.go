package main

import "fmt"

var isExit = false

func main() {

	for !isExit {
		magicSquare()
	}
}
func magicSquare() {
	var i, j, n, ldsum, rdsum, magic int
	var ar [30][30]int
	var sor [10]int
	var soc [10]int
reset:
	fmt.Print("\n-> Enter the size of the Magic Square (2 for 2x2 , 3 for 3x3): ")
	fmt.Scanln(&n)
	if n <= 1 {
		fmt.Printf("\n! The Magic Square should be have minimum size of 2x2\n\n ")
		goto reset
	}
	fmt.Println("-> Enter the elements in the matrix:")
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			fmt.Printf("%d - %d: ", i+1, j+1)
			fmt.Scanln(&ar[i][j])
			soc[i], sor[i] = 0, 0
		}
	}
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			sor[i] += ar[i][j]
			soc[i] += ar[j][i]
			if i == j {
				rdsum += ar[i][j]
			}
			if i+j == n-1 {
				ldsum += ar[i][j]
			}
		}
		if rdsum == sor[i] && ldsum == soc[i] {
			magic++
		}
	}
	if magic > 0 {
		fmt.Printf("\n  yes!!\n")
		fmt.Println("       :---------------------------------:")
		fmt.Println("       |*                               *|")
		fmt.Println("       | !!!You have the MAGIC SQUARE!!! |")
		fmt.Println("       |*                               *|")
		fmt.Println("       :---------------------------------:")
		fmt.Println()
		isExit = true
	} else {
		fmt.Println("Sorry.. it is a not a Magic Square, but thats cool..")
		fmt.Println("                   Please try again!")
		fmt.Printf("\n\n\n")
	}
}
