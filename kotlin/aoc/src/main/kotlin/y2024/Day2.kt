package y2024

class Day2 {
    fun part1(input: String): Int {
        return input.lines()
            .map { line -> line.split("\\s+".toRegex()).map { it.toInt() } }
            .count(::isSafePart1)
    }

    fun part2(input: String): Int {
        return input.lines()
            .map { line -> line.split("\\s+".toRegex()).map { it.toInt() } }
            .count(::isSafePart2)
    }

    private fun isSafePart1(items: List<Int>): Boolean {
        val diffs = items.windowed(2).map { it[1] - it[0] }
        return diffs.all { it in -1 downTo -3 } || diffs.all { it in 1..3 }
    }

    private fun isSafePart2(items: List<Int>): Boolean {
        if (isSafePart1(items)) {
            return true
        }

        // remove each one and check again
        for (i in 0..items.size) {
            val filteredList = items.filterIndexed { j, _ -> i != j }
            if (isSafePart1(filteredList)) {
                return true
            }
        }

        return false
    }
}
