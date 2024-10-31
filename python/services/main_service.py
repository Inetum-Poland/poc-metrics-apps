import asyncio
import json
from typing import Dict, Any
from .utils import add, subtract, multiply, divide
from core.telemetry import get_tracer


tracer = get_tracer(__name__)


async def long_run() -> Dict[str, Any]:
    
    await add(1, 2)
    await subtract(1, 2)
    await multiply(1, 2)
    await divide(1, 2)

    return {"data": "ok"}


async def short_run() -> Dict[str, Any]:
    _ = await asyncio.sleep(0.1)
    return {"data": "ok"}


async def database_run(db) -> Dict[str, Any]:
    with tracer.start_as_current_span("read_collection") as span:
        span.add_event("read_collection started")
        opts = {
            "data": {
                "$gte": 0
            }
        }
        span.set_attribute("opts", json.dumps(opts))
        span.set_attribute("collection.name", "Data")
        results = await db["Data"].find(opts).to_list()
        serialized_results = list(
            map(
                lambda doc: {**doc, "_id": str(doc["_id"])},
                results
            )
        )
        span.set_attribute("collection.length", len(serialized_results))
        span.add_event("read_collection done")
        return serialized_results


async def failed_run() -> Dict[str, Any]:
    _ = await asyncio.sleep(0.1)
    raise Exception
    return {"data": "Failed"}