from contextlib import asynccontextmanager
from fastapi import FastAPI
from opentelemetry.instrumentation.fastapi import FastAPIInstrumentor

from api.endpoints import router
from core.config import settings
from core.tracing import setup_tracing
from core.logging import setup_logging
from db.mongodb import connect_to_mongo, close_mongo_connection


@asynccontextmanager
async def lifespan(app: FastAPI) -> None:
    await connect_to_mongo()
    yield
    await close_mongo_connection()


app = FastAPI(
    title=settings.PROJECT_NAME,
    openapi_url="/openapi.json"
)

setup_logging(settings.PROJECT_NAME)
setup_tracing(settings.PROJECT_NAME)
FastAPIInstrumentor.instrument_app(app)

app.include_router(router)
app.router.lifespan_context = lifespan
