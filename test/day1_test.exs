defmodule Day1Test do
  use ExUnit.Case
  require Logger

  @example File.read!("priv/day1/example.txt") |> String.split("\n")
  @input File.read!("priv/day1/input.txt") |> String.split("\n")

  test "part 1" do
    assert Day1.part1(@example) == 24000
    Logger.info("Day 1, Part 1: #{Day1.part1(@input)}")
  end

  test "part 2" do
    assert Day1.part2(@example) == 45000
    Logger.info("Day 1, Part 2: #{Day1.part2(@input)}")
  end
end
