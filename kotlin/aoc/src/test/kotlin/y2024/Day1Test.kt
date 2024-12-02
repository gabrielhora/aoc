package y2024

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import utils.readFile

class Day1Test {
    private val day = Day1()

    @Test
    fun part1() {
        assertEquals(day.part1(readFile("/y2024/day1_sample.txt")), 11)
        assertEquals(day.part1(readFile("/y2024/day1.txt")), 1765812)
    }

    @Test
    fun part2() {
        assertEquals(day.part2(readFile("/y2024/day1_sample.txt")), 31)
        assertEquals(day.part2(readFile("/y2024/day1.txt")), 20520794)
    }
}
