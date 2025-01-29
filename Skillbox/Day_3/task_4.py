from functools import wraps
from time import sleep
from typing import Callable, Any

def slowdown_2s(func: Callable) -> Callable:
    """
    Декоратор, который замедляет выполнение декорируемой функции на 2 секунды.
    Args:
        func: Функция, которую нужно декорировать.

    Returns: Ссылка на функцию-обёртку.
    """
    @wraps(func)
    def wrapper(*args, **kwargs) -> Any:
        """
        Обёртка для декорируемой функции, осуществляющая задержку перед вызовом.
        Задерживает выполнение на 2 секунды, а затем вызывает оригинальную функцию
        с переданными ей аргументами.
        Args:
            *args: Позиционные аргументы, передаваемые в декорируемую функцию.
            **kwargs: Именованные аргументы, передаваемые в декорируемую функцию.
        Returns: Результат, возвращаемый декорируемой функцией.
        """
        sleep(2)
        result = func(*args, **kwargs)
        return result
    return wrapper


@slowdown_2s
def summer():
    print("Скоро лето!!! ")


summer()