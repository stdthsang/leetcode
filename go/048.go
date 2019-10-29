func rotate(matrix [][]int) {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return
    }
    row, col := len(matrix), len(matrix[0])
    for i := 0; i < row; i++ {
        for j := i + 1; j < col; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
    for i := 0; i < col/2; i++ {
        for j := 0; j < row; j++ {
            matrix[j][i], matrix[j][row-i-1] = matrix[j][row-i-1], matrix[j][i]
        }
    }
}
