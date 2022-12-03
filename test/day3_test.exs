defmodule Day3Test do
  use ExUnit.Case
  require Logger

  @example File.read!("priv/day3/example.txt") |> String.split("\n")
  @input File.read!("priv/day3/input.txt") |> String.split("\n")

  test "part 1" do
    assert Day3.part1(@example) == 157
    Logger.info("Day 3, Part 1: #{Day3.part1(@input)}")
  end

  test "part 2" do
    assert Day3.part2(@example) == 70
    Logger.info("Day 3, Part 2: #{Day3.part2(@input)}")
  end
end
