defmodule Day4Test do
  use ExUnit.Case

  import Y2022.Day4

  @example File.read!("priv/Y2022/day4/example.txt") |> String.split("\n")
  @input File.read!("priv/Y2022/day4/input.txt") |> String.split("\n")

  test "part 1" do
    assert part1(@example) == 2
    IO.puts("\y2022 day4 part1: #{part1(@input)}")
  end

  test "part 2" do
    assert part2(@example) == 4
    IO.puts("\n#y2022 day4 part2: #{part2(@input)}")
  end
end
