defmodule Day3Test do
  use ExUnit.Case

  import Y2022.Day3

  @example File.read!("priv/Y2022/day3/example.txt") |> String.split("\n")
  @input File.read!("priv/Y2022/day3/input.txt") |> String.split("\n")

  test "part 1" do
    assert part1(@example) == 157
    IO.puts("\y2022 day3 part1: #{part1(@input)}")
  end

  test "part 2" do
    assert part2(@example) == 70
    IO.puts("\n#y2022 day3 part2: #{part2(@input)}")
  end
end
