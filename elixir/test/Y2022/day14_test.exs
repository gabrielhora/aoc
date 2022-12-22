defmodule Day14Test do
  use ExUnit.Case

  import Y2022.Day14

  @example File.read!("priv/Y2022/day14/example.txt") |> String.split("\n")
  @input File.read!("priv/Y2022/day14/input.txt") |> String.split("\n")

  test "part 1" do
    assert part1(@example) == 24
    IO.puts("\ny2022 day14 part1: #{part1(@input)}")
  end

  test "part 2" do
    assert part2(@example) == 93
    IO.puts("\ny2022 day14 part2: #{part2(@input)}")
  end
end
