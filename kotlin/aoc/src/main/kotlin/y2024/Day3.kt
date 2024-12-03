package y2024

import utils.findAll

sealed interface Op
data class Mul(val fst: Int, val snd: Int) : Op
data object Do : Op
data object Dont : Op

class Day3 {
    fun part1(input: String): Int = input.parsePart1().sumOf { it.fst * it.snd }

    fun part2(input: String): Int =
        input
            .parsePart2()
            .fold(0 to true) { acc, op ->
                when {
                    op is Do -> acc.first to true
                    op is Dont -> acc.first to false
                    acc.second && op is Mul -> (acc.first + op.fst * op.snd) to true
                    else -> acc
                }
            }
            .first

    private fun String.parsePart1(): Sequence<Mul> =
        findAll("mul\\((\\d+),(\\d+)\\)")
            .map { it.groupValues.drop(1) }
            .map { (fst, snd) -> Mul(fst.toInt(), snd.toInt()) }

    private fun String.parsePart2(): Sequence<Op> =
        findAll("mul\\((\\d+),(\\d+)\\)|do\\(\\)|don't\\(\\)")
            .map {
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
