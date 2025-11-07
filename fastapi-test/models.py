from database import Base
from sqlalchemy.orm import Mapped, mapped_column


class TestModel(Base):

    __tablename__ = "test_model"
    id: Mapped[int] = mapped_column(primary_key=True)
    text1: Mapped[str] = mapped_column()
    text2: Mapped[str] = mapped_column()
    text3: Mapped[str] = mapped_column()