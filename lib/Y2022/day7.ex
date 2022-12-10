defmodule Y2022.Day7 do
  def part1(lines) do
    Enum.reduce(lines, %{}, &parse(&1, &2))
    |> then(&Map.delete(&1, :current_dir))
    |> Map.reject(fn {_, size} -> size > 100_000 end)
    |> Map.values()
    |> Enum.sum()
  end

  def part2(lines) do
    state =
      Enum.reduce(lines, %{}, &parse(&1, &2))
      |> then(&Map.delete(&1, :current_dir))

    unused = 70_000_000 - state["/"]

    state
    |> Enum.sort_by(fn {_, size} -> size end)
    |> Enum.find_value(fn {_, size} ->
      if unused + size >= 30_000_000, do: size
    end)
  end

  defp parse("$ cd ..", state) do
    {_, new_state} =
      Map.get_and_update(state, :current_dir, fn cur ->
        {cur, Path.dirname(cur)}
      end)

    new_state
  end

  defp parse("$ cd " <> dir, state) do
    {_, new_state} =
      Map.get_and_update(state, :current_dir, fn cur ->
        {cur, Path.join(cur || "/", dir)}
      end)

    new_state
  end

  defp parse("dir " <> _, state), do: state
  defp parse("$ ls", state), do: state

  defp parse(file_line, state) do
    size =
      file_line
      |> String.split(" ")
      |> then(fn [size, _] -> String.to_integer(size) end)

    add_size(state, size, state[:current_dir])
  end

  defp add_size(state, size, current_dir) do
    # add size to current dir and all it's parents
    {_, state} =
      Map.get_and_update(state, current_dir, fn cur ->
        {cur, (cur || 0) + size}
      end)

    if current_dir == "/" do
      state
    else
      add_size(state, size, Path.dirname(current_dir))
    end
  end
end
