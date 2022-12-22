defmodule Y2022.Day13 do
  def part1(input) do
    input
    |> String.split("\n\n")
    |> Enum.map(&parse_packet/1)
    |> Enum.with_index(1)
    |> Enum.filter(fn {[l, r], _} -> cmp(l, r) end)
    |> Enum.map(&elem(&1, 1))
    |> Enum.sum()
  end

  def part2(input) do
    input
    |> String.split("\n\n")
    |> Enum.flat_map(&parse_packet/1)
    |> then(&(&1 ++ [[[2]], [[6]]]))
    |> Enum.sort(&cmp/2)
    |> Enum.with_index(1)
    |> Enum.filter(fn {p, _} -> p == [[2]] || p == [[6]] end)
    |> Enum.map(&elem(&1, 1))
    |> Enum.reduce(&Kernel.*/2)
  end

  # compare returns true if the lists are in the correct order

  def cmp(left, right) do
    case {left, right} do
      {[], r} when is_list(r) -> true
      {l, []} when is_list(l) -> false
      {[l | _], [r | _]} when l < r and is_integer(l) and is_integer(r) -> true
      {[l | _], [r | _]} when l > r and is_integer(l) and is_integer(r) -> false
      {[l | ls], [r | rs]} when is_integer(l) and is_integer(r) -> cmp(ls, rs)
      {[l | ls], r} when is_integer(l) and is_list(r) -> cmp([[l] | ls], r)
      {l, [r | rs]} when is_integer(r) and is_list(l) -> cmp(l, [[r] | rs])
      {[l | ls], [r | rs]} when l == r and is_list(l) and is_list(r) -> cmp(ls, rs)
      {[l | _], [r | _]} when is_list(l) and is_list(r) -> cmp(l, r)
    end
  end

  defp parse_packet(line) do
    String.split(line, "\n")
    |> then(fn [p1, p2] -> [parse_list(p1), parse_list(p2)] end)
  end

  defp parse_list(list), do: Code.eval_string(list) |> elem(0)
end
