from typing import Any
from fastapi import FastAPI, WebSocket, WebSocketDisconnect
import json
import asyncio
import uvicorn

app = FastAPI()


@app.get("/")
def read_root():
    return {"message": "FastAPI WebSocket Benchmark Server"}


@app.websocket("/ws")
async def websocket_endpoint(ws: WebSocket):
    await ws.accept()
    print("✅ Client connected")

    try:
        while True:
            msg = await ws.receive_text()
            try:
                data = json.loads(msg)
            except json.JSONDecodeError:
                await ws.send_text(json.dumps({"error": "invalid JSON"}))
                continue

            # Phoenix와 유사한 포맷: ["join_ref","ref","topic","event","payload"]
            if isinstance(data, list) and len(data) >= 4:
                event = data[3]
                topic = data[2] if len(data) > 2 else None

                if event == "phx_join":
                    await ws.send_text(
                        json.dumps(["1", "1", topic, "phx_reply", {"status": "ok"}])
                    )
                    print(f"📡 JOIN: {topic}")
                elif event == "ping":
                    await ws.send_text(
                        json.dumps(["2", "2", topic, "pong", {"msg": "hello"}])
                    )
                    # 서버 응답이 너무 빨라 부하가 약하면 약간 delay를 줄 수도 있음
                    await asyncio.sleep(0.001)
                else:
                    await ws.send_text(json.dumps(["?", "?", topic, "unknown_event"]))
            else:
                await ws.send_text(json.dumps({"error": "invalid format"}))

    except WebSocketDisconnect:
        print("❌ Client disconnected")
    except Exception as e:
        print("⚠️ Error:", e)
        await ws.close()


if __name__ == "__main__":
    # 로컬 테스트용 서버 실행 (ex: ws://localhost:8000/ws)
    uvicorn.run(app, host="0.0.0.0", port=8000)
