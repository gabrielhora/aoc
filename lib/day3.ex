defmodule Day3 do
  def part1(lines) do
    lines
    |> Enum.map(&String.split_at(&1, div(String.length(&1), 2)))
    |> Enum.map(&{String.graphemes(elem(&1, 0)), String.graphemes(elem(&1, 1))})
    |> Enum.map(&[MapSet.new(elem(&1, 0)), MapSet.new(elem(&1, 1))])
    |> Enum.map(&MapSet.intersection(Enum.at(&1, 0), Enum.at(&1, 1)))
    |> Enum.map(&score/1)
    |> Enum.sum()
  end

  def part2(lines) do
    lines
    |> Enum.chunk_every(3)
    |> Enum.map(fn el -> Enum.map(el, &MapSet.new(String.graphemes(&1))) end)
    |> Enum.map(fn el -> Enum.reduce(el, &MapSet.intersection(&1, &2)) end)
    |> Enum.map(&score/1)
    |> Enum.sum()
  end

  defp score(map) when is_map(map) do
    MapSet.to_list(map)
    |> List.first()
    |> :binary.first()
    |> score()
  end

  defp score(letter) when is_integer(letter) do
    if(letter in ?a..?z, do: letter - 96, else: letter - 64 + 26)
  end
end
