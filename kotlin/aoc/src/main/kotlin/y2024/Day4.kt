package y2024

class Day4 {

    // TODO: clean up the code

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

    fun part2(input: String): Int {
        val matrix = input.lines().map { it.toList() }

        val threeByThrees = matrix
            .windowed(3)
            .flatMap { rows ->
                (0..rows[0].size - 3)
                    .map { col ->
                        listOf(
                            rows[0].subList(col, col + 3),
                            rows[1].subList(col, col + 3),
                            rows[2].subList(col, col + 3),
                        )
                    }
            }

        return threeByThrees.count {
            val x1 = listOf(it[0][0], it[1][1], it[2][2]).joinToString("")
            val x2 = listOf(it[2][0], it[1][1], it[0][2]).joinToString("")
            (x1 == "MAS" || x1 == "SAM") && (x2 == "MAS" || x2 == "SAM")
        }
    }
}
