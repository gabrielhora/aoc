defmodule Y2022.Day13 do
  def part1(input) do
    input
    |> String.split("\n\n")
    |> Enum.map(&parse_packet/1)
    |> Enum.with_index(1)
    |> Enum.filter(fn {[l, r], _} -> compare(l, r) end)
    |> Enum.map(&elem(&1, 1))
    |> Enum.sum()
  end

  def part2(input) do
    input
    |> String.split("\n\n")
    |> Enum.flat_map(&parse_packet/1)
    |> then(&(&1 ++ [[[2]], [[6]]]))
    |> Enum.sort(&compare/2)
    |> Enum.with_index(1)
    |> Enum.filter(fn {p, _} -> p == [[2]] || p == [[6]] end)
    |> Enum.map(&elem(&1, 1))
    |> Enum.reduce(&Kernel.*/2)
  end

  # compare returns true if the lists are in the correct order

  defp compare([left | lt], [right | rt]) when is_integer(left) and is_integer(right) do
    cond do
      left < right -> true
      left > right -> false
      true -> compare(lt, rt)
    end
  end

  defp compare([], right) when is_list(right), do: true
  defp compare(left, []) when is_list(left), do: false

  defp compare([left | lt], [right | rt]) when is_list(left) and is_list(right) do
    cond do
      left == right -> compare(lt, rt)
      true -> compare(left, right)
    end
  end

  defp compare([left | t], right) when is_integer(left), do: compare([[left] | t], right)
  defp compare(left, [right | t]) when is_integer(right), do: compare(left, [[right] | t])

  defp parse_packet(line) do
    String.split(line, "\n")
    |> then(fn [p1, p2] -> [parse_list(p1), parse_list(p2)] end)
  end

  defp parse_list(list), do: Code.eval_string(list) |> elem(0)
end
