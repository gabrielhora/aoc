defmodule Y2022.Day6 do
  def part1(line) do
    line |> String.graphemes() |> first_unique_index(4)
  end

  def part2(line) do
    line |> String.graphemes() |> first_unique_index(14)
  end

  defp first_unique_index(chars, size) do
    try do
      for i <- 0..(length(chars) - 1) do
        start = Enum.slice(chars, i, size) |> Enum.uniq()
        if length(start) == size, do: throw(i + size)
      end
    catch
      index -> index
    end
  end
end
