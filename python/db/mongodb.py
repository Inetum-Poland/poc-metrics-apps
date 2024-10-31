from motor.motor_asyncio import AsyncIOMotorClient
from core.config import settings


class MongoDB:
    """
    Class that represents MongoDB and allows connection.
    """
    client: AsyncIOMotorClient = None
    db = None


async def get_database() -> MongoDB:
    """
    Return instance of MongoDB class.
    Returns:
      MongoDB: instance of MongoDB class.
    """
    return db.db


async def connect_to_mongo() -> None:
    """
    Establish connection to MongoDB.
    Done by assigning values to instance of MongoDB class. 
    """
    db.client = AsyncIOMotorClient(settings.MONGO_URL)
    db.db = db.client[settings.MONGO_DB_NAME]


async def close_mongo_connection() -> None:
    """
    Close connection to MongoDB.
    """
    if db.client:
        db.client.close()


db = MongoDB()
