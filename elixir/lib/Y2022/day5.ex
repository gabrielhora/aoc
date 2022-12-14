defmodule Y2022.Day5 do
  def part1(lines) do
    [crates, moves] = String.split(lines, "\n\n")

    parse_moves(moves)
    |> Enum.flat_map(fn {qty, from, to} -> for _ <- 0..(qty - 1), do: {1, from, to} end)
    |> move(parse_stacks(crates))
  end

  def part2(lines) do
    String.split(lines, "\n\n")
    |> then(fn [c, m] -> parse_moves(m) |> move(parse_stacks(c)) end)
  end

  defp move(moves, stacks) do
    moves
    |> Enum.reduce(stacks, fn {qty, from, to}, acc ->
      {head, tail} = Enum.split(Enum.at(acc, from), qty)

      for {stack, i} <- Enum.with_index(acc) do
        cond do
          i == to -> head ++ stack
          i == from -> tail
          true -> stack
        end
      end
    end)
    |> Enum.map(&hd/1)
    |> Enum.join()
  end

  defp parse_moves(moves) do
    moves
    |> String.split("\n")
    |> Enum.map(fn m ->
      %{"move" => move, "from" => from, "to" => to} =
        Regex.named_captures(~r/move (?<move>\d+) from (?<from>\d+) to (?<to>\d+)/, m)

      from = String.to_integer(from) - 1
      to = String.to_integer(to) - 1
      qty = String.to_integer(move)
      {qty, from, to}
    end)
  end

  defp parse_stacks(crates) do
    cols = String.to_integer(String.last(crates))
    crates = String.split(crates, "\n")
    crates = Enum.slice(crates, 0, length(crates) - 1)
    rows = length(crates)

    Enum.into(0..(cols - 1), [], fn col ->
      for row <- 0..rows, line = Enum.at(crates, row) do
        String.slice(line, col * 4, 4)
        |> String.replace(["[", "]"], "")
        |> String.trim()
      end
      |> Enum.reject(&(&1 == ""))
    end)
  end
end
