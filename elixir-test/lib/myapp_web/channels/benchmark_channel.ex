defmodule MyappWeb.BenchmarkChannel do
  use MyappWeb, :channel

  def join("benchmark:lobby", _payload, socket) do
    {:ok, socket}
  end

  def handle_in("ping", payload, socket) do
    push(socket, "pong", %{received: payload})
    {:noreply, socket}
  end
end
