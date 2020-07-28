# Паттерн facade
Синтетический пример заказа товара покупателем.

Покупатель делает заказ, происходит набор действий:
- проверка аккаунта
- проверка наличия
- резервирование
- возврат

Фасад дает возможность реализовать операции бронирования и возврата товара без знания реализации этих методов.

Для наглядности выводим результаты работы внутренних методов для набора разных операций.