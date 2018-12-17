defmodule PigLatin do
  @doc """
  Given a `phrase`, translate it a word at a time to Pig Latin.

  Words beginning with consonants should have the consonant moved to the end of
  the word, followed by "ay".

  Words beginning with vowels (aeiou) should have "ay" added to the end of the
  word.

  Some groups of letters are treated like consonants, including "ch", "qu",
  "squ", "th", "thr", and "sch".

  Some groups are treated like vowels, including "yt" and "xr".
  """

  @vowels ["a", "e", "i", "o", "u"]
  @fake_vowels ["yt", "yd", "xr", "xb"]

  @spec translate(phrase :: String.t()) :: String.t()
  def translate(phrase) do
    phrase
    |> String.split()
    |> Enum.map(&translate_to_pig_latin/1)
    |> Enum.join(" ")
  end

  defp translate_to_pig_latin(phrase) do
    if begin_with_vowels?(phrase), do: phrase <> "ay", else: build_world(phrase)
  end

  defp begin_with_vowels?(phrase) do
    String.starts_with?(phrase, @vowels ++ @fake_vowels)
  end

  defp begin_with_consonants?(phrase) do
    !begin_with_vowels?(phrase)
  end

  defp begin_with_two_consonants(phrase) do
    !begin_with_vowels?(String.at(phrase, 0)) && !begin_with_vowels?(String.at(phrase, 1))
  end

  defp build_world(phrase) do
    group = ~r/(ch)|(qu)|(thr)|(th)|(sch)|[^aeiou]qu|[^aeiou]{1,}/

    [phrase_head, phrase_tail] =
      String.split(phrase, group, parts: 2, trim: true, include_captures: true)

    phrase_tail <> phrase_head <> "ay"
  end
end
