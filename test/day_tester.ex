defmodule DayTester do
  defmacro __using__(mod: mod, part1: part1, part2: part2) do
    quote do
      use ExUnit.Case
      import unquote(mod)

      [_, year, day] = String.split(Atom.to_string(unquote(mod)), ".")
      @path "#{year}/#{String.downcase(day)}"

      @example File.read!("priv/#{@path}/example.txt") |> String.split("\n")
      @input File.read!("priv/#{@path}/input.txt") |> String.split("\n")

      test "part 1" do
        assert part1(@example) == unquote(part1)
        IO.puts("\n#{@path}/part1: #{part1(@input)}")
      end

      test "part 2" do
        assert part2(@example) == unquote(part2)
        IO.puts("\n#{@path}/part2: #{part2(@input)}")
      end
    end
  end
end
