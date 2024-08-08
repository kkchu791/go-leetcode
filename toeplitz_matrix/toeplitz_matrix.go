func isToeplitzMatrix(matrix [][]int) bool {
    for r := 0; r < len(matrix) - 1; r++ {
        for c := 0; c < len(matrix[0]) - 1; c++ {
            curr := matrix[r][c]
            directDiagonalOfCurrent := matrix[r+1][c+1]

            if curr != directDiagonalOfCurrent {
                return false
            }
        }
    }

    return true
}


