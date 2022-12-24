defmodule Y2022.Day21 do
  def part1(input) do
    data = parse(input)
    solve(data, data["root"])
  end

  def part2(input) do
    data = parse(input)

    # find the side of the equation that doesn't change with "humn"
    {k1, k2, _} = data["root"]
    {humn, no_humn} = if find_human(data, data[k1]), do: {k1, k2}, else: {k2, k1}
    target = solve(data, data[no_humn])

    # search for the correct value by changing the side that changes with "humn"
    binary_search(
      data,
      data[humn],
      0,
      1_000_000_000_000_000,
      target
    )
  end

  defp binary_search(_input, _op, l, r, _target) when l > r, do: :not_found

  defp binary_search(input, op, l, r, target) do
    m = div(l + r, 2)
    input = Map.replace(input, "humn", m)
    res = solve(input, op)

    cond do
      res < target -> binary_search(input, op, l, m + 1, target)
      res > target -> binary_search(input, op, m - 1, r, target)
      true -> m
    end
  end

  defp find_human(_input, x) when is_integer(x), do: false
  defp find_human(_input, {"humn", _, _}), do: true
  defp find_human(_input, {_, "humn", _}), do: true

  defp find_human(input, {k1, k2, _}),
    do: find_human(input, input[k1]) || find_human(input, input[k2])

  defp solve(_input, x) when is_integer(x), do: x

  defp solve(input, {k1, k2, op}) do
    val1 = solve(input, input[k1])
    val2 = solve(input, input[k2])
    op.(val1, val2)
  end

  defp parse(input) do
    input
    |> Enum.map(&Regex.named_captures(~r/(?<name>.*): (?<op>.*)/, &1))
    |> Enum.map(fn m ->
      case Integer.parse(m["op"]) do
        :error ->
          {
            m["name"],
            case String.split(m["op"]) do
              [a, "+", b] -> {a, b, &Kernel.+/2}
              [a, "-", b] -> {a, b, &Kernel.-/2}
              [a, "*", b] -> {a, b, &Kernel.*/2}
              [a, "/", b] -> {a, b, &Kernel.div/2}
            end
          }

        {num, _} ->
          {m["name"], num}
      end
    end)
    |> Map.new()
  end
end
