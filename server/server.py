from fastapi import FastAPI

router = FastAPI()


# Добавление пароля (POST)
@router.post("/passwords/")
async def add_password(url: str, password: str):
    return {"password": password}


# Удаление пароля (DELETE)
@router.delete("/passwords/{password_id}")
async def delete_password(password_id: int):
    return {"message": f"Password {password_id} deleted"}


# Получение паролей (GET)
@router.get("/passwords/")
async def get_passwords():
    return {"passwords": ["password1", "password2"]}
