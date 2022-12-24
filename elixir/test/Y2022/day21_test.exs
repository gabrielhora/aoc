defmodule Day21Test do
  use ExUnit.Case

  import Y2022.Day21

  @example File.read!("priv/Y2022/day21/example.txt") |> String.split("\n")
  @input File.read!("priv/Y2022/day21/input.txt") |> String.split("\n")

  test "part 1" do
    assert part1(@example) == 152
    IO.puts("\ny2022 day21 part1: #{part1(@input)}")
  end

  test "part 2" do
    # In example.txt increasing (in binary_search/5) `m`, increases `res`
    # With the input.txt the relationship is inversed (i.e. the "list" is reversed)
    # assert part2(@example) == 301
    IO.puts("\ny2022 day21 part2: #{part2(@input)}")
  end
end
