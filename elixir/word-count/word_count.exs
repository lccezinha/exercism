defmodule Words do
  @doc """
  Count the number of words in the sentence.

  Words are compared case-insensitively.
  """
  @spec count(String.t()) :: map
  def count(sentence) do
    sentence
      |> String.downcase
      |> String.split(~r/[^[:alnum:]-]/u, trim: true) # Return only alphanumeric chars, more info: https://www.petefreitag.com/cheatsheets/regex/character-classes/
      |> count_words
  end

  defp count_words(words) do
    Enum.reduce(words, %{}, fn(word, map) -> update_count(word, map) end)
  end

  defp update_count(word, map), do: Map.update(map, word, 1, &(&1 + 1)) # fn(n) -> n + 1 end
end
