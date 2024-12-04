package y2024

import utils.mapIgnoreErrors

class Day4 {

    fun part1(input: String): Int {
        val matrix = input.lines()
        val directions = listOf(
            listOf(0 to 0, -1 to 1, -2 to 2, -3 to 3), // diag up
            listOf(0 to 0, 0 to 1, 0 to 2, 0 to 3),    // right
            listOf(0 to 0, 1 to 1, 2 to 2, 3 to 3),    // diag down
            listOf(0 to 0, 1 to 0, 2 to 0, 3 to 0)     // down
        )

        var result = 0

        for (rowIdx in matrix.indices) {
            for (colIdx in 0..<matrix[0].length) {
                result += directions
                    .map { dir ->
                        dir
                            .mapIgnoreErrors { matrix[rowIdx + it.first][colIdx + it.second] }
                            .joinToString("")
                    }
                    .count { it == "XMAS" || it == "SAMX" }
            }
        }

        return result
    }

    fun part2(input: String): Int {
        return input
            .lines()
            .map { it.toList() }
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
            }.count {
                val x1 = listOf(it[0][0], it[1][1], it[2][2]).joinToString("")
                val x2 = listOf(it[2][0], it[1][1], it[0][2]).joinToString("")
                (x1 == "MAS" || x1 == "SAM") && (x2 == "MAS" || x2 == "SAM")
            }
    }
}
