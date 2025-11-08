package cmd

import "fmt"

const Version = "v0.1.0"

// PrintBanner prints the RiOS ASCII art banner
func PrintBanner() {
	fmt.Println()
	fmt.Println("+-------------------------------------------------+")
	fmt.Println("|   RRRRR   IIIII   OOOOO   SSSSS                 |")
	fmt.Println("|   R    R    I    O     O  S                     |")
	fmt.Println("|   R    R    I    O     O  S                     |")
	fmt.Println("|   RRRRR     I    O     O   SSSSS                |")
	fmt.Println("|   R   R     I    O     O        S               |")
	fmt.Println("|   R    R    I    O     O        S               |")
	fmt.Println("|   R     R IIIII   OOOOO   SSSSS                 |")
	fmt.Println("+-------------------------------------------------+")
	fmt.Printf("|  RiOS Compute Network - Worker %-16s |\n", Version)
	fmt.Println("+-------------------------------------------------+")
	fmt.Println()
}
