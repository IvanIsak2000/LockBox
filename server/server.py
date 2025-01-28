import argparse

import uvicorn
from fastapi import FastAPI

router = FastAPI()


# Добавление пароля (POST)
@router.post("/passwords/")
async def add_password(url: str, password: str):
    return {"url": url, "password": password}


# Удаление пароля (DELETE)
@router.delete("/passwords/{password_id}")
async def delete_password(password_id: int):
    return {"message": f"Password {password_id} deleted"}


# Получение паролей (GET)
@router.get("/passwords/")
async def get_passwords():
    return {"passwords": ["password1", "password2"]}


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    subparsers = parser.add_subparsers(dest="command", help="Доступные команды")

    subparsers.add_parser("init", help='Первоначальная инициализация сервера')
    subparsers.add_parser("serve", help='Запуск сервера.')

    args = parser.parse_args()
    if args.command == "serve":
        uvicorn.run(router)
