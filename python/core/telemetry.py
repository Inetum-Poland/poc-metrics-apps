from opentelemetry import trace
from opentelemetry.trace import SpanKind
from opentelemetry.exporter.otlp.proto.http.trace_exporter import OTLPSpanExporter
from opentelemetry.sdk.trace import TracerProvider, ReadableSpan
from opentelemetry.sdk.trace.export import BatchSpanProcessor
from opentelemetry.sdk.resources import Resource
from .config import settings


class CustomBatchSpanProcessor(BatchSpanProcessor):
    """Custom Class to alter behave of BatchSpanProcessor"""

    def on_end(self, span: ReadableSpan) -> None:
        """
        Overwrite behavior of BatchSpanProcessor.on_end method.
        The purpose of this is to stop sending meaningless spans of types:
          - http.request
          - http.response.start
          - http.response.body
        If span's type is different then proceed with regular on_end execution.
        Parameters:
            span: Span that should be sent
        """
        if span.kind == SpanKind.INTERNAL and (
            span.attributes.get("asgi.event.type") in ("http.request", "http.response.start", "http.response.body")
        ):
            return
        super().on_end(span=span)


def setup_telemetry(service_name: str="fastapi-service") -> TracerProvider:
    """
    Setup telemetry for an app.
    Parameters:
        service_name: name of the service to include within spans
    Returns:
        TracerProvider: tracer that can be used to create spans
    """
    resource = Resource.create({"service.name": service_name})
    tracer_provider = TracerProvider(resource=resource)
    otlp_exporter = OTLPSpanExporter(endpoint=settings.OTLP_ENDPOINT)
    span_processor = CustomBatchSpanProcessor(otlp_exporter)
    tracer_provider.add_span_processor(span_processor)
    trace.set_tracer_provider(tracer_provider)
    return tracer_provider


def get_tracer(name: str) -> TracerProvider:
    """
    Get TracerProvider to create spans.
    Parameters:
      name: TracerProvider's name
    Returns:
      TracerProvider: TracerProvider that can be used to create spans
    """
    return trace.get_tracer(name)