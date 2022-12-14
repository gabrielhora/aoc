defmodule Y2022.Day14 do
  def part1(input) do
    blocked = parse_input(input)
    y_limit = blocked |> Enum.max_by(fn {_, y} -> y end) |> elem(1)
    origin = {500, 0}

    Stream.iterate(0, &(&1 + 1))
    |> Enum.reduce_while(blocked, fn n, blocked ->
      case drop(blocked, origin, y_limit) do
        {:ok, blocked} -> {:cont, blocked}
        {:void, _} -> {:halt, n}
      end
    end)
  end

  def part2(input) do
    blocked = parse_input(input)
    y_limit = blocked |> Enum.max_by(fn {_, y} -> y end) |> then(&(elem(&1, 1) + 2))
    origin = {500, 0}

    Stream.iterate(0, &(&1 + 1))
    |> Enum.reduce_while(blocked, fn n, blocked ->
      {_, blocked} = drop(blocked, origin, y_limit)
      if !free?(blocked, origin), do: {:halt, n + 1}, else: {:cont, blocked}
    end)
  end

  defp drop(blocked, {x, y}, y_limit) do
    down = {x, y + 1}
    left = {x - 1, y + 1}
    right = {x + 1, y + 1}

    cond do
      y >= y_limit -> {:void, MapSet.put(blocked, {x, y - 1})}
      free?(blocked, down) -> drop(blocked, down, y_limit)
      free?(blocked, left) -> drop(blocked, left, y_limit)
      free?(blocked, right) -> drop(blocked, right, y_limit)
      true -> {:ok, MapSet.put(blocked, {x, y})}
    end
  end

  defp free?(blocked, coords), do: !Enum.member?(blocked, coords)

  defp parse_input(input) do
    input
    |> Enum.map(&String.split(&1, " -> "))
    |> Enum.map(fn coord ->
      coord
      |> Enum.map(&String.split(&1, ","))
      |> Enum.map(fn [x, y] -> {String.to_integer(x), String.to_integer(y)} end)
    end)
    |> Enum.flat_map(&Enum.chunk_every(&1, 2, 1, :discard))
    |> Enum.flat_map(fn [{x1, y1}, {x2, y2}] ->
      for x <- x1..x2, y <- y1..y2, do: {x, y}
    end)
    |> MapSet.new()
  end
end
