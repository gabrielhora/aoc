defmodule Day3Test do
  use ExUnit.Case
  require Logger

  @example File.read!("priv/day4/example.txt") |> String.split("\n")
  @input File.read!("priv/day4/input.txt") |> String.split("\n")

  test "part 1" do
    assert Day4.part1(@example) == 2
    Logger.info("Day 4, Part 1: #{Day4.part1(@input)}")
  end

  test "part 2" do
    assert Day4.part2(@example) == 4
    Logger.info("Day 4, Part 2: #{Day4.part2(@input)}")
  end
end
