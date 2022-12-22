defmodule Y2022.Day9 do
  def part1(lines) do
    parse_moves(lines) |> move_rope(2)
  end

  def part2(lines) do
    parse_moves(lines) |> move_rope(10)
  end

  defp move_rope(moves, num_knots) do
    moves
    |> Enum.reduce({List.duplicate({0, 0}, num_knots), MapSet.new()}, fn m, {knots, visited} ->
      knots = move(m, knots)
      {knots, MapSet.put(visited, List.last(knots))}
    end)
    |> elem(1)
    |> MapSet.size()
  end

  defp parse_moves(lines) do
    lines
    |> Enum.map(&String.split(&1, " "))
    |> Enum.map(fn [dir, count] -> {dir, String.to_integer(count)} end)
    |> Enum.flat_map(fn {dir, count} -> for _ <- 0..(count - 1), do: dir end)
  end

  defp move_knots([head]), do: [head]

  defp move_knots([{hx, hy}, {tx, ty} | rest]) do
    delta_x = hx - tx
    delta_y = hy - ty

    tail =
      case {delta_x, delta_y} do
        {-2, -2} -> {tx - 1, ty - 1}
        {-2, -1} -> {tx - 1, ty - 1}
        {-2, 0} -> {tx - 1, ty}
        {-2, 1} -> {tx - 1, ty + 1}
        {-2, 2} -> {tx - 1, ty + 1}
        {-1, -2} -> {tx - 1, ty - 1}
        {-1, 2} -> {tx - 1, ty + 1}
        {0, -2} -> {tx, ty - 1}
        {0, 2} -> {tx, ty + 1}
        {1, -2} -> {tx + 1, ty - 1}
        {1, 2} -> {tx + 1, ty + 1}
        {2, -2} -> {tx + 1, ty - 1}
        {2, -1} -> {tx + 1, ty - 1}
        {2, 0} -> {tx + 1, ty}
        {2, 1} -> {tx + 1, ty + 1}
        {2, 2} -> {tx + 1, ty + 1}
        _ -> {tx, ty}
      end

    [{hx, hy} | move_knots([tail | rest])]
  end

  # move applies the move logic and return the updated knots positions

  defp move("R", [{x, y} | knots]) do
    head = {x + 1, y}
    move_knots([head | knots])
  end

  defp move("L", [{x, y} | knots]) do
    head = {x - 1, y}
    move_knots([head | knots])
  end

  defp move("U", [{x, y} | knots]) do
    head = {x, y + 1}
    move_knots([head | knots])
  end

  defp move("D", [{x, y} | knots]) do
    head = {x, y - 1}
    move_knots([head | knots])
  end
end
