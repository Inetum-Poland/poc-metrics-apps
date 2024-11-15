from db.mongodb import get_database, MongoDB


async def get_db() -> MongoDB:
    """
    Get instance of MongoDB class.
    Returns:
      MongoDB: Instance of MongoDB class
    """
    db = await get_database()
    return db
