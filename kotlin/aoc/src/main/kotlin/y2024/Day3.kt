package y2024

import utils.findAll

class Day3 {
    fun part1(input: String): Int {
        return input
            .findAll("mul\\((\\d+),(\\d+)\\)")
            .map { it.groupValues.drop(1) }
            .sumOf { (fst, snd) -> fst.toInt() * snd.toInt() }
    }

    fun part2(input: String): Int {
        return input
            .findAll("mul\\((\\d+),(\\d+)\\)|do\\(\\)|don't\\(\\)")
            .fold(0 to 1) { acc, match ->
                val gv = match.groupValues
                val op = gv.first()
                when {
                    op == "do()" -> acc.first to 1
                    op == "don't()" -> acc.first to 0
                    op.startsWith("mul") -> acc.first + acc.second * gv[1].toInt() * gv[2].toInt() to acc.second
                    else -> throw Exception("invalid")
                }
            }
            .first
    }
}
