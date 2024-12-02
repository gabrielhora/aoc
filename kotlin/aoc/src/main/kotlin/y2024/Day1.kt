package y2024

import utils.listOfInts
import kotlin.math.absoluteValue

class Day1 {
    fun part1(input: String): Int {
        val (left, right) = splitLists(input)
        return left
            .sorted()
            .zip(right.sorted())
            .sumOf { (l, r) -> (l - r).absoluteValue }
    }

    fun part2(input: String): Int {
        val (left, right) = splitLists(input)
        return left.sumOf { num -> num * right.count { it == num } }
    }

    private fun splitLists(input: String) = input
        .lines()
        .listOfInts()
        .map { (first, second) -> first to second }
        .unzip()
}
