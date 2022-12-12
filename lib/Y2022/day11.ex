defmodule Y2022.Day11 do
  def part1(input) do
    monkeys = parse_monkeys(input)

    play_rounds(20, monkeys, &div(&1, 3))
    |> Enum.map(& &1.moves)
    |> Enum.sort(:desc)
    |> Enum.take(2)
    |> then(fn [a, b] -> a * b end)
  end

  def part2(input) do
    monkeys = parse_monkeys(input)

    # use the least common denominator to reduce the worry sizes
    divisor = monkeys |> Enum.map(& &1.test) |> lcm()

    play_rounds(10000, monkeys, &rem(&1, divisor))
    |> Enum.map(& &1.moves)
    |> Enum.sort(:desc)
    |> Enum.take(2)
    |> then(fn [a, b] -> a * b end)
  end

  defp play_rounds(0, monkeys, _worry_func), do: monkeys

  defp play_rounds(round, monkeys, worry_func) do
    monkeys = play_monkeys(0, monkeys, worry_func)
    play_rounds(round - 1, monkeys, worry_func)
  end

  defp play_monkeys(idx, monkeys, _worry_func) when idx == length(monkeys), do: monkeys

  defp play_monkeys(idx, monkeys, worry_func) do
    monkey = Enum.at(monkeys, idx)
    monkeys = move_items(monkey.items, monkey, idx, monkeys, worry_func)
    play_monkeys(idx + 1, monkeys, worry_func)
  end

  defp move_items([], _monkey, _monkey_idx, monkeys, _worry_func), do: monkeys

  defp move_items([worry | rest], monkey, monkey_idx, monkeys, worry_func) do
    worry = op(monkey.op, worry) |> worry_func.()

    target_idx =
      if rem(worry, monkey.test) == 0,
        do: monkey.if_true,
        else: monkey.if_false

    monkeys = move_item(monkeys, monkey_idx, target_idx, worry)
    move_items(rest, monkey, monkey_idx, monkeys, worry_func)
  end

  defp move_item(monkeys, from, to, item) do
    monkeys
    |> Enum.with_index()
    |> Enum.map(fn {m, i} ->
      cond do
        i == to -> %{m | items: [item | m.items]}
        i == from -> %{m | items: tl(m.items), moves: m.moves + 1}
        true -> m
      end
    end)
  end

  defp op({"old * old", _}, item), do: item * item
  defp op({"old * " <> _, num}, item), do: num * item
  defp op({"old + " <> _, num}, item), do: num + item

  defp parse_monkeys(input) do
    String.split(input, "\n\n") |> Enum.map(&monkey_data/1)
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
      if_false: caps["if_false"] |> String.to_integer(),
      moves: 0
    }
  end

  # https://rosettacode.org/wiki/Least_common_multiple#Elixir
  def lcm([a, b]), do: div(abs(a * b), Integer.gcd(a, b))
  def lcm([h | t]), do: lcm([h, lcm(t)])
end
