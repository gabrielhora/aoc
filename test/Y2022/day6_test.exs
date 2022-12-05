defmodule Day6Test do
  use ExUnit.Case

  import Y2022.Day6

  @example File.read!("priv/Y2022/day6/example.txt")
  @input File.read!("priv/Y2022/day6/input.txt")

  test "part 1" do
    assert part1(@example) == 0
    IO.puts("\ny2022 day6 part1: #{part1(@input)}")
  end

  test "part 2" do
    assert part2(@example) == 0
    IO.puts("\n#y2022 day6 part2: #{part2(@input)}")
  end
end
