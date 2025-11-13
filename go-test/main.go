package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Gorilla WebSocket 업그레이더 설정
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 벤치마크용이니까 origin 체크는 그냥 통과시켜도 됨
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	r := gin.Default()

	// ✅ static 파일 제공 (/static 폴더 안의 파일)
	r.Static("/static", "./static")

	// ✅ 루트("/") 요청 시 index.html 반환
	r.GET("/home", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	// ✅ WebSocket 엔드포인트 (FastAPI의 /ws 와 동일)
	r.GET("/ws", websocketHandler)

	log.Println("🚀 Gin WebSocket Benchmark Server started on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// /ws 핸들러
func websocketHandler(c *gin.Context) {
	// HTTP → WebSocket 업그레이드
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("❌ WebSocket upgrade error:", err)
		return
	}
	log.Println("✅ Client connected")

	defer func() {
		conn.Close()
		log.Println("❌ Client disconnected")
	}()

	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("⚠️ Read error:", err)
			return
		}
		if mt != websocket.TextMessage {
			// 텍스트만 처리
			continue
		}

		// FastAPI 버전처럼 JSON 파싱
		var data []interface{}
		if err := json.Unmarshal(msg, &data); err != nil {
			// invalid JSON 응답
			resp, _ := json.Marshal(map[string]string{"error": "invalid JSON"})
			_ = conn.WriteMessage(websocket.TextMessage, resp)
			continue
		}

		if len(data) < 4 {
			resp, _ := json.Marshal(map[string]string{"error": "invalid format"})
			_ = conn.WriteMessage(websocket.TextMessage, resp)
			continue
		}

		// ["join_ref","ref","topic","event","payload"]
		// 인덱스 2: topic, 3: event
		topic, _ := data[2].(string)
		event, _ := data[3].(string)

		switch event {
		case "phx_join":
			// ["1","1",topic,"phx_reply",{"status":"ok"}]
			reply := []interface{}{"1", "1", topic, "phx_reply", map[string]string{"status": "ok"}}
			if err := writeJSON(conn, reply); err != nil {
				log.Println("⚠️ Write error:", err)
				return
			}
			log.Println("📡 JOIN:", topic)

		case "ping":
			// ["2","2",topic,"pong",{"msg":"hello"}]
			reply := []interface{}{"2", "2", topic, "pong", map[string]string{"msg": "hello"}}
			if err := writeJSON(conn, reply); err != nil {
				log.Println("⚠️ Write error:", err)
				return
			}
			// FastAPI 버전처럼 살짝 딜레이 줄 수도 있음 (부하 조절)
			time.Sleep(1 * time.Millisecond)

		default:
			reply := []interface{}{"?", "?", topic, "unknown_event"}
			if err := writeJSON(conn, reply); err != nil {
				log.Println("⚠️ Write error:", err)
				return
			}
		}
	}
}

// JSON 배열/객체를 WebSocket 텍스트로 보내는 helper
func writeJSON(conn *websocket.Conn, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("json marshal error: %w", err)
	}
	return conn.WriteMessage(websocket.TextMessage, b)
}
