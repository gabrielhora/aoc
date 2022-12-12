defmodule Y2022.Day12 do
  def part1(input) do
    state = init_state(input)
    state = move(state, MapSet.new([state.initial]))
    state.moves
  end

  def part2(input) do
    state = init_state(input)

    state.nodes
    |> Enum.filter(&(elem(&1, 1) == ?a))
    |> Enum.map(&move(state, MapSet.new([&1])))
    |> Enum.reject(&(&1 == %{}))
    |> Enum.map(& &1.moves)
    |> Enum.sort()
    |> hd()
  end

  defp move(_state, visit) when visit == %MapSet{}, do: %{}

  defp move(state, visit) do
    visit =
      Enum.reduce_while(visit, MapSet.new(), fn {{x, y}, height}, acc ->
        if {x, y} == elem(state.target, 0) do
          {:halt, MapSet.new([])}
        else
          th = state.nodes[{x, y - 1}]
          rh = state.nodes[{x + 1, y}]
          bh = state.nodes[{x, y + 1}]
          lh = state.nodes[{x - 1, y}]

          acc = if th && th - height <= 1, do: MapSet.put(acc, {{x, y - 1}, th}), else: acc
          acc = if rh && rh - height <= 1, do: MapSet.put(acc, {{x + 1, y}, rh}), else: acc
          acc = if bh && bh - height <= 1, do: MapSet.put(acc, {{x, y + 1}, bh}), else: acc
          acc = if lh && lh - height <= 1, do: MapSet.put(acc, {{x - 1, y}, lh}), else: acc

          {:cont, acc}
        end
      end)

    if visit == %MapSet{} do
      # found the target
      state
    else
      visit = visit |> MapSet.difference(state.visited)
      state = %{state | moves: state.moves + 1, visited: MapSet.union(state.visited, visit)}
      move(state, visit)
    end
  end

  defp init_state(input) do
    letters = input |> String.split("\n") |> Enum.map(&String.graphemes/1)
    rows = length(letters)
    cols = length(Enum.at(letters, 0))

    nodes =
      for y <- 0..(rows - 1), x <- 0..(cols - 1) do
        weight = letters |> Enum.at(y) |> Enum.at(x)
        {{x, y}, :binary.last(weight)}
      end
      |> Map.new()

    initial = {{cx, cy}, _} = Enum.find_value(nodes, fn {k, v} -> if v == ?S, do: {k, ?a} end)
    target = {{tx, ty}, _} = Enum.find_value(nodes, fn {k, v} -> if v == ?E, do: {k, ?z} end)

    %{
      nodes: nodes |> Map.replace({cx, cy}, ?a) |> Map.replace({tx, ty}, ?z),
      initial: initial,
      target: target,
      moves: 0,
      visited: MapSet.new()
    }
  end
end
