package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func isValidCell(board *[9][9]int, value int, x int, y int) bool {
	// Check in rows
	for j := range board[x] {
		if j == y {
			continue
		}
		if board[x][j] == value {
			return false
		}
	}

	// Check in columns
	for i := range board {
		if i == x {
			continue
		}
		if board[i][y] == value {
			return false
		}
	}

	// Check in 3x3
	k := (x / 3) * 3
	l := (y / 3) * 3

	for i := k; i < k+3; i++ {
		for j := l; j < l+3; j++ {
			if i == x && j == y {
				continue
			}
			if board[i][j] == value {
				return false
			}
		}
	}

	return true
}

func isValidBoard(board *[9][9]int) bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 0 {
				continue
			}

			if !isValidCell(board, board[i][j], i, j) {
				return false
			}

		}
	}

	return true
}

func isValidChar(c int32) bool {
	return c >= 49 && c <= 57 || c == 45
}

func solveSudoku(board [9][9]int, i int, j int) ([9][9]int, error) {

	if j >= 9 {
		j = 0
		i++
	}

	if i >= 9 {
		return board, nil
	}

	if board[i][j] != 0 {
		return solveSudoku(board, i, j+1)
	}

	for value := 1; value <= 9; value++ {
		if !isValidCell(&board, value, i, j) {
			continue
		}

		board[i][j] = value
		newBoard, e := solveSudoku(board, i, j+1)
		if e != nil {
			continue
		}

		return newBoard, nil
	}

	return board, fmt.Errorf("invalid previous number")
}

func printBoard(board *[9][9]int) {
	for i := range board {
		if i%3 == 0 {
			fmt.Println("  |-----------|-----------|-----------|  ")
		}

		for j := range board[i] {
			if j%3 == 0 {
				fmt.Print("  |  ")
			} else {
				fmt.Print("  ")
			}

			fmt.Print(board[i][j])
		}
		fmt.Println("  |  ")
	}
	fmt.Println("  |-----------|-----------|-----------|  ")
}

func readBoardFromFile(board *[9][9]int, fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	var tempBoard [9][9]int
	j := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if j >= 9 {
			break
		}

		for i, c := range line {
			if !isValidChar(c) {
				return fmt.Errorf("file contains invalid charecters")
			}

			if c == 45 {
				tempBoard[j][i] = 0
			} else {
				tempBoard[j][i] = int(c) - 48
			}
		}

		j++
	}

	for i := range board {
		for j := range board[i] {
			board[i][j] = tempBoard[i][j]
		}
	}

	return nil
}

func main() {
	var board = [9][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	if len(os.Args) > 1 {
		fileName := os.Args[1]

		if err := readBoardFromFile(&board, fileName); err != nil {
			fmt.Println(err)
		}
	}

	if !isValidBoard(&board) {
		fmt.Println("Given sudoku is invalid.")
		return
	}

	startTime := time.Now()
	solvedBoard, err := solveSudoku(board, 0, 0)
	duration := time.Now().Sub(startTime)
	fmt.Println(duration)
	if err != nil {
		fmt.Println("The sudoku board is unsolvable.")
		return
	}

	printBoard(&solvedBoard)
}
