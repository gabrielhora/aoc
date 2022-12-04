defmodule Day4Test do
  use ExUnit.Case
  require Logger

  import Y2022.Day4

  @example File.read!("priv/Y2022/day4/example.txt") |> String.split("\n")
  @input File.read!("priv/Y2022/day4/input.txt") |> String.split("\n")

  test "part 1" do
    assert part1(@example) == 2
    Logger.info("Y2022.Day4.Part1: #{part1(@input)}")
  end

  test "part 2" do
    assert part2(@example) == 4
    Logger.info("Y2022.Day4.Part2: #{part2(@input)}")
  end
end
