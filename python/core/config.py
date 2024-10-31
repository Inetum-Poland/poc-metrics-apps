from pydantic_settings import BaseSettings


class Settings(BaseSettings):
    PROJECT_NAME: str = "AppPython"
    MONGO_URL: str = "mongodb://root:Password123@mongodb:27017"
    MONGO_DB_NAME: str = "Data"
    OTLP_ENDPOINT: str = "http://opentelemetry:4318/v1/traces"

    class Config:
        case_sensitive = True

settings = Settings()
