package main

import (
	"fmt"
)

var table = [][]string{}
var rows, columns int

// using in func countRoundOfPlayers()
var round int

func createTable() {
	// using index 1, 2, 3 to play
	row1 := []string{"|", " ", " ", " ", "|"}
	row2 := []string{"|", " ", " ", " ", "|"}
	row3 := []string{"|", " ", " ", " ", "|"}

	table = append(table, row1)
	table = append(table, row2)
	table = append(table, row3)
}

func main() {
	go createTable()

	for {
		// display XO game.
		go display()
		fmt.Scan(&rows)
		if rows == 0 {
			break
		}
		fmt.Scan(&columns)
		go input()
	}
}

func display() {
	fmt.Println("  -----------")

	for i, s := range table {
		fmt.Println(i, s)
	}

	// Basic
	// fmt.Println(table[0])
	// fmt.Println(table[1])
	// fmt.Println(table[2])

	fmt.Println("  -----------")

	gameOver := checkGameOver()
	if gameOver != true {
		fmt.Print("Enter the coordinates: ")
	}
}

func input() {

	switch {
	case rows == 1 && columns == 1:
		if table[0][1] == "X" || table[0][1] == "O" {
			fmt.Println("Please new select to play.")
			break
		}
		table[0][1] = countRoundOfPlayers()
		break
	case rows == 1 && columns == 2:
		if table[0][2] == "X" || table[0][2] == "O" {
			fmt.Println("Please new select to play.")
			break
		}
		table[0][2] = countRoundOfPlayers()
		break
	case rows == 1 && columns == 3:
		if table[0][3] == "X" || table[0][3] == "O" {
			fmt.Println("Please new select to play.")
			break
		}
		table[0][3] = countRoundOfPlayers()
		break
	case rows == 2 && columns == 1:
		if table[1][1] == "X" || table[1][1] == "O" {
			fmt.Println("Please new select to play.")
			break
		}
		table[1][1] = countRoundOfPlayers()
		break
	case rows == 2 && columns == 2:
		if table[1][2] == "X" || table[1][2] == "O" {
			fmt.Println("Please new select to play.")
			break
		}
		table[1][2] = countRoundOfPlayers()
		break
	case rows == 2 && columns == 3:
		if table[1][3] == "X" || table[1][3] == "O" {
			fmt.Println("Please new select to play.")
			break
		}
		table[1][3] = countRoundOfPlayers()
		break
	case rows == 3 && columns == 1:
		if table[2][1] == "X" || table[2][1] == "O" {
			fmt.Println("Please new select to play.")
			break
		}
		table[2][1] = countRoundOfPlayers()
		break
	case rows == 3 && columns == 2:
		if table[2][2] == "X" || table[2][2] == "O" {
			fmt.Println("Please new select to play.")
			break
		}
		table[2][2] = countRoundOfPlayers()
		break
	case rows == 3 && columns == 3:
		if table[2][3] == "X" || table[2][3] == "O" {
			fmt.Println("Please new select to play.")
			break
		}
		table[2][3] = countRoundOfPlayers()
		break
	default:
		fmt.Println("Please new select to play.")
		break
	}
}

func countRoundOfPlayers() string {
	round++
	var players string
	if round%2 != 0 {
		players = "X"
	} else if round%2 == 0 {
		players = "O"
	}
	return players
}

func checkGameOver() bool {
	switch {
	// row 1
	case table[0][1] == "X" && table[0][2] == "X" && table[0][3] == "X":
		fmt.Println("X Wins")
		fmt.Println()
		fmt.Println("Process finished with exit code 0")
		return true
	case table[0][1] == "O" && table[0][2] == "O" && table[0][3] == "O":
		fmt.Println("O Wins")
		fmt.Println()
		fmt.Println("Process finished with exit code 0")
		return true

		// row 2
	case table[1][1] == "X" && table[1][2] == "X" && table[1][3] == "X":
		fmt.Println("X Wins")
		fmt.Println()
		fmt.Println("Process finished with exit code 0")
		return true
	case table[1][1] == "O" && table[1][2] == "O" && table[1][3] == "O":
		fmt.Println("O Wins")
		fmt.Println()
		fmt.Println("Process finished with exit code 0")
		return true

		// row 3
	case table[2][1] == "X" && table[2][2] == "X" && table[2][3] == "X":
		fmt.Println("X Wins")
		fmt.Println()
		fmt.Println("Process finished with exit code 0")
		return true
	case table[2][1] == "O" && table[2][2] == "O" && table[2][3] == "O":
		fmt.Println("O Wins")
		fmt.Println()
		fmt.Println("Process finished with exit code 0")
		return true

		// col 1
	case table[0][1] == "X" && table[1][1] == "X" && table[2][1] == "X":
		fmt.Println("X Wins")
		fmt.Println()
		fmt.Println("Process finished with exit code 0")
		return true
	case table[0][1] == "O" && table[1][1] == "O" && table[2][1] == "O":
		fmt.Println("O Wins")
		fmt.Println()
		fmt.Println("Process finished with exit code 0")
		return true

		// col 2
	case table[0][2] == "X" && table[1][2] == "X" && table[2][2] == "X":
		fmt.Println("X Wins")
		fmt.Println()
		fmt.Println("Process finished with exit code 0")
		return true
	case table[0][2] == "O" && table[1][2] == "O" && table[2][2] == "O":
		fmt.Println("O Wins")
		fmt.Println()
		fmt.Println("Process finished with exit code 0")
		return true

		// col 3
	case table[0][3] == "X" && table[1][3] == "X" && table[2][3] == "X":
		fmt.Println("X Wins")
		fmt.Println()
		fmt.Println("Process finished with exit code 0")
		return true
	case table[0][3] == "O" && table[1][3] == "O" && table[2][3] == "O":
		fmt.Println("O Wins")
		fmt.Println()
		fmt.Println("Process finished with exit code 0")
		return true

		// diagonal "\"
	case table[0][1] == "X" && table[1][2] == "X" && table[2][3] == "X":
		fmt.Println("X Wins")
		fmt.Println()
		fmt.Println("Process finished with exit code 0")
		return true
	case table[0][1] == "O" && table[1][2] == "O" && table[2][3] == "O":
		fmt.Println("O Wins")
		fmt.Println()
		fmt.Println("Process finished with exit code 0")
		return true

		// diagonal "/"
	case table[0][3] == "X" && table[1][2] == "X" && table[2][1] == "X":
		fmt.Println("X Wins")
		fmt.Println()
		fmt.Println("Process finished with exit code 0")
		return true
	case table[0][3] == "O" && table[1][2] == "O" && table[2][1] == "O":
		fmt.Println("O Wins")
		fmt.Println()
		fmt.Println("Process finished with exit code 0")
		return true
	}
	return false
}
