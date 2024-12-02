package y2024

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import utils.readFile

class Day2Test {
    private val day = Day2()
    private val file = "day2"

    @Test
    fun part1() {
        assertEquals(2, day.part1(readFile("/y2024/${file}_sample.txt")))
        assertEquals(502, day.part1(readFile("/y2024/${file}.txt")))
    }

    @Test
    fun part2() {
        assertEquals(4, day.part2(readFile("/y2024/${file}_sample.txt")))
        assertEquals(0, day.part2(readFile("/y2024/${file}.txt")))
    }
}
