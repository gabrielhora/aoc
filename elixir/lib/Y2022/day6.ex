defmodule Y2022.Day6 do
  def part1(line) do
    line |> String.graphemes() |> first_unique_set_index(4)
  end

  def part2(line) do
    line |> String.graphemes() |> first_unique_set_index(14)
  end

  defp first_unique_set_index(chars, size) do
    chars
    |> Enum.chunk_every(size, 1, :discard)
    |> Enum.find_index(&(Enum.uniq(&1) |> length() == size))
    |> then(&(&1 + size))
  end
end
