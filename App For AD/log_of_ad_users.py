import os
import ldap3
import slack
from dotenv import load_dotenv
from datetime import datetime
import logging

# Загружаем конфигурационные параметры из файла .env
load_dotenv()

# Конфигурация
AD_SERVER = os.getenv('AD_SERVER')
AD_USER = os.getenv('AD_USER')
AD_PASSWORD = os.getenv('AD_PASSWORD')
AD_SEARCH_BASE = os.getenv('AD_SEARCH_BASE')
LOG_FILE = 'disabled_users.log'

SLACK_TOKEN = os.getenv('SLACK_TOKEN')

# Настроим логирование
logging.basicConfig(
    filename='app.log',
    level=logging.INFO,
    format='%(asctime)s - %(levelname)s - %(message)s',
)

# Инициализация клиента Slack
client = slack.WebClient(token=SLACK_TOKEN)

# Функция для поиска отключенных пользователей в AD
def get_disabled_users():
    logging.info("[*] Connecting to AD...")
    server = ldap3.Server(AD_SERVER)
    connection = ldap3.Connection(server, user=AD_USER, password=AD_PASSWORD)
    
    try:
        connection.bind()
        logging.info("[*] Connected successfully!")
    except ldap3.LDAPBindError as e:
        logging.error(f"[ERROR] Failed to connect: {e}")
        return []

    # Фильтр для поиска отключенных пользователей
    search_filter = '(&(objectClass=user)(userAccountControl:1.2.840.113556.1.4.803:=2))'
    attributes = ['sAMAccountName', 'userPrincipalName']

    logging.info("[*] Searching for disabled users in AD...")
    connection.search(search_base=AD_SEARCH_BASE, search_filter=search_filter, attributes=attributes)
    disabled_users = []

    for entry in connection.entries:
        username = entry.sAMAccountName.value
        email = entry.userPrincipalName.value
        disabled_users.append({'username': username, 'email': email})

    connection.unbind()
    return disabled_users

# Функция для логирования отключенных пользователей
def log_disabled_users(users):
    logging.info("[*] Logging disabled users...")
    with open(LOG_FILE, 'a') as log_file:
        timestamp = datetime.now().strftime('%Y-%m-%d %H:%M:%S')
        log_file.write(f'[{timestamp}] Found {len(users)} disabled users:\n')
        for user in users:
            log_file.write(f"Username: {user['username']}, Email: {user['email']}\n")
        log_file.write('\n')

# Функция для получения всех пользователей Slack
def get_slack_users():
    logging.info("[*] Fetching users from Slack...")
    try:
        response = client.users_list()
        if response["ok"]:
            return response["members"]
        else:
            logging.error("[ERROR] Failed to fetch users from Slack")
            return []
    except slack.errors.SlackApiError as e:
        logging.error(f"[ERROR] Failed to fetch users from Slack: {e.response['error']}")
        return []

# Основной скрипт
def main():
    disabled_users = get_disabled_users()
    
    if disabled_users:
        log_disabled_users(disabled_users)
        logging.info(f'[INFO] Logged {len(disabled_users)} disabled users to {LOG_FILE}')
        
        slack_users = get_slack_users()
        if slack_users:
            logging.info("[*] Checking which users can be disabled in Slack...")
            for user in disabled_users:
                # Проверка, если этот пользователь есть в Slack
                slack_user = next((u for u in slack_users if u.get('profile', {}).get('email') == user['email']), None)
                if slack_user:
                    logging.info(f"[INFO] Can disable: {user['username']} (Slack email: {user['email']})")
                else:
                    logging.info(f"[INFO] No corresponding Slack user found for {user['username']}")
        else:
            logging.error("[ERROR] No Slack users found or failed to fetch them.")
    else:
        logging.info("[INFO] No disabled users found in AD.")

if __name__ == '__main__':
    main()
