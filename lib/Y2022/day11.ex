defmodule Y2022.Day11 do
  def part1(input) do
    monkeys = parse_monkeys(input)

    play_rounds(20, monkeys, List.duplicate(0, length(monkeys)), &div(&1, 3))
    |> elem(1)
    |> Enum.sort(:desc)
    |> Enum.take(2)
    |> then(fn [a, b] -> a * b end)
  end

  def part2(input) do
    monkeys = parse_monkeys(input)

    # use the least common denominator to reduce the worry sizes
    divisor = monkeys |> Enum.map(fn {m, _} -> m.test end) |> lcm()

    play_rounds(10000, monkeys, List.duplicate(0, length(monkeys)), &rem(&1, divisor))
    |> elem(1)
    |> Enum.sort(:desc)
    |> Enum.take(2)
    |> then(fn [a, b] -> a * b end)
  end

  defp play_rounds(0, monkeys, moves, _worry_func), do: {monkeys, moves}

  defp play_rounds(round, monkeys, moves, worry_func) do
    {monkeys, moves} = play_monkeys(0, monkeys, moves, worry_func)
    play_rounds(round - 1, monkeys, moves, worry_func)
  end

  defp play_monkeys(idx, monkeys, moves, _worry_func) when idx == length(monkeys),
    do: {monkeys, moves}

  defp play_monkeys(idx, monkeys, moves, worry_func) do
    {monkey, idx} = Enum.at(monkeys, idx)
    {monkeys, moves} = move_items(monkey.items, monkey, idx, monkeys, moves, worry_func)
    play_monkeys(idx + 1, monkeys, moves, worry_func)
  end

  defp move_items([worry | rest], monkey, monkey_idx, monkeys, moves, worry_func) do
    worry = op(monkey.op, worry) |> worry_func.()

    target_idx =
      if rem(worry, monkey.test) == 0,
        do: monkey.if_true,
        else: monkey.if_false

    moves = List.replace_at(moves, monkey_idx, (Enum.at(moves, monkey_idx) || 0) + 1)
    monkeys = move_item(monkeys, monkey_idx, target_idx, worry)
    move_items(rest, monkey, monkey_idx, monkeys, moves, worry_func)
  end

  defp move_items([], _monkey, _monkey_idx, monkeys, moves, _worry_func), do: {monkeys, moves}

  defp move_item(monkeys, from, to, item) do
    monkeys
    |> Enum.map(fn {m, i} ->
      cond do
        i == to -> {%{m | items: [item | m.items]}, i}
        i == from -> {%{m | items: tl(m.items)}, i}
        true -> {m, i}
      end
    end)
  end

  defp op({"old * old", _}, item), do: item * item
  defp op({"old * " <> _, num}, item), do: num * item
  defp op({"old + " <> _, num}, item), do: num + item

  defp parse_monkeys(input) do
    String.split(input, "\n\n")
    |> Enum.map(&monkey_data/1)
    |> Enum.with_index()
  end

  defp monkey_data(monkey) do
    caps =
      Regex.named_captures(
        ~r"Starting items: (?<items>.+)
  Operation: new = (?<op>.+)
  Test: divisible by (?<test>\d+)
    If true: throw to monkey (?<if_true>\d+)
    If false: throw to monkey (?<if_false>\d+)",
        monkey
      )

    op_match = Regex.named_captures(~r/old . (?<num>\d+)/, caps["op"])

    %{
      items:
        caps["items"]
        |> String.split(",")
        |> Enum.map(&String.trim/1)
        |> Enum.map(&String.to_integer/1),
      op:
        if(op_match == nil,
          do: {caps["op"], 0},
          else: {caps["op"], String.to_integer(op_match["num"])}
        ),
      test: caps["test"] |> String.to_integer(),
      if_true: caps["if_true"] |> String.to_integer(),
      if_false: caps["if_false"] |> String.to_integer()
    }
  end

  # https://rosettacode.org/wiki/Least_common_multiple#Elixir
  def lcm([a, b]), do: div(abs(a * b), Integer.gcd(a, b))
  def lcm([h | t]), do: lcm([h, lcm(t)])
end
