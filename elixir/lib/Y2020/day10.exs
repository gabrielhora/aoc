#!/usr/bin/env elixir

defmodule Y2020.Day10 do
  @input File.read!("#{__DIR__}/../../priv/Y2020/day10/example.txt")

  def part1 do
    result =
      @input
      |> parse
      |> Enum.chunk_every(2, 1, :discard)
      |> Enum.reduce([], fn [a, b], acc -> [b - a | acc] end)
      |> Enum.reverse()
      |> Enum.frequencies()
      |> then(fn %{1 => ones, 3 => threes} -> ones * threes end)

    IO.puts("part 1: #{result}")
  end

  def part2 do
    @input
    |> parse
    |> Enum.chunk_every(2, 1, :discard)
    |> Enum.reduce([], fn [a, b], acc -> [b - a | acc] end)
    |> Enum.reverse()
    |> IO.inspect()
  end

  defp parse(input) do
    adapters =
      input
      |> String.split("\n")
      |> Enum.map(&String.to_integer/1)
      |> Enum.sort()

    adapters = List.insert_at(adapters, 0, 0)
    adapters = List.insert_at(adapters, length(adapters), List.last(adapters) + 3)
    adapters
  end
end

Y2020.Day10.part2()
