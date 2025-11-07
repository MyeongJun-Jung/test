from contextlib import asynccontextmanager
from typing import Annotated, Union

from fastapi import FastAPI, Depends
from database import AsyncSessionLocal, get_db, engine, Base
from models import TestModel
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import select

@asynccontextmanager
async def lifespan(app: FastAPI):
    async with engine.begin() as conn:
        await conn.run_sync(Base.metadata.drop_all)
        await conn.run_sync(Base.metadata.create_all)

    async with AsyncSessionLocal() as session:
        for i in range(10):
            model = TestModel(
                text1=f"{i}",
                text2=f"{i}",
                text3=f"{i}",
            )
            session.add(model)
        await session.commit()
    yield
    # clean up items

app = FastAPI(lifespan=lifespan)

SessionDep = Annotated[AsyncSession, Depends(get_db)]

@app.get("/")
async def read_root(session: SessionDep):
    result = await session.scalars(select(TestModel).limit(5))
    data = result.all()
    return data
