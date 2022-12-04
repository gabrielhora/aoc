defmodule Day2Test do
  use ExUnit.Case
  require Logger

  import Y2022.Day2

  @example File.read!("priv/Y2022/day2/example.txt") |> String.split("\n")
  @input File.read!("priv/Y2022/day2/input.txt") |> String.split("\n")

  test "part 1" do
    assert part1(@example) == 15
    Logger.info("Y2022.Day2.Part1: #{part1(@input)}")
  end

  test "part 2" do
    assert part2(@example) == 12
    Logger.info("Y2022.Day2.Part2: #{part2(@input)}")
  end
end
