defmodule Day11Test do
  use ExUnit.Case

  import Y2022.Day11

  @example File.read!("priv/Y2022/day11/example.txt")
  @input File.read!("priv/Y2022/day11/input.txt")

  test "part 1" do
    assert part1(@example) == 10605
    IO.puts("\ny2022 day11 part1: #{part1(@input)}")
  end

  test "part 2" do
    assert part2(@example) == 2_713_310_158
    IO.puts("\ny2022 day11 part2: #{part2(@input)}")
  end
end
