func totalNQueens(n int) int {
    cur := make([][]byte, n)
    for i := 0; i < n; i++ {
        cur[i] = make([]byte, n)
    }
    var res [][]string
    helper(0, n, cur, &res)
    return len(res)
}

func helper(row, n int, cur [][]byte, res *[][]string) {
    if row == n {
        temp := make([]string, n)
        for i := 0; i < n; i++ {
            str := make([]byte, n)
            copy(str, cur[i])
            temp[i] = string(str)
        }
        *res = append(*res, temp)
        return
    }
    for col := 0; col < n; col++ {
        if isValid(row, col, n, cur) {
            cur[row][col] = 'Q'
            helper(row+1, n, cur, res)
            cur[row][col] = '.'
        }
    }
}

func isValid(row, col, n int, cur [][]byte) bool {
    for i := 0; i < row; i++ {
        if cur[i][col] == 'Q' {
            return false
        }
    }
    for i := 0; i < col; i++ {
        if cur[row][i] == 'Q' {
            return false
        }
    }
    for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if cur[i][j] == 'Q' {
            return false
        }
    }
    for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
        if cur[i][j] == 'Q' {
            return false
        }
    }
    return true
}
