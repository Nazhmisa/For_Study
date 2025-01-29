from functools import wraps
from typing import Callable

def counter(func: Callable) -> Callable:
    """
    Декоратор, который считает количество вызовов декорируемой функции.
    """
    @wraps(func)
    def wrapper(*args, **kwargs):
        # Увеличиваем счётчик вызовов
        wrapper.count += 1
        print(f"Функция '{func.__name__}' была вызвана {wrapper.count} раз(а).")
        return func(*args, **kwargs)
    
    # Инициализируем атрибут счётчика
    wrapper.count = 0
    return wrapper

# Пример использования
@counter
def greet(name: str) -> None:
    print(f"Привет, {name}!")

# Проверка работы
greet("Алиса")
greet("Боб")
greet("Чарли")
