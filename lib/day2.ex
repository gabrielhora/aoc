defmodule Day2 do
  def part1(lines) do
    lines
    |> Enum.map(&String.split/1)
    |> Enum.reject(&(&1 == []))
    |> Enum.map(&score_part1/1)
    |> Enum.sum()
  end

  defp score_part1(["A", "X"]), do: 1 + 3
  defp score_part1(["A", "Y"]), do: 2 + 6
  defp score_part1(["A", "Z"]), do: 3 + 0
  defp score_part1(["B", "X"]), do: 1 + 0
  defp score_part1(["B", "Y"]), do: 2 + 3
  defp score_part1(["B", "Z"]), do: 3 + 6
  defp score_part1(["C", "X"]), do: 1 + 6
  defp score_part1(["C", "Y"]), do: 2 + 0
  defp score_part1(["C", "Z"]), do: 3 + 3

  def part2(lines) do
    lines
    |> Enum.map(&String.split/1)
    |> Enum.reject(&(&1 == []))
    |> Enum.map(&score_part2/1)
    |> Enum.sum()
  end

  defp score_part2(["A", "X"]), do: 3 + 0
  defp score_part2(["A", "Y"]), do: 1 + 3
  defp score_part2(["A", "Z"]), do: 2 + 6
  defp score_part2(["B", "X"]), do: 1 + 0
  defp score_part2(["B", "Y"]), do: 2 + 3
  defp score_part2(["B", "Z"]), do: 3 + 6
  defp score_part2(["C", "X"]), do: 2 + 0
  defp score_part2(["C", "Y"]), do: 3 + 3
  defp score_part2(["C", "Z"]), do: 1 + 6
end
