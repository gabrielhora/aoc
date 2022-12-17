defmodule Y2022.Day15 do
  @doc """
  That was freaking hard! But I'm very happy I figured it out :D Code could probably be better but
  that will do for now (already spend too much time here)!

  My solution is to calculate the "signal strength" of the sensors at a particular Y row. This takes
  the format of a range from a..b which are the X coordinates that one sensor "senses" in that Y
  row. After that I merge all sensor strength into the the minimum amount of ranges possible
  (important for part 2).

  For part 1 we know the Y row we are interested in, so we calculate the "signal strength" of all
  sensors at that Y get the size of the range.

  For part 2 we need to find a hole somewhere in a 4mil x 4mil grid. I do that by looping through
  the Y axis, calculating the signal strength of all sensors at that Y and looking for the hole,
  i.e. if the calculation of the signal strenth returns two different ranges.

  The second part is still a bit slow but very doable, it runs in about 9 seconds on my machine.
  """

  def part1(input, row) do
    data = parse(input)
    sensors = data |> Enum.map(fn {s, _, d} -> {s, d} end)

    sensors
    |> Enum.reduce([], fn {s, d}, acc -> [sensor_x_range(s, d, row) | acc] end)
    |> Enum.reject(&(&1 == nil))
    |> then(&Enum.sort/1)
    |> then(&merge_ranges(&1, -1_000_000_000..1_000_000_000))
    |> then(fn [a..b] -> b - a end)
  end

  def part2(input, space) do
    data = parse(input)
    sensors = data |> Enum.map(fn {s, _, d} -> {s, d} end)

    0..space
    |> Enum.find_value(fn y ->
      ranges =
        sensors
        |> Enum.reduce([], fn {s, d}, acc -> [sensor_x_range(s, d, y) | acc] end)
        |> Enum.reject(&(&1 == nil))
        |> then(&Enum.sort/1)
        |> then(&merge_ranges(&1, 0..space))

      if ranges != [0..space] do
        [_..x, _] = ranges
        {x + 1, y}
      end
    end)
    |> then(fn {x, y} -> x * 4_000_000 + y end)
  end

  # reduce while we can't reduce anymore
  defp merge_ranges(ranges = [h | t], max_range) when is_list(ranges) do
    new_ranges =
      Enum.reduce(t, [h], fn r, acc ->
        m = merge_ranges(hd(acc), r, max_range)
        List.flatten(Enum.reverse(m), tl(acc))
      end)

    result =
      if length(new_ranges) < length(ranges),
        do: merge_ranges(new_ranges, max_range),
        else: new_ranges

    Enum.sort(result)
  end

  defp merge_ranges(a = a1..a2, b = b1..b2, m1..m2) do
    a1 = max(a1, m1)
    a2 = min(a2, m2)
    b1 = max(b1, m1)
    b2 = min(b2, m2)

    cond do
      # overlap, merge them
      !Range.disjoint?(a, b) -> [min(a1, b1)..max(a2, b2)]
      # ends are "touching", merge them
      a2 < b1 && abs(a2 - b1) == 1 -> [a1..b2]
      b2 < a1 && abs(b2 - a1) == 1 -> [b1..a2]
      # no overlap, sort them
      true -> [min(a, b), max(a, b)]
    end
  end

  defp sensor_x_range({sx, sy}, dist, row_y) do
    dy = abs(row_y - sy)

    if dy > dist,
      # sensor does not reach this row_y
      do: nil,
      else: (sx - (dist - dy))..(sx + (dist - dy))
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
