defmodule Day4 do
  def part1(lines) do
    lines
    |> Enum.map(&String.split(&1, ","))
    |> Enum.map(fn [e1, e2] -> [parse_range(e1), parse_range(e2)] end)
    |> Enum.filter(&either_contains/1)
    |> Enum.count()
  end

  def part2(lines) do
    lines
    |> Enum.map(&String.split(&1, ","))
    |> Enum.map(fn [e1, e2] -> [parse_range(e1), parse_range(e2)] end)
    |> Enum.filter(&overlap/1)
    |> Enum.count()
  end

  defp parse_range(range) do
    [n1, n2] = String.split(range, "-")
    {n1, _} = Integer.parse(n1)
    {n2, _} = Integer.parse(n2)
    Enum.to_list(n1..n2) |> MapSet.new()
  end

  defp either_contains([s1, s2]) do
    cond do
      MapSet.intersection(s1, s2) == s2 -> true
      MapSet.intersection(s2, s1) == s1 -> true
      true -> false
    end
  end

  defp overlap([s1, s2]) do
    MapSet.intersection(s1, s2) != %MapSet{}
  end
end
