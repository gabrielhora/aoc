defmodule Day15Test do
  use ExUnit.Case

  import Y2022.Day15

  @example File.read!("priv/Y2022/day15/example.txt") |> String.split("\n")
  @input File.read!("priv/Y2022/day15/input.txt") |> String.split("\n")

  @tag timeout: :infinity
  test "part 1" do
    assert part1(@example, 10, -100..100) == 26
    # FIXME: This is slow!
    # IO.puts("\ny2022 day15 part1: #{part1(@input, 2_000_000, -10_000_000..10_000_000)}")
  end

  @tag timeout: :infinity
  test "part 2" do
    assert part2(@example, 20) == 56_000_011
    # IO.puts("\ny2022 day15 part2: #{part2(@input, 4_000_000)}")
  end
end
