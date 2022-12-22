defmodule Day5Test do
  use ExUnit.Case

  import Y2022.Day5

  @example File.read!("priv/Y2022/day5/example.txt")
  @input File.read!("priv/Y2022/day5/input.txt")

  test "part 1" do
    assert part1(@example) == "CMZ"
    IO.puts("\ny2022 day5 part1: #{part1(@input)}")
  end

  test "part 2" do
    assert part2(@example) == "MCD"
    IO.puts("\n#y2022 day5 part2: #{part2(@input)}")
  end
end
