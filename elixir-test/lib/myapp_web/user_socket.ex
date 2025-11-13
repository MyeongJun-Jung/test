# lib/myapp_web/user_socket.ex
defmodule MyappWeb.UserSocket do
  use Phoenix.Socket

  # ✅ benchmark:* topic을 BenchmarkChannel로 연결
  channel "benchmark:*", MyappWeb.BenchmarkChannel

  # 연결 허용
  def connect(_params, socket, _connect_info) do
    {:ok, socket}
  end

  # socket id는 필요 없으면 nil
  def id(_socket), do: nil
end
