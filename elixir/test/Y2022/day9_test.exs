defmodule Day9Test do
  use ExUnit.Case

  import Y2022.Day9

  @example File.read!("priv/Y2022/day9/example.txt") |> String.split("\n")
  @example2 File.read!("priv/Y2022/day9/example_part2.txt") |> String.split("\n")
  @input File.read!("priv/Y2022/day9/input.txt") |> String.split("\n")

  test "part 1" do
    assert part1(@example) == 13
    IO.puts("\ny2022 day9 part1: #{part1(@input)}")
  end

  test "part 2" do
    assert part2(@example2) == 36
    IO.puts("\n#y2022 day9 part2: #{part2(@input)}")
  end
end
