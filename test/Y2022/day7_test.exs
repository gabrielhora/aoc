defmodule Day7Test do
  use ExUnit.Case

  import Y2022.Day7

  @example File.read!("priv/Y2022/day7/example.txt") |> String.split("\n")
  @input File.read!("priv/Y2022/day7/input.txt") |> String.split("\n")

  test "part 1" do
    assert part1(@example) == 95437
    #   IO.puts("\ny2022 day7 part1: #{part1(@input)}")
  end

  test "part 2" do
    assert part2(@example) == 24_933_642
    IO.puts("\n#y2022 day7 part2: #{part2(@input)}")
  end
end
