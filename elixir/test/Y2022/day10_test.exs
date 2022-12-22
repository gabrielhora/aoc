defmodule Day10Test do
  use ExUnit.Case

  import Y2022.Day10

  @example File.read!("priv/Y2022/day10/example.txt") |> String.split("\n")
  @input File.read!("priv/Y2022/day10/input.txt") |> String.split("\n")

  test "part 1" do
    assert part1(@example) == 13140
    IO.puts("\ny2022 day10 part1: #{part1(@input)}")
  end

  test "part 2" do
    assert part2(@example) ==
             "##  ##  ##  ##  ##  ##  ##  ##  ##  ##  \n###   ###   ###   ###   ###   ###   ### \n####    ####    ####    ####    ####    \n#####     #####     #####     #####     \n######      ######      ######      ####\n#######       #######       #######     \n "

    IO.puts("\ny2022 day10 part2:\n #{part2(@input)}")
  end
end
