package y2024

sealed interface Op
data class Mul(val first: Int, val second: Int) : Op
data object Do : Op
data object Dont : Op

class Day3 {
    fun part1(input: String): Int {
        return parsePart1(input.replace("\n", ""))
            .sumOf { it.first * it.second }
    }

    fun part2(input: String): Int {
        return parsePart2(input.replace("\n", ""))
            .fold(0 to true) { acc, op ->
                when {
                    op is Do -> acc.first to true
                    op is Dont -> acc.first to false
                    acc.second && op is Mul -> (acc.first + op.first * op.second) to true
                    else -> acc
                }
            }
            .first
    }

    private fun parsePart1(input: String): Sequence<Pair<Int, Int>> {
        val matches = Regex("mul\\((\\d+),(\\d+)\\)").findAll(input)
        return matches
            .map { it.groupValues.drop(1) }
            .map { (fst, snd) -> fst.toInt() to snd.toInt() }
    }

    private fun parsePart2(input: String): Sequence<Op> {
        val matches = Regex("mul\\((\\d+),(\\d+)\\)|do\\(\\)|don't\\(\\)").findAll(input)
        return matches.map {
            val values = it.groupValues
            val op = values.first()
            when {
                op.startsWith("mul") -> Mul(values[1].toInt(), values[2].toInt())
                op == "do()" -> Do
                op == "don't()" -> Dont
                else -> throw Exception("unknown value")
            }
        }
    }

}
