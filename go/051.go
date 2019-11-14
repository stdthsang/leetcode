func solveNQueens(n int) [][]string {
    cur := make([][]byte, n)
    for i := 0; i < n; i++ {
        cur[i] = make([]byte, n)
        for j := 0; j < n; j++ {
            cur[i][j] = '.'
        }
    }
    var res [][]string
    helper(0, n, cur, &res)
    return res
}

func helper(row, n int, cur [][]byte, out *[][]string) {
    if row == n {
        temp := make([]string, n)
        for i := 0; i < n; i++ {
            bytes := make([]byte, n)
            copy(bytes, cur[i])
            temp[i] = string(bytes)
        }
        *out = append(*out, temp)
        return
    }
    for i := 0; i < n; i++ {
        if isValid(row, i, n, cur) {
            cur[row][i] = 'Q'
            helper(row+1, n, cur, out)
            cur[row][i] = '.'
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

