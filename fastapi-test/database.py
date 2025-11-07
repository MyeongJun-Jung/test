import os
from sqlalchemy import MetaData
from sqlalchemy.orm import DeclarativeBase
from sqlalchemy.ext.asyncio import AsyncSession, create_async_engine, async_sessionmaker, AsyncAttrs

DB_URL = os.getenv("DB_URL")

meta = MetaData()
# ✅ 1. Async Engine 생성
DATABASE_URL = f"postgresql+asyncpg://{DB_URL}"

engine = create_async_engine(
    DATABASE_URL,
    echo=True,  # SQL 출력 (디버그용)
)

# ✅ 2. Session maker
AsyncSessionLocal = async_sessionmaker(
    bind=engine,
    class_=AsyncSession,
    expire_on_commit=False,
)

# ✅ 3. Declarative Base
class Base(AsyncAttrs, DeclarativeBase):
    pass

# ✅ 4. Dependency (FastAPI 등에서 사용)
async def get_db():
    async with AsyncSessionLocal() as session:
        yield session