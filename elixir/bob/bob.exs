defmodule Bob do
  def hey(input) do
    input =  String.trim(input)

    cond do
      shout?(input) && ask?(input) -> "Calm down, I know what I'm doing!"
      ask?(input) -> "Sure."
      silence?(input) -> "Fine. Be that way!"
      shout?(input) -> "Whoa, chill out!"
      true -> "Whatever."
    end
  end

  defp shout?(input) do
    String.upcase(input) == input && String.downcase(input) != input
  end

  defp ask?(input) do
    String.ends_with?(input, "?")
  end

  defp silence?(input) do
    input == ""
  end
end
