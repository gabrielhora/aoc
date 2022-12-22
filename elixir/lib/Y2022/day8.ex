defmodule Y2022.Day8 do
  def part1(lines) do
    grid = parse(lines)
    size = lines |> String.split() |> length()

    for x <- 0..(size - 1), y <- 0..(size - 1) do
      is_visible(grid, size, x, y)
    end
    |> Enum.count(& &1)
  end

  def part2(lines) do
    grid = parse(lines)
    size = lines |> String.split() |> length()

    for x <- 0..(size - 1), y <- 0..(size - 1) do
      scenic_score(grid, size, x, y)
    end
    |> Enum.max()
  end

  defp parse(lines) do
    lines
    |> String.split("\n")
    |> Enum.map(fn l -> String.graphemes(l) |> Enum.map(&String.to_integer/1) end)
  end

  defp split_at(enum, idx) do
    left = Enum.slice(enum, 0..(idx - 1))
    right = Enum.slice(enum, (idx + 1)..-1)
    [left, right]
  end

  defp visible_to(enum, height) do
    Enum.find(enum, fn el -> el >= height end) == nil
  end

  defp is_visible(grid, size, x, y) do
    if x == size - 1 || x == 0 || y == size - 1 || y == 0 do
      true
    else
      row = Enum.at(grid, x)
      col = for i <- 0..(size - 1), do: Enum.at(grid, i) |> Enum.at(y)
      height = Enum.at(row, y)

      [left, right] = split_at(row, y)
      [up, down] = split_at(col, x)

      visible_to(right, height) || visible_to(left, height) ||
        visible_to(up, height) || visible_to(down, height)
    end
  end

  defp calculate_score(enum, height) do
    Enum.reduce_while(enum, 0, fn el, acc ->
      if el >= height do
        {:halt, acc + 1}
      else
        {:cont, acc + 1}
      end
    end)
  end

  defp scenic_score(grid, size, x, y) do
    if x == size - 1 || x == 0 || y == size - 1 || y == 0 do
      0
    else
      row = Enum.at(grid, x)
      col = for i <- 0..(size - 1), do: Enum.at(grid, i) |> Enum.at(y)
      height = Enum.at(row, y)

      [left, right] = split_at(row, y)
      [up, down] = split_at(col, x)

      left_score = calculate_score(Enum.reverse(left), height)
      right_score = calculate_score(right, height)
      up_score = calculate_score(Enum.reverse(up), height)
      down_score = calculate_score(down, height)

      left_score * right_score * up_score * down_score
    end
  end
end
