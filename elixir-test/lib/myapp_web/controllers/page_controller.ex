defmodule MyappWeb.PageController do
  use MyappWeb, :controller

  def hello(conn, _params) do
    json(conn, %{hello: "world"})
  end

  def home(conn, _params) do
    render(conn, :home)
  end
end
