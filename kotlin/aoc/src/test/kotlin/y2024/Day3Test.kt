package y2024

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import utils.readFile

class Day3Test {
    private val day = Day3()
    private val file = "day3"

    @Test
    fun part1() {
        assertEquals(161, day.part1(readFile("/y2024/${file}_sample.txt")))
        assertEquals(161289189, day.part1(readFile("/y2024/${file}.txt")))
    }

    @Test
    fun part2() {
        assertEquals(48, day.part2(readFile("/y2024/${file}_sample2.txt")))
        assertEquals(83595109, day.part2(readFile("/y2024/${file}.txt")))
    }
}
