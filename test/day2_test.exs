defmodule Day2Test do
  use ExUnit.Case
  require Logger

  @example File.read!("priv/day2/example.txt") |> String.split("\n")
  @input File.read!("priv/day2/input.txt") |> String.split("\n")

  test "part 1" do
    assert Day2.part1(@example) == 15
    Logger.info("Day 2, Part 1: #{Day2.part1(@input)}")
  end

  test "part 2" do
    assert Day2.part2(@example) == 12
    Logger.info("Day 1, Part 2: #{Day2.part2(@input)}")
  end
end
