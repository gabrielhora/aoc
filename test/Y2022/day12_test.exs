defmodule Day12Test do
  use ExUnit.Case

  import Y2022.Day12

  @example File.read!("priv/Y2022/day12/example.txt")
  @input File.read!("priv/Y2022/day12/input.txt")

  test "part 1" do
    assert part1(@example) == 31
    IO.puts("\ny2022 day12 part1: #{part1(@input)}")
  end

  test "part 2" do
    assert part2(@example) == 29
    IO.puts("\ny2022 day12 part2: #{part2(@input)}")
  end
end
