package main

import (
    "fmt"
)

func validateSingleBoard(board [][]byte, str, stc, enr, enc int) bool {
    singleBoardMap := make(map[byte]bool)
    for row := str; row < enr; row++ {
        for col := stc; col < enc; col++ {
            if board[row][col] == '.' {
                continue
            }
            if singleBoardMap[board[row][col]] {
                return false
            }
            singleBoardMap[board[row][col]] = true
        }
    }
    return true
}

func isValidSudoku(board [][]byte) bool {
    // Validate rows
    for row := 0; row < 9; row++ {
        rowMap := make(map[byte]bool)
        for col := 0; col < 9; col++ {
            if board[row][col] == '.' {
                continue
            }
            if rowMap[board[row][col]] {
                return false
            }
            rowMap[board[row][col]] = true
        }
    }

    // Validate columns
    for col := 0; col < 9; col++ {
        colMap := make(map[byte]bool)
        for row := 0; row < 9; row++ {
            if board[row][col] == '.' {
                continue
            }
            if colMap[board[row][col]] {
                return false
            }
            colMap[board[row][col]] = true
        }
    }

    // Validate 3x3 sub-grids
    for str := 0; str < 9; str += 3 {
        for stc := 0; stc < 9; stc += 3 {
            if !validateSingleBoard(board, str, stc, str+3, stc+3) {
                return false
            }
        }
    }

    return true
}


func main() {
    board1 := [][]byte{
        {'5', '3', '.', '.', '7', '.', '.', '.', '.'},
        {'6', '.', '.', '1', '9', '5', '.', '.', '.'},
        {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
        {'8', '.', '.', '.', '6', '.', '.', '.', '3'},
        {'4', '.', '.', '8', '.', '3', '.', '.', '1'},
        {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
        {'.', '6', '.', '.', '.', '.', '2', '8', '.'},
        {'.', '.', '.', '4', '1', '9', '.', '.', '5'},
        {'.', '.', '.', '.', '8', '.', '.', '7', '9'},
    }

    board2 := [][]byte{
        {'8', '3', '.', '.', '7', '.', '.', '.', '.'},
        {'6', '.', '.', '1', '9', '5', '.', '.', '.'},
        {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
        {'8', '.', '.', '.', '6', '.', '.', '.', '3'},
        {'4', '.', '.', '8', '.', '3', '.', '.', '1'},
        {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
        {'.', '6', '.', '.', '.', '.', '2', '8', '.'},
        {'.', '.', '.', '4', '1', '9', '.', '.', '5'},
        {'.', '.', '.', '.', '8', '.', '.', '7', '9'},
    }

    fmt.Println("Board 1 is valid:", isValidSudoku(board1)) // Output: true
    fmt.Println("Board 2 is valid:", isValidSudoku(board2)) // Output: false
}
