defmodule Day1 do
  def part1(lines) do
    lines
    |> Enum.chunk_while([], &chunk_on_space/2, &{:cont, &1})
    |> Enum.reduce([], &[Enum.sum(&1) | &2])
    |> Enum.max()
  end

  def part2(lines) do
    lines
    |> Enum.chunk_while([], &chunk_on_space/2, &{:cont, &1})
    |> Enum.reduce([], &[Enum.sum(&1) | &2])
    |> Enum.sort(:desc)
    |> Enum.take(3)
    |> Enum.sum()
  end

  defp chunk_on_space("", acc), do: {:cont, acc, []}
  defp chunk_on_space(el, acc), do: {:cont, [String.to_integer(el) | acc]}
end
