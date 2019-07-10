defmodule QiitaApiTest do
  use ExUnit.Case
  doctest QiitaApi

  test "greets the world" do
    assert QiitaApi.hello() == :world
  end
end
