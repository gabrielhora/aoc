defmodule Y2022.Day1Test do
  use ExUnit.Case

  import Y2022.Day1

  @example File.read!("priv/Y2022/day1/example.txt") |> String.split("\n")
  @input File.read!("priv/Y2022/day1/input.txt") |> String.split("\n")

  test "part 1" do
    assert part1(@example) == 24000
    IO.puts("\ny2022 day1 part1: #{part1(@input)}")
  end

  test "part 2" do
    assert part2(@example) == 45000
    IO.puts("\n#y2022 day1 part2: #{part2(@input)}")
  end
end
