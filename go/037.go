func solveSudoku(board [][]byte) {
    help(board, 0, 0)
}

func help(board [][]byte, row, col int) bool {
    if col == len(board[0]) {
        return help(board, row+1, 0)
    }
    if row == len(board) {
        return true
    }
    if board[row][col] != '.' {
        return help(board, row, col+1)
    }
    for i := 1; i <= 9; i++ {
        ch := byte(i) + '0'
        if valid(board, row, col, ch) {
            board[row][col] = ch
            if help(board, row, col+1) {
                return true
            }
            board[row][col] = '.'
        }
    }
    return false
}

func valid(board [][]byte, row, col int, ch byte) bool {
    m, n := len(board), len(board[0])
    for i := 0; i < m; i++ {
        if board[i][col] == ch {
            return false
        }
    }
    for i := 0; i < n; i++ {
        if board[row][i] == ch {
            return false
        }
    }
    for i := row / 3 * 3; i < row/3*3+3; i++ {
        for j := col / 3 * 3; j < col/3*3+3; j++ {
            if board[i][j] == ch {
                return false
            }
        }
    }
    return true
}
