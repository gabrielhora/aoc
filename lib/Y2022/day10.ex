defmodule Y2022.Day10 do
  def part1(lines) do
    cycles = execute(lines)

    20 * Enum.at(cycles, 20 - 1) +
      60 * Enum.at(cycles, 60 - 1) +
      100 * Enum.at(cycles, 100 - 1) +
      140 * Enum.at(cycles, 140 - 1) +
      180 * Enum.at(cycles, 180 - 1) +
      220 * Enum.at(cycles, 220 - 1)
  end

  def part2(lines) do
    execute(lines)
    |> Enum.with_index()
    |> Enum.map(fn {x, cycle} ->
      x = div(cycle, 40) * 40 + x

      if x - 1 == cycle || x == cycle || x + 1 == cycle,
        do: "#",
        else: " "
    end)
    |> Enum.chunk_every(40)
    |> Enum.map(&Enum.join/1)
    |> Enum.join("\n")
  end

  defp execute(instructions) do
    Enum.reduce(instructions, {1, [1]}, fn cmd, {x, xs} ->
      {new_x, new_xs} = op(cmd, x)
      {new_x, xs ++ new_xs}
    end)
    |> elem(1)
  end

  defp op("noop", x), do: {x, [x]}

  defp op("addx " <> num, x) do
    new_x = x + String.to_integer(num)
    {new_x, [x, new_x]}
  end
end
