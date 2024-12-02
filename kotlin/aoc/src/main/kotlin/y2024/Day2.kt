package y2024

class Day2 {
    fun part1(input: String): Int {
        return input.lines()
            .map { line -> line.split("\\s+".toRegex()).map { it.toInt() } }
            .map { it.windowed(2) }
            .map { item -> item.map { it[1] - it[0] } }
            .count { diffs -> diffs.all { it in -1 downTo -3 } || diffs.all { it in 1..3 } }
    }

    fun part2(input: String): Int {
        return 0
    }
}
