package com.example.demo.hello;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.web.socket.*;
import org.springframework.web.socket.handler.TextWebSocketHandler;

public class WsHandler extends TextWebSocketHandler {

    private final ObjectMapper mapper = new ObjectMapper();

    @Override
    public void afterConnectionEstablished(WebSocketSession session) {
        System.out.println("✅ Client connected");
    }

    @Override
    protected void handleTextMessage(WebSocketSession session, TextMessage message) {
        try {
            String payload = message.getPayload();
            Object[] arr = mapper.readValue(payload, Object[].class);

            if (arr.length < 4) {
                session.sendMessage(new TextMessage(
                        mapper.writeValueAsString(new ErrorMsg("invalid format"))
                ));
                return;
            }

            String topic = arr[2].toString();
            String event = arr[3].toString();

            switch (event) {
                case "phx_join":
                    Object[] joinReply = {"1", "1", topic, "phx_reply",
                            mapper.readValue("{\"status\":\"ok\"}", Object.class)};
                    session.sendMessage(new TextMessage(mapper.writeValueAsString(joinReply)));
                    break;

                case "ping":
                    Object[] pongReply = {"2", "2", topic, "pong",
                            mapper.readValue("{\"msg\":\"hello\"}", Object.class)};
                    session.sendMessage(new TextMessage(mapper.writeValueAsString(pongReply)));
                    break;

                default:
                    Object[] unknown = {"?", "?", topic, "unknown_event"};
                    session.sendMessage(new TextMessage(mapper.writeValueAsString(unknown)));
                    break;
            }

        } catch (Exception e) {
            System.out.println("⚠️ Error: " + e.getMessage());
        }
    }

    @Override
    public void afterConnectionClosed(WebSocketSession session, CloseStatus status) {
        System.out.println("❌ Client disconnected");
    }

    static class ErrorMsg {
        public String error;

        public ErrorMsg(String error) {
            this.error = error;
        }
    }
}
