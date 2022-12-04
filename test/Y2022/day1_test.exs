defmodule Y2022.Day1Test do
  use ExUnit.Case
  require Logger

  import Y2022.Day1

  @example File.read!("priv/Y2022/day1/example.txt") |> String.split("\n")
  @input File.read!("priv/Y2022/day1/input.txt") |> String.split("\n")

  test "part 1" do
    assert part1(@example) == 24000
    Logger.info("Y2022.Day1.Part1: #{part1(@input)}")
  end

  test "part 2" do
    assert part2(@example) == 45000
    Logger.info("Y2022.Day1.Part2: #{part2(@input)}")
  end
end
