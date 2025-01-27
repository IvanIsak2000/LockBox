import asyncio

from cryptography.fernet import Fernet


class Crypter:
    """
    Класс для криптографических функций на сервере
    """
    def __init__(self):
        session_key = Fernet.generate_key()
        print(f"{session_key=}")
        self.fernet = Fernet(session_key)

    async def encrypt(self, data: bytes) -> str:
        return self.fernet.encrypt(data)

    async def decrypt(self, token: str) -> bytes:
        return self.fernet.decrypt(token)


my_password = input("Введите пароль для хранения: ")
c = Crypter()
encrypt_token = asyncio.run(c.encrypt(my_password.encode()))
print(f"{encrypt_token=}")
decrypt_result = asyncio.run((c.decrypt(encrypt_token))).decode()
print(f"{decrypt_result=}")
