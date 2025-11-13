import { WebSocketGateway, WebSocketServer } from '@nestjs/websockets';
import { Server, WebSocket } from 'ws';

@WebSocketGateway({
  cors: {
    origin: '*',
  },
  path: '/ws', // ws://localhost:3000/ws
})
export class WsGateway {
  @WebSocketServer()
  server: Server;

  handleConnection(client: WebSocket) {
    console.log('🔌 Client connected');

    client.on('message', (msg: string) => {
      try {
        const data = JSON.parse(msg); // Phoenix style array
        const [joinRef, ref, topic, event, payload] = data;

        if (event === 'phx_join') {
          console.log(`👥 Join request for topic: ${topic}`);
          client.send(JSON.stringify([joinRef, ref, topic, 'phx_reply', { status: 'ok' }]));
        }

        if (event === 'ping') {
          // Echo back like Phoenix
          client.send(
            JSON.stringify([
              joinRef,
              ref,
              topic,
              'pong',
              { time: Date.now(), ...payload },
            ]),
          );
        }
      } catch (e) {
        console.error('❌ Invalid message', msg);
      }
    });

    client.on('close', () => {
      console.log('❌ Client disconnected');
    });
  }
}
