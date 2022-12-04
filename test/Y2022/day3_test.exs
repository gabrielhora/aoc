defmodule Day3Test do
  use ExUnit.Case
  require Logger

  import Y2022.Day3

  @example File.read!("priv/Y2022/day3/example.txt") |> String.split("\n")
  @input File.read!("priv/Y2022/day3/input.txt") |> String.split("\n")

  test "part 1" do
    assert part1(@example) == 157
    Logger.info("Y2022.Day3.Part1: #{part1(@input)}")
  end

  test "part 2" do
    assert part2(@example) == 70
    Logger.info("Y2022.Day3.Part2: #{part2(@input)}")
  end
end
