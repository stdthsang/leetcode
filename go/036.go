func isValidSudoku(board [][]byte) bool {
    if len(board) == 0 || len(board[0]) == 0 {
        return false
    }
    m, n := len(board), len(board[0])
    row, col, cell := make([][]bool, m), make([][]bool, m), make([][]bool, m)
    for i := 0; i < m; i++ {
        row[i], col[i], cell[i] = make([]bool, n), make([]bool, n), make([]bool, n)

    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == '.' {
                continue
            }
            c := int(board[i][j]-'0') - 1
            if row[i][c] || col[c][j] || cell[3*(i/3)+j/3][c] {
                return false
            }
            row[i][c], col[c][j], cell[3*(i/3)+j/3][c] = true, true, true
        }
    }
    return true
}
