from fastapi import APIRouter, HTTPException, Depends
from motor.motor_asyncio import AsyncIOMotorDatabase
from .deps import get_db
from core.telemetry import get_tracer
from services.main_service import long_run, short_run, database_run, failed_run


router = APIRouter(prefix="/api")
tracer = get_tracer(__name__)


@router.get("/long_run")
async def get_long_run():
    with tracer.start_as_current_span("LongRun") as span:
        span.add_event("LongRun started")
        try:
            result = await long_run()
            span.set_status(200, "ok")
            span.add_event("LongRun done")
            return result
        except Exception as e:
            raise HTTPException(status_code=500, detail=str(e))
            


@router.get("/short_run")
async def get_short_run():
    with tracer.start_as_current_span("ShortRun") as span:
        span.add_event("ShortRun started")
        try:
            result = await short_run()
            span.set_status(200, "ok")
            span.add_event("ShortRun done")
            return result
        except Exception as e:
            raise HTTPException(status_code=500, detail=str(e))


@router.get("/database_run")
async def get_database_run(db: AsyncIOMotorDatabase = Depends(get_db)):
    with tracer.start_as_current_span("DatabaseRun") as span:
        span.add_event("DatabaseRun started")
        try:
            result = await database_run(db)
            span.set_status(200, "ok")
            span.add_event("DatabaseRun done")
            return result
        except Exception as e:
            raise HTTPException(status_code=500, detail=str(e))


@router.get("/failed_run")
async def get_failed_run():
    with tracer.start_as_current_span("FailedRun") as span:
        span.add_event("FailedRun started")
        try:
            result = await failed_run()
            span.set_status(200, "ok")
            span.add_event("FailedRun done")
            return result
        except Exception as e:
            span.set_status(500, "Unexpected error")
            span.record_exception(e)
            raise HTTPException(status_code=500, detail=str(e))
