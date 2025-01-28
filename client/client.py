import argparse
import asyncio
import aiohttp
from colorama import Fore, Style


URL = "http://127.0.0.1:8000/"


async def add_password(url: str, password: str):
    async with aiohttp.ClientSession() as session:
        response = await session.post(
            url=URL+'passwords/',
            params={
                "url": url,
                "password": password
            }
        )
        print(f"{response.json=}")


if __name__ == '__main__':
    parser = argparse.ArgumentParser(
        prog="Lockbox",
        description="Простой self-hosted менеджер паролей",
        formatter_class=argparse.RawTextHelpFormatter  # Для красивого форматирования
    )

    # Создаем подпарсеры для каждой команды
    subparsers = parser.add_subparsers(dest="command", help="Доступные команды")

    # Команда connect
    connect_parser = subparsers.add_parser(
        "connect",
        help="Соединиться с сервером",
        description="Соединиться с сервером. Аргументы: server_url: str",
        formatter_class=argparse.RawTextHelpFormatter
    )
    connect_parser.add_argument("server_url", type=str, help="URL сервера")

    # Команда sync
    sync_parser = subparsers.add_parser(
        "sync",
        help="Синхронизировать локальные пароли с сервером",
        description="Синхронизировать локальные пароли с сервером."
    )

    # Команда add
    add_parser = subparsers.add_parser(
        "add",
        help="Добавление нового пароля",
        description="Добавление нового пароля. Аргументы: url: str, password: str"
    )
    add_parser.add_argument("url", type=str, help="URL сайта")
    add_parser.add_argument("password", type=str, help="Пароль")

    # Парсинг аргументов
    args = parser.parse_args()

    # Обработка команд
    if args.command == "connect":
        print(f"Соединяемся с сервером: {args.server_url}")
    elif args.command == "sync":
        print(Fore.BLUE + "Синхронизация паролей...")
    elif args.command == "add":
        print(f"Добавлен новый пароль для {args.url}: {args.password}")
    else:
        parser.print_help()
    # asyncio.run(add_password(URL, "123"))

    print(Style.RESET_ALL)
