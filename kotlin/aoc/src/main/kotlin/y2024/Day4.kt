package y2024

class Day4 {

    fun part1(input: String): Int {
        val matrix = input.lines()
        val combinations = mutableListOf<String>()

        val maxRowIdx = matrix.size - 1
        val maxColIdx = matrix[0].length - 1

        for (rowIdx in 0..maxRowIdx) {
            for (colIdx in 0..maxColIdx) {
                // diag up
                if (rowIdx >= 3 && colIdx <= maxColIdx - 3) {
                    combinations += (0..3).map { matrix[rowIdx - it][colIdx + it] }.joinToString("")
                }

                // right
                if (colIdx <= maxColIdx - 3) {
                    combinations += (0..3).map { matrix[rowIdx][colIdx + it] }.joinToString("")
                }

                // diag down
                if (rowIdx <= maxRowIdx - 3 && colIdx <= maxColIdx - 3) {
                    combinations += (0..3).map { matrix[rowIdx + it][colIdx + it] }.joinToString("")

                }

                // down
                if (rowIdx <= maxRowIdx - 3) {
                    combinations += (0..3).map { matrix[rowIdx + it][colIdx] }.joinToString("")
                }
            }
        }

        return combinations.count { it == "XMAS" || it == "SAMX" }
    }

    fun part2(input: String): Int = 0
}
