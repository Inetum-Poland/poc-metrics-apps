import asyncio
from random import randint
from typing import Callable
from core.tracing import get_tracer


tracer = get_tracer(__name__)


async def perform_operation(
        operation: str,
        a: int,
        b: int,
        function: Callable[[int, int], int]
        ) -> int:
    """
    Run function with provided parametrs with added delay.
    Delay is a value between 20ms to 800ms.

    Parameters:
      operation: Name of the operation
      a: First number
      b: Second number
      function: Function to call
    Returns:
      int: Result of provided function
    """
    with tracer.start_as_current_span(operation) as span:
        span.add_event(f"{operation} started")

        delay = randint(20, 800) / 1000
        await asyncio.sleep(delay)
        result = function(a, b)

        span.set_status(200, "ok")
        span.add_event(f"{operation} done")

        return result


async def add(a: int, b: int) -> int:
    """
    Add two numbers.

    Parameters:
      a: First addend
      b: Second addend
    Returns:
      int: Sum
    """
    return await perform_operation("Addition", a, b, lambda x, y: x + y)


async def subtract(a: int, b: int) -> int:
    """
    Subtract two numbers.

    Parameters:
      a: Minuend
      b: Subtrahend
    Returns:
      int: Difference
    """
    return await perform_operation("Substraction", a, b, lambda x, y: x - y)


async def multiply(a: int, b: int) -> int:
    """
    Multiply two numbers.

    Parameters:
      a: First factor
      b: Second factor
    Returns:
      int: Product
    """
    return await perform_operation("Multiplication", a, b, lambda x, y: x * y)


async def divide(a: int, b: int) -> int:
    """
    Divide two numbers.

    Parameters:
      a: Dividend
      b: Divisor
    Returns:
      int: Quotient
    """
    return await perform_operation("Division", a, b, lambda x, y: x / y)
