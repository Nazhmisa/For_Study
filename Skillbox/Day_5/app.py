import openai
import os
from dotenv import load_dotenv

load_dotenv()

openai.api_key = os.getenv("OPENAI_API_KEY")
print("API Key:", openai.api_key)

try:
    response = openai.Model.list()
    print([model["id"] for model in response["data"]])
except Exception as e:
    print("Ошибка:", e)
