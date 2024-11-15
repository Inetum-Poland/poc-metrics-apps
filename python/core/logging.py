import logging
from opentelemetry.sdk.resources import Resource
from opentelemetry.sdk._logs import LoggerProvider, LoggingHandler
from opentelemetry.exporter.otlp.proto.grpc._log_exporter import (
    OTLPLogExporter
)
from opentelemetry._logs import set_logger_provider
from opentelemetry.sdk._logs.export import BatchLogRecordProcessor
from .config import settings


# https://github.com/open-telemetry/opentelemetry-python/issues/3389
class PyMongoLoggingFilter(logging.Filter):
    """Filter for pymongo log records"""
    def filter(self, record) -> bool:
        """
        Filter out records from pymongo.
        Parameters:
          record: Record to check
        Returns:
          bool: True if "pymongo" is not part of record name. Otherwise False.
        """
        return "pymongo" not in record.name


def setup_logging(service_name: str = "fastapi-service") -> None:
    """
    Setup OTEL Logging.
    Parameters:
      service_name: name of the service
    """
    resource = Resource.create({
        "service.name": service_name,
        "deployment.environment": "local"
    }, schema_url="https://opentelemetry.io/schemas/1.21.0")

    logger_provider = LoggerProvider(resource=resource)
    set_logger_provider(logger_provider)
    otlp_exporter = OTLPLogExporter(
        endpoint=settings.OTLP_ENDPOINT,
        insecure=True
    )
    logger_provider.add_log_record_processor(
        BatchLogRecordProcessor(otlp_exporter)
    )
    handler = LoggingHandler(
        level=logging.NOTSET,
        logger_provider=logger_provider
    )
    handler.addFilter(PyMongoLoggingFilter())
    logging.getLogger().setLevel(logging.NOTSET)
    logging.getLogger().addHandler(handler)
