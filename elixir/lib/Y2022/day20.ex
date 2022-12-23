defmodule Y2022.Day20 do
  def part1(input) do
    input
    |> Enum.map(&String.to_integer/1)
    |> Enum.with_index()
    |> then(&mix({&1, &1}))
    |> then(fn {_, mixed} -> solve(mixed) end)
  end

  def part2(input) do
    input
    |> Enum.map(&String.to_integer/1)
    |> Enum.map(&(&1 * 811_589_153))
    |> Enum.with_index()
    |> then(fn data ->
      Enum.reduce(0..9, {data, data}, fn _, acc -> mix(acc) end)
    end)
    |> then(fn {_, mixed} -> solve(mixed) end)
  end

  defp solve(data) do
    size = length(data)
    zero = Enum.find_index(data, fn {val, _} -> val == 0 end)
    a = data |> Enum.at(mod(zero + 1000, size)) |> elem(0)
    b = data |> Enum.at(mod(zero + 2000, size)) |> elem(0)
    c = data |> Enum.at(mod(zero + 3000, size)) |> elem(0)
    a + b + c
  end

  defp mix({original, mixed}) do
    size = length(original)

    mixed =
      original
      |> Enum.reduce(mixed, fn d = {val, _}, acc ->
        cur_idx = Enum.find_index(acc, &(&1 == d))
        new_idx = mod(cur_idx + val, size - 1)
        new_idx = if(new_idx == 0, do: -1, else: new_idx)
        {_, acc} = List.pop_at(acc, cur_idx)
        List.insert_at(acc, new_idx, d)
      end)

    {original, mixed}
  end

  # https://stackoverflow.com/a/858649
  defp mod(x, y) when x > 0, do: rem(x, y)
  defp mod(x, y) when x < 0, do: rem(x, y) + y
  defp mod(0, _), do: 0
end
