defmodule Y2022.Day15 do
  # FIXME: Still haven't figure out a way to solve for the input data without looping through everything.
  #  Part 1 is slow but manageble and computing a grid 4_000_000 by 4_000_000 won't work for part 2.
  #  I'll come back to this later.

  def part1(input, row, range) do
    data = parse(input)
    beacons = data |> Enum.map(&elem(&1, 1)) |> MapSet.new()

    range
    |> Enum.reduce(MapSet.new(), fn x, acc ->
      coords =
        for {s, _, d} <- data do
          if m_dist(s, {x, row}) <= d do
            {x, row}
          end
        end
        |> MapSet.new()

      MapSet.union(acc, coords)
    end)
    |> then(&MapSet.difference(&1, beacons))
    |> then(&MapSet.reject(&1, fn x -> x == nil end))
    |> MapSet.size()
  end

  def part2(input, space) do
    data = parse(input)

    0..space
    |> Enum.find_value(fn x ->
      ys = find_ys(x, space, data)

      if MapSet.size(ys) != space + 1 do
        # find missing y and halt
        y = 0..space |> Enum.find(fn y -> !MapSet.member?(ys, {x, y}) end)
        {x, y}
      end
    end)
    |> then(fn {x, y} -> x * 4_000_000 + y end)
  end

  defp find_ys(x, space, data) do
    0..space
    |> Enum.reduce(MapSet.new(), fn y, acc ->
      coords =
        for {s, _, d} <- data do
          if m_dist(s, {x, y}) <= d do
            {x, y}
          end
        end
        |> MapSet.new()

      MapSet.union(acc, coords)
    end)
    |> then(&MapSet.reject(&1, fn x -> x == nil end))
  end

  defp m_dist({ax, ay}, {bx, by}), do: abs(ax - bx) + abs(ay - by)

  defp parse(input) do
    input
    |> Enum.map(&parse_line/1)
    |> Enum.map(fn {s, b} -> {s, b, m_dist(s, b)} end)
  end

  defp parse_line(line) do
    Regex.named_captures(
      ~r/Sensor at x=(?<sx>-?\d+), y=(?<sy>-?\d+): closest beacon is at x=(?<bx>-?\d+), y=(?<by>-?\d+)/,
      line
    )
    |> then(fn m ->
      {{String.to_integer(m["sx"]), String.to_integer(m["sy"])},
       {String.to_integer(m["bx"]), String.to_integer(m["by"])}}
    end)
  end
end
