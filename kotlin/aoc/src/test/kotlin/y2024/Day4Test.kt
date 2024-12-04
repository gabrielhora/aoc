package y2024

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import utils.readFile

class Day4Test {
    private val day = Day4()
    private val file = "day4"

    @Test
    fun part1() {
        assertEquals(18, day.part1(readFile("/y2024/${file}_sample.txt")))
        assertEquals(2718, day.part1(readFile("/y2024/${file}.txt")))
    }

    @Test
    fun part2() {
        assertEquals(9, day.part2(readFile("/y2024/${file}_sample.txt")))
        assertEquals(2046, day.part2(readFile("/y2024/${file}.txt")))
    }
}
