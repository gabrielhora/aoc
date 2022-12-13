defmodule Day13Test do
  use ExUnit.Case

  import Y2022.Day13

  @example File.read!("priv/Y2022/day13/example.txt")
  @input File.read!("priv/Y2022/day13/input.txt")

  test "part 1" do
    assert part1(@example) == 13
    IO.puts("\ny2022 day13 part1: #{part1(@input)}")
  end

  test "part 2" do
    assert part2(@example) == 140
    IO.puts("\ny2022 day13 part2: #{part2(@input)}")
  end
end
