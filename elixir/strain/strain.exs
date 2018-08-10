defmodule Strain do
  @doc """
  Given a `list` of items and a function `fun`, return the list of items where
  `fun` returns true.

  Do not use `Enum.filter`.
  """
  @spec keep(list :: list(any), fun :: (any -> boolean)) :: list(any)
  def keep([head | tail], fun) do
    list = [(if fun.(head), do: head) | keep(tail, fun)]
    Enum.reject(list, fn(item) -> is_nil(item) end)
  end

  def keep([], fun) do
    []
  end

  @doc """
  Given a `list` of items and a function `fun`, return the list of items where
  `fun` returns false.

  Do not use `Enum.reject`.
  """
  @spec discard(list :: list(any), fun :: (any -> boolean)) :: list(any)
  def discard([head | tail], fun) do
    list = [(unless fun.(head), do: head) | discard(tail, fun)]
    Enum.filter(list, fn(item) -> !is_nil(item) end)
  end

  def discard([], fun) do
    []
  end
end
