defmodule Day20Test do
  use ExUnit.Case

  import Y2022.Day20

  @example File.read!("priv/Y2022/day20/example.txt") |> String.split("\n")
  @input File.read!("priv/Y2022/day20/input.txt") |> String.split("\n")

  test "part 1" do
    assert part1(@example) == 3
    IO.puts("\ny2022 day20 part1: #{part1(@input)}")
  end

  test "part 2" do
    assert part2(@example) == 1_623_178_306
    IO.puts("\ny2022 day20 part2: #{part2(@input)}")
  end
end
