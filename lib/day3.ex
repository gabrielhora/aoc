defmodule Day3 do
  def part1(lines) do
    lines
    |> Enum.map(&String.split_at(&1, div(String.length(&1), 2)))
    |> Enum.map(fn el ->
      fst = MapSet.new(String.graphemes(elem(el, 0)))
      snd = MapSet.new(String.graphemes(elem(el, 1)))
      MapSet.intersection(fst, snd) |> Enum.at(0)
    end)
    |> Enum.map(&score/1)
    |> Enum.sum()
  end

  def part2(lines) do
    lines
    |> Enum.chunk_every(3)
    |> Enum.map(fn el -> Enum.map(el, &MapSet.new(String.graphemes(&1))) end)
    |> Enum.map(fn el -> Enum.reduce(el, &MapSet.intersection(&1, &2)) |> Enum.at(0) end)
    |> Enum.map(&score/1)
    |> Enum.sum()
  end

  defp score(letter) do
    asc = :binary.first(letter)
    if asc in ?a..?z, do: asc - 96, else: asc - 64 + 26
  end
end
