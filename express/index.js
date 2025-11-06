import express from 'express';

const app = express();

// 포트 설정 (기본: 3000)
const PORT = process.env.PORT || 3000;

// 기본 라우트
app.get('/', (req, res) => {
  res.json({ hello: 'world3' });
});

// 서버 실행
app.listen(PORT, () => {
  console.log(`🚀 Express server running on http://localhost:${PORT}`);
});