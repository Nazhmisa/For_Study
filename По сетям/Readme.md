# 🔥 1️⃣ Сетевые технологии
## 1.1 IP-адресация
📌 IPv4 – 32-битный адрес (например, 192.168.1.1).
📌 IPv6 – 128-битный адрес (например, 2001:db8::ff00:42:8329).

🔹 Классы IP-адресов (IPv4)

![alt text](<Снимок экрана от 2025-02-19 18-00-32.png>)

🔹 Частные и публичные IP
✅ Частные (не выходят в интернет напрямую, используются в локальных сетях):

10.0.0.0 – 10.255.255.255
172.16.0.0 – 172.31.255.255
192.168.0.0 – 192.168.255.255

✅ Публичные – доступны в интернете (выдаются провайдерами).

❓ Вопрос: Чем отличаются частные и публичные IP-адреса?

## 1.2 DHCP (Dynamic Host Configuration Protocol)
🔹 Протокол, который автоматически раздаёт IP-адреса устройствам в сети.

🔹 Как работает DHCP?
1️⃣ Клиент отправляет DHCPDISCOVER (ищет сервер).
2️⃣ DHCP-сервер отвечает DHCPOFFER (предлагает IP).
3️⃣ Клиент отправляет DHCPREQUEST (подтверждает выбор).
4️⃣ DHCP-сервер отвечает DHCPACK (назначает IP).

📌 Команда в Windows:

bash
```
ipconfig /release
ipconfig /renew
```
❓ Вопрос: Что делать, если компьютер не получает IP-адрес по DHCP?

## 1.3 DNS (Domain Name System)
🔹 Преобразует доменные имена (google.com) в IP-адреса (142.250.185.46).

🔹 Типы DNS-записей

![alt text](<Снимок экрана от 2025-02-19 18-03-08.png>)

📌 Команды:

Проверить IP домена:
```
nslookup google.com
```
Проверить конкретный DNS-сервер:
```
nslookup google.com 8.8.8.8
```
❓ Вопрос: Что делать, если сайт не открывается, но пинг до IP-адреса проходит?

## 1.4 MAC-адрес и ARP
📌 MAC-адрес – уникальный адрес сетевой карты (например, 00:1A:2B:3C:4D:5E).
📌 ARP (Address Resolution Protocol) – преобразует IP в MAC.

📌 Команды:

Посмотреть ARP-таблицу:
```
arp -a
```
Очистить ARP-кеш:
```
arp -d
```
❓ Вопрос: Как ARP помогает устройствам в локальной сети обмениваться данными?

## 1.5 TCP и UDP
📌 TCP (Transmission Control Protocol)
✅ Надёжный, устанавливает соединение (3-way handshake).
✅ Используется в HTTP(S), FTP, SSH.

📌 UDP (User Datagram Protocol)
✅ Быстрее, но без подтверждения доставки.
✅ Используется в DNS, VoIP, стриминге.

❓ Вопрос: В каких случаях лучше использовать UDP вместо TCP?

# 🔥 2️⃣ Windows и Linux (основы системного администрирования)
## 2.1 Командная строка Windows (CMD, PowerShell)
🔹 Основные команды CMD

![alt text](<Снимок экрана от 2025-02-19 18-04-52.png>)

📌 PowerShell – более мощный инструмент

Посмотреть IP-адрес:
```
Get-NetIPAddress
```
Остановить процесс по имени:
```
Stop-Process -Name "notepad" -Force
```
❓ Вопрос: Как сбросить DNS-кеш в Windows?

## 2.2 Linux – основные команды

![alt text](<Снимок экрана от 2025-02-19 18-05-54.png>)

📌 Работа с сетью в Linux

Узнать IP-адрес:
```
ip a
```
Проверить маршрут:
```
traceroute google.com
```
Посмотреть открытые порты:
```
netstat -tulnp
```
❓ Вопрос: Чем df -h отличается от du -sh?

## 2.3 Active Directory (AD) – основа корпоративных сетей
🔹 Что такое AD?

Это база данных пользователей и устройств в сети.
Позволяет централизованно управлять учетными записями, доступом и политиками безопасности.
🔹 Основные компоненты AD:
✅ Domain Controller (DC) – сервер, управляющий пользователями и политиками.
✅ Organizational Units (OU) – группы пользователей/компьютеров.
✅ Group Policy (GPO) – правила безопасности и настройки ПК.

📌 Команды для работы с AD в Windows

Проверить, в каком домене ПК:
```
whoami /fqdn
```
Получить список пользователей домена:
```
Get-ADUser -Filter *
```
Посмотреть группы пользователя:
```
net user имя /domain
```
❓ Вопрос: Чем отличается локальная учетная запись от доменной?

## 2.4 Виртуализация (VMware, Hyper-V, VirtualBox)
🔹 Почему это важно?

Позволяет запускать несколько ОС на одном ПК.
Используется для тестирования и серверных решений.
✅ Hyper-V – встроенная в Windows виртуализация (Pro/Enterprise).
✅ VMware Workstation – продвинутая виртуализация.
✅ VirtualBox – бесплатная и удобная.

📌 Основные команды PowerShell для Hyper-V

Создать ВМ:
```
New-VM -Name "LinuxTest" -MemoryStartupBytes 2GB -NewVHDPath "C:\VMs\LinuxTest.vhdx" -NewVHDSizeBytes 20GB
```
Запустить ВМ:
```
Start-VM -Name "LinuxTest"
```
❓ Вопрос: В чём разница между Hyper-V и VMware?

# 🔥 3️⃣ Железо (аппаратная часть компьютера)
## 3.1 Основные компоненты ПК и сервера

![alt text](<Снимок экрана от 2025-02-19 18-09-29.png>)
📌 Различия между HDD и SSD:
✅ HDD – дешевле, но медленнее.
✅ SSD – быстрее, но дороже.

❓ Вопрос: В чем разница между DDR4 и DDR5 RAM?

## 3.2 Сборка ПК и серверов
🔹 Что важно при сборке?
✅ Совместимость комплектующих.
✅ Запас по мощности блока питания.
✅ Охлаждение для процессора и видеокарты.
✅ Количество слотов для RAM и M.2 SSD.

📌 RAID – массивы для защиты данных

![alt text](<Снимок экрана от 2025-02-19 18-09-57.png>)

❓ Вопрос: Какие типы RAID лучше для сервера?

## 3.3 Диагностика и устранение проблем с железом
🔹 Основные инструменты диагностики:
✅ BIOS/UEFI – проверка железа на старте.
✅ S.M.A.R.T. – проверка состояния HDD/SSD (wmic diskdrive get status).
✅ MemTest86 – тест оперативной памяти.
✅ HWMonitor – мониторинг температуры.

📌 Как проверить загрузку процессора и памяти в Windows?
1️⃣ Ctrl + Shift + Esc → Диспетчер задач.
2️⃣ perfmon → Монитор ресурсов.
3️⃣ wmic cpu get loadpercentage – нагрузка CPU через командную строку.

❓ Вопрос: Как понять, что БП не хватает мощности?

# 🔥 4️⃣ Информационная безопасность (ИБ)
## 4.1 Основные угрозы безопасности

![alt text](<Снимок экрана от 2025-02-19 18-11-47.png>)

📌 Защита:
✅ Использовать антивирус (например, Windows Defender, Kaspersky, ESET).
✅ Включать двухфакторную аутентификацию (2FA).
✅ Не переходить по подозрительным ссылкам.
✅ Проверять файлы перед открытием (VirusTotal, SandBox).

❓ Вопрос: Какие типы вредоносного ПО ты знаешь?

## 4.2 Аутентификация и управление доступом
🔹 Способы аутентификации:
1️⃣ Пароли (надежные, сложные, уникальные).
2️⃣ Двухфакторная аутентификация (2FA).
3️⃣ Биометрия (Face ID, отпечаток пальца).
4️⃣ Аппаратные ключи (YubiKey).

📌 Основные принципы безопасности паролей:
✅ Не использовать одинаковые пароли.
✅ Длина пароля 12+ символов.
✅ Использовать менеджеры паролей (Bitwarden, KeePass).

❓ Вопрос: Какие уязвимости есть у SMS-кодов в 2FA?

## 4.3 Сетевая безопасность
🔹 Как защитить сеть?
✅ Использовать шифрование (WPA3 для Wi-Fi).
✅ Включать фаервол (Windows Firewall, iptables в Linux).
✅ Отключать ненужные порты (netstat -an).
✅ Использовать VPN для защиты трафика.

📌 Проверка безопасности сети в Windows:
1️⃣ netstat -an – какие порты открыты.
2️⃣ ipconfig /all – инфо о сети.
3️⃣ arp -a – просмотр ARP-таблицы (поиск спуфинга).

❓ Вопрос: Как определить, что кто-то подключился к твоему Wi-Fi?

## 4.4 Резервное копирование и защита данных
🔹 Виды резервного копирования:

![alt text](<Снимок экрана от 2025-02-19 18-12-40.png>)

📌 Лучшие практики:
✅ Использовать 3-2-1 стратегию (3 копии, 2 носителя, 1 в облаке).
✅ Шифровать резервные копии.
✅ Регулярно тестировать восстановление.

❓ Вопрос: Какие инструменты ты знаешь для резервного копирования?

# 🔥 5️⃣ Автоматизация и скрипты (PowerShell, Bash)
## 5.1 PowerShell (Windows) – Основные команды
PowerShell — мощный инструмент для автоматизации задач в Windows.

🔹 Основные команды:

![alt text](<Снимок экрана от 2025-02-19 18-13-14.png>)

📌 Пример скрипта PowerShell:

```
# Скрипт проверяет доступность сайтов и сохраняет результаты в лог
$sites = @("google.com", "yandex.ru", "microsoft.com")
foreach ($site in $sites) {
    $result = Test-Connection -ComputerName $site -Count 2
    Add-Content -Path "C:\logs\ping_results.txt" -Value "$site is reachable: $($result.Status)"
}
```
❓ Вопрос: Как можно автоматически завершить процесс по его имени?

## 5.2 Bash (Linux) – Основы автоматизации
Bash — стандартный инструмент для автоматизации в Linux.

🔹 Основные команды:

![alt text](<Снимок экрана от 2025-02-19 18-14-04.png>)

📌 Пример Bash-скрипта:

```
#!/bin/bash
* Скрипт проверяет занятое место на диске и отправляет уведомление
THRESHOLD=90
USAGE=$(df / | grep / | awk '{print $5}' | sed 's/%//g')

if [ "$USAGE" -gt "$THRESHOLD" ]; then
    echo "Warning: Disk usage is above $THRESHOLD%" | mail -s "Disk Alert" user@example.com
fi
```
❓ Вопрос: Как можно посмотреть, какие порты слушает система в Linux?

## 5.3 Автоматизация резервного копирования
🔹 PowerShell (Windows) – создание резервной копии файлов:


```
Copy-Item -Path "C:\Users\User\Documents" -Destination "D:\Backup" -Recurse
```
🔹 Bash (Linux) – автоматическое копирование данных с архивацией:
```
tar -czvf /backup/my_backup_$(date +%F).tar.gz /home/user/data
```
❓ Вопрос: Как настроить автоматический запуск скрипта в Linux при загрузке системы?



# 🔥 6️⃣ VPN, NAT и базовая безопасность сети
## 6.1 VPN (Virtual Private Network) – Основы
VPN используется для создания защищенного канала связи через интернет.

🔹 Виды VPN:

PPTP (устаревший, слабая защита)
L2TP/IPSec (более безопасный, но медленный)
OpenVPN (надежный и популярный, использует TLS-шифрование)
WireGuard (новый, быстрый, легкий в настройке)
📌 Как проверить, подключен ли VPN?

В Windows: ipconfig /all – ищем адаптер с VPN-адресом
В Linux: ifconfig или ip a
Через веб-сервис: whoer.net
❓ Вопрос: В чем разница между OpenVPN и WireGuard?



## 6.2 NAT (Network Address Translation)
NAT преобразует частные IP-адреса (локальной сети) в один внешний IP.

🔹 Виды NAT:

SNAT (Source NAT) – подмена исходящего IP-адреса
DNAT (Destination NAT) – подмена адреса получателя (используется для проброса портов)
PAT (Port Address Translation) – многие внутренние IP-адреса используют один внешний IP
📌 Как посмотреть NAT в Windows?

netsh interface ip show config
tracert 8.8.8.8 – если первая строка с 192.168.x.x или 10.x.x.x, значит, NAT активен
📌 Как настроить NAT в Linux?


```
iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
```
❓ Вопрос: В чем разница между SNAT и DNAT?

## 6.3 Firewall и безопасность сети
Брандмауэр (firewall) – фильтрует трафик по правилам.

🔹 Windows Firewall:

wf.msc – графическая утилита
netsh advfirewall set allprofiles state off – отключение (не рекомендуется)
🔹 Linux (iptables / UFW):

iptables -L -v – посмотреть текущие правила
ufw enable – включить UFW
ufw allow 22/tcp – открыть SSH
📌 Пример запрета всех входящих соединений, кроме SSH (Linux):

```
ufw default deny incoming
ufw allow 22/tcp
ufw enable
```
❓ Вопрос: Как открыть определенный порт в Windows Firewall?


# 7️⃣ Диагностика сетевых проблем
## 7.1 Основные команды диагностики сети
📌 1. Проверка конфигурации сети

ipconfig /all (Windows) или ifconfig / ip a (Linux) – показывает IP-адреса, шлюз, DNS
route print (Windows) или ip route (Linux) – таблица маршрутизации
📌 2. Проверка доступности узла (Ping)

ping 8.8.8.8 – проверка доступности Google DNS
ping -t 8.8.8.8 – бесконечный пинг (Windows)
ping -c 5 8.8.8.8 – 5 запросов (Linux/Mac)
📌 3. Отслеживание маршрута (Traceroute)

tracert 8.8.8.8 (Windows)
traceroute 8.8.8.8 (Linux)
📌 4. Проверка DNS

nslookup google.com (Windows)
dig google.com (Linux, требует установки)
📌 5. Проверка открытых портов

netstat -an – показывает активные соединения
telnet 192.168.1.1 80 – проверка доступа к порту
nc -zv 192.168.1.1 80 (Linux)
📌 6. Просмотр сетевой активности

netstat -bn – какие процессы используют интернет
tasklist | findstr chrome.exe – поиск процесса
❓ Вопрос: Что покажет ping 127.0.0.1?

## 7.2 Частые проблемы и их диагностика
🔴 Проблема 1: Нет интернета
✅ Решение:

ipconfig /release && ipconfig /renew – обновить IP (Windows)
dhclient -r && dhclient – обновить IP (Linux)
Проверить доступность шлюза: ping 192.168.1.1
Проверить DNS: nslookup google.com
🔴 Проблема 2: Компьютер не видит другие устройства в сети
✅ Решение:

arp -a – показывает соседние устройства
ping <локальный IP> – проверить связь
net use \\<IP-адрес> – проверить доступ к общим папкам
🔴 Проблема 3: Высокий пинг и потери пакетов
✅ Решение:

ping 8.8.8.8 -n 100 – проверить стабильность
pathping 8.8.8.8 (Windows) – где теряются пакеты
tracert 8.8.8.8 – диагностика маршрута
🔴 Проблема 4: Открыт лишний порт, рискуем безопасностью
✅ Решение:

Проверить открытые порты: netstat -an
Отключить ненужные службы: sc stop <имя_службы>
Закрыть порт в брандмауэре:

```
netsh advfirewall firewall add rule name="Block Port 23" dir=in action=block protocol=TCP localport=23
```
❓ Вопрос: Как проверить, не занимает ли процесс определенный порт?

# 8️⃣ Active Directory (AD) и управление пользователями
## 8.1 Что такое Active Directory?
📌 Active Directory (AD) – это служба каталогов от Microsoft для управления пользователями, компьютерами и ресурсами в корпоративных сетях.

📌 Основные компоненты AD:

Домен (Domain) – единая база пользователей и устройств
Контроллер домена (Domain Controller, DC) – сервер, который управляет AD
Organizational Units (OU) – организационные подразделения для группировки пользователей/устройств
Групповые политики (GPO, Group Policy Object) – правила, применяемые к пользователям и компьютерам
Лес (Forest) – набор связанных доменов
Дерево (Tree) – группа доменов в одном пространстве имен
📌 Команды для проверки работы AD:

whoami – имя пользователя
gpresult /r – проверка примененных GPO
net user <имя_пользователя> – информация о пользователе
dsquery user -name * – поиск всех пользователей в AD
❓ Вопрос: Как узнать, к какому домену подключён компьютер?

## 8.2 Создание и управление пользователями в AD
📌 Создание пользователя через GUI:

Открыть Active Directory Users and Computers (ADUC)
Выбрать OU, нажать New → User
Ввести имя, логин и задать пароль
Настроить группы и политики
📌 Создание пользователя через PowerShell:

```
New-ADUser -Name "Ivan Petrov" -GivenName "Ivan" -Surname "Petrov" -SamAccountName "ipetrov" -UserPrincipalName "ipetrov@domain.local" -Path "OU=Users,DC=domain,DC=local" -AccountPassword (ConvertTo-SecureString "P@ssw0rd!" -AsPlainText -Force) -Enabled $true
```
📌 Добавление пользователя в группу:

```
Add-ADGroupMember -Identity "IT_Users" -Members "ipetrov"
```
📌 Удаление пользователя:

```
Remove-ADUser -Identity "ipetrov"
```
❓ Вопрос: Чем отличается SamAccountName от UserPrincipalName?

## 8.3 Группы в Active Directory
🔹 Локальные (Local) – действуют только на одном сервере
🔹 Глобальные (Global) – используются в пределах одного домена
🔹 Универсальные (Universal) – работают во всех доменах леса

📌 Просмотр групп пользователя:

```
Get-ADUser -Identity ipetrov -Properties MemberOf
```
📌 Создание группы и добавление пользователей:

```
New-ADGroup -Name "Developers" -GroupScope Global -Path "OU=Groups,DC=domain,DC=local"
Add-ADGroupMember -Identity "Developers" -Members "ipetrov"
```
8.4 Групповые политики (GPO)
📌 Примеры GPO:
✅ Отключение USB-накопителей
✅ Назначение домашней страницы браузера
✅ Автоматическая блокировка экрана через 10 минут бездействия

📌 Команды для работы с GPO:

gpedit.msc – редактирование локальных политик
gpresult /r – просмотр применённых политик
gpupdate /force – принудительное обновление политик
❓ Вопрос: Что делает gpupdate /force и когда его использовать?

# 9️⃣ Бэкапы и отказоустойчивость
## 9.1 Важность резервного копирования (Backup)
📌 Бэкапы нужны для:
✅ Защиты от потери данных (сбой оборудования, вирусы, ошибки пользователей)
✅ Восстановления данных в случае аварии
✅ Защиты от атак (например, шифровальщиков)

📌 Основные типы бэкапов:
1️⃣ Полный (Full Backup) – копирование всех данных 🏋‍♂ (много места)
2️⃣ Дифференциальный (Differential Backup) – сохраняются изменения после полного бэкапа 📈
3️⃣ Инкрементальный (Incremental Backup) – сохраняются только новые изменения с последнего бэкапа 🔄

❓ Вопрос: Чем инкрементальный бэкап отличается от дифференциального?

## 9.2 Способы резервного копирования
📌 Локальные бэкапы:
🔹 HDD, SSD, NAS (Synology, QNAP)
🔹 Серверные решения (Windows Server Backup, Veeam)

📌 Облачные бэкапы:
☁ OneDrive, Google Drive, Dropbox – для личных данных
☁ AWS S3, Azure Backup – для корпоративных данных

📌 RAID ≠ Бэкап!
RAID защищает от отказа дисков, но не спасает от удаления файлов или вирусов 🛑

## 9.3 Бэкап в Windows
📌 Windows Server Backup (WSB)
✅ Встроенное средство резервного копирования
✅ Поддержка планирования задач
✅ Можно делать полные и инкрементальные бэкапы

📌 Создание бэкапа через командную строку (wbadmin)

```
wbadmin start backup -backupTarget:D: -include:C: -allCritical -quiet
```
📌 Восстановление файлов из бэкапа:

```
wbadmin get versions
wbadmin start recovery -version:04/18/2023-10:00 -itemType:File -items:C:\Data
```
❓ Вопрос: В чем разница между образом системы и бэкапом отдельных файлов?

## 9.4 Отказоустойчивость (High Availability, HA)
📌 Методы повышения отказоустойчивости:
✅ RAID-массивы (RAID 1, 5, 10)
✅ Кластеризация серверов (Failover Cluster)
✅ Резервные каналы связи (двойной интернет, VPN)
✅ Автоматическое переключение (Failover)

📌 RAID Уровни:
🔹 RAID 0 – скорость, но нет защиты 🚀
🔹 RAID 1 – зеркалирование 🪞 (данные дублируются)
🔹 RAID 5 – компромисс скорости и надежности 🔄
🔹 RAID 10 – скорость + надежность 💪

❓ Вопрос: Какой RAID использовать для базы данных?

## 9.5 Мониторинг отказов и тестирование бэкапов
📌 Что важно проверять?
✅ Делать тестовые восстановления данных
✅ Следить за состоянием жестких дисков (S.M.A.R.T.)
✅ Контролировать использование ресурсов серверов

📌 Инструменты мониторинга:
🔹 Zabbix – мониторинг серверов и сети
🔹 PRTG – отслеживание доступности сервисов
🔹 Veeam One – мониторинг бэкапов

❓ Вопрос: Как часто нужно тестировать бэкапы?

# 🔟 Информационная безопасность и защита данных
## 10.1 Основные принципы информационной безопасности (ИБ)
🔹 Конфиденциальность – доступ к данным только у авторизованных пользователей 🛑🔑
🔹 Целостность – защита данных от изменений и повреждений 🏛
🔹 Доступность – данные должны быть доступны в нужное время 🚀

⚠️ Основные угрозы ИБ:
❌ Вирусы, трояны, шифровальщики (Ransomware)
❌ Фишинг (мошеннические письма, сайты)
❌ DDoS-атаки (перегрузка серверов)
❌ Внутренние угрозы (недобросовестные сотрудники)

## 10.2 Антивирусы и Endpoint Protection
📌 Популярные антивирусные решения:
✅ Kaspersky, ESET, Bitdefender – защита от вирусов 🛡
✅ Windows Defender – встроенный антивирус от Microsoft 🏢
✅ CrowdStrike, SentinelOne – продвинутая защита бизнеса

📌 Команды проверки Windows Defender:

```
Get-MpComputerStatus  # Проверка статуса антивируса
Start-MpScan -ScanType QuickScan  # Быстрое сканирование
```
❓ Вопрос: Чем отличается антивирус от EDR (Endpoint Detection & Response)?

## 10.3 Фаерволы (Firewall) и контроль трафика
📌 Фаервол (Брандмауэр) фильтрует сетевой трафик 🔥
✅ Аппаратные: Cisco ASA, FortiGate, Palo Alto
✅ Программные: Windows Firewall, iptables (Linux)

📌 Команды настройки фаервола в Windows:

```
netsh advfirewall set allprofiles state on  # Включить фаервол
netsh advfirewall firewall add rule name="Block_Telnet" dir=in action=block protocol=TCP localport=23  
```
📌 Команды настройки iptables в Linux:

```
sudo iptables -A INPUT -p tcp --dport 22 -j ACCEPT  # Разрешить SSH
sudo iptables -A INPUT -p tcp --dport 23 -j DROP  # Заблокировать Telnet
```
❓ Вопрос: В чем разница между Stateful и Stateless фаерволами?

## 10.4 Аутентификация, авторизация и контроль доступа
📌 Методы аутентификации:
🔹 Логин/пароль – стандарт, но уязвим 🛑
🔹 2FA/MFA – дополнительный уровень защиты (Google Authenticator, SMS)
🔹 Biometric (Face ID, отпечаток пальца)

📌 Политики паролей (Windows GPO):

```
net accounts /maxpwage:30  # Срок действия пароля – 30 дней
net accounts /minpwlen:12  # Минимальная длина – 12 символов
```
❓ Вопрос: Почему 2FA не всегда безопасен?

## 10.5 Шифрование данных
📌 Методы шифрования:
✅ Symmetric (AES, DES) – один ключ 🗝
✅ Asymmetric (RSA, ECC) – открытый и закрытый ключи 🔑

📌 Шифрование файлов в Windows (BitLocker):

```
manage-bde -on C: -RecoveryPassword
```
📌 Шифрование в Linux (LUKS):

```
cryptsetup luksFormat /dev/sdb
```
❓ Вопрос: В чем разница между хэшированием и шифрованием?

## 10.6 SIEM и мониторинг безопасности
📌 SIEM-системы (Security Information and Event Management):
✅ Splunk – анализ логов, выявление атак
✅ ELK Stack (Elasticsearch + Logstash + Kibana) – сбор и визуализация данных
✅ Microsoft Sentinel – SIEM от Microsoft

❓ Вопрос: Какие логи важны для анализа атак?


# 11 Администрирование серверов и виртуализация
## 11.1 Основы администрирования Windows Server
📌 Основные роли Windows Server:
✅ Active Directory (AD DS) – управление пользователями и группами
✅ DNS – преобразование доменных имен в IP-адреса
✅ DHCP – автоматическая выдача IP-адресов
✅ File Server – общий доступ к файлам и папкам

📌 Команды Windows Server:

```
Get-ADUser -Filter *  # Вывести всех пользователей AD
Get-ADGroupMember "Domain Admins"  # Проверить, кто в группе администраторов
```
📌 Команда для перезапуска сервера:

```
shutdown /r /t 0
```
❓ Вопрос: Какие группы есть в Active Directory?

11.2 Основы администрирования Linux Server
📌 Основные службы Linux-сервера:
✅ OpenSSH – удаленное подключение по SSH
✅ Apache/Nginx – веб-сервер
✅ Samba – файлообмен между Windows и Linux
✅ Cron – автоматизация задач

📌 Основные команды Linux-сервера:

```
sudo systemctl restart ssh  # Перезапуск SSH
sudo systemctl status apache2  # Проверить статус веб-сервера
```
📌 Создание пользователя:

```
sudo useradd -m newuser
sudo passwd newuser
```
❓ Вопрос: Чем отличается Debian от RHEL?

11.3 Основы виртуализации (Hyper-V, VMware, VirtualBox, KVM)
📌 Типы виртуализации:
✅ Гипервизоры 1-го типа – работают напрямую с железом (VMware ESXi, Hyper-V, KVM)
✅ Гипервизоры 2-го типа – работают внутри ОС (VirtualBox, VMware Workstation)

📌 Основные команды Hyper-V (Windows):

```
Get-VM  # Посмотреть список ВМ
New-VM -Name "TestVM" -MemoryStartupBytes 2GB -Path "C:\VMs"
```
📌 Основные команды KVM (Linux):

```
virsh list --all  # Список виртуальных машин
virsh start VM_NAME  # Запуск ВМ
```
❓ Вопрос: Что лучше для сервера – Hyper-V или VMware?

11.4 Резервное копирование и восстановление (Backup & Recovery)
📌 Методы бэкапа:
✅ Full Backup – полный снимок данных
✅ Incremental Backup – копирование только изменений
✅ Differential Backup – разница с последним полным бэкапом

📌 Резервное копирование в Windows:

```
wbadmin start backup -backupTarget:E: -include:C:
```
📌 Резервное копирование в Linux (tar + rsync):
```
tar -cvzf backup.tar.gz /home/user
rsync -av /home/user remote:/backup/
```
❓ Вопрос: Что лучше – локальный бэкап или облачный?

11.5 Контейнеризация (Docker, Kubernetes)
📌 Зачем Docker?
✅ Легче, чем виртуальные машины
✅ Быстрое развертывание приложений
✅ Удобная изоляция окружений

📌 Основные команды Docker:

```
docker ps  # Посмотреть работающие контейнеры
docker run -d -p 80:80 nginx  # Запустить контейнер Nginx
```
📌 Основные команды Kubernetes:

```
kubectl get pods  # Посмотреть все поды
kubectl apply -f deployment.yaml  # Запустить приложение
```
❓ Вопрос: В чем разница между Docker и Kubernetes?

# 🌩️ 12. Облачные технологии (AWS, Azure, Google Cloud)
## 12.1 Что такое облачные технологии?
📌 Основные модели облаков:
✅ IaaS (Infrastructure as a Service) – виртуальные машины, сети (AWS EC2, Azure VM)
✅ PaaS (Platform as a Service) – готовая среда для разработки (AWS Elastic Beanstalk, Azure App Services)
✅ SaaS (Software as a Service) – готовый софт в облаке (Google Docs, Office 365)

❓ Вопрос: В чем разница между IaaS, PaaS и SaaS?

## 12.2 Основные сервисы AWS
📌 Популярные AWS-сервисы:
✅ EC2 – виртуальные серверы
✅ S3 – облачное хранилище
✅ RDS – управляемая база данных
✅ Lambda – серверлесс-функции

📌 Команды AWS CLI:

```
aws s3 ls  # Просмотр содержимого бакета
aws ec2 describe-instances  # Список ВМ
```
❓ Вопрос: Какой сервис AWS используют для хранения бэкапов?

## 12.3 Основные сервисы Azure
📌 Популярные сервисы Azure:
✅ Azure Virtual Machines (VM) – виртуальные серверы
✅ Azure Blob Storage – облачное хранилище
✅ Azure SQL Database – управляемая база данных
✅ Azure Functions – серверлесс-функции

📌 Команды Azure CLI:

```
az vm list  # Список виртуальных машин
az storage account list  # Список хранилищ
```
❓ Вопрос: В чем разница между Azure Blob Storage и AWS S3?

## 12.4 Основные сервисы Google Cloud
📌 Популярные сервисы Google Cloud:
✅ Compute Engine – виртуальные машины
✅ Cloud Storage – облачное хранилище
✅ BigQuery – анализ больших данных
✅ Cloud Functions – серверлесс-функции

📌 Команды gcloud CLI:

```
gcloud compute instances list  # Список ВМ
gcloud storage buckets list  # Список хранилищ
```
❓ Вопрос: В чем разница между Google Cloud и AWS?

## 12.5 Облачные базы данных
📌 Основные облачные СУБД:
✅ Amazon RDS (MySQL, PostgreSQL, MSSQL)
✅ Azure SQL Database
✅ Google Cloud SQL

📌 Подключение к RDS из терминала:

```
mysql -h mydb-instance.amazonaws.com -u admin -p
```
❓ Вопрос: В чем преимущество облачных баз данных?

## 12.6 Безопасность в облаках
📌 Основные принципы безопасности:
✅ Многофакторная аутентификация (MFA)
✅ Шифрование данных (AWS KMS, Azure Key Vault)
✅ Контроль доступа (IAM, RBAC, ACL)

📌 Пример политики IAM (AWS):

```
{
  "Effect": "Allow",
  "Action": "s3:ListBucket",
  "Resource": "arn:aws:s3:::my-bucket"
}
```
❓ Вопрос: Как защитить данные в облаке?

# 🚀 13. DevOps-инструменты и инфраструктура как код (IaC)
## 13.1 Что такое DevOps?
📌 DevOps – это практика объединения разработки (Dev) и операционной поддержки (Ops) для ускорения развертывания и улучшения качества ПО.

Основные принципы DevOps:
✅ Автоматизация – минимизация ручных действий
✅ CI/CD – непрерывная интеграция и доставка
✅ Мониторинг и логирование – контроль состояния системы
✅ IaC (Infrastructure as Code) – управление инфраструктурой через код

❓ Вопрос: Что такое DevOps и почему он важен?

## 13.2 CI/CD: Непрерывная интеграция и доставка
📌 CI (Continuous Integration) – частые слияния кода в репозиторий с автоматическими тестами.
📌 CD (Continuous Deployment/Delivery) – автоматизированный процесс развертывания обновлений в продакшн.

Популярные инструменты CI/CD:
✅ Jenkins – мощный open-source инструмент для автоматизации
✅ GitHub Actions – CI/CD внутри GitHub
✅ GitLab CI/CD – встроенный в GitLab CI/CD
✅ CircleCI, Travis CI – альтернативные решения

📌 Пример .gitlab-ci.yml для CI/CD в GitLab:

```
stages:
  - build
  - test
  - deploy

build:
  script:
    - echo "Building project..."

test:
  script:
    - echo "Running tests..."

deploy:
  script:
    - echo "Deploying to production..."
```
❓ Вопрос: В чем разница между CI и CD?

## 13.3 Docker и контейнеризация
📌 Docker – инструмент для упаковки приложений и их зависимостей в контейнеры.

✅ Облегчает переносимость приложений
✅ Позволяет запускать несколько сервисов изолированно
✅ Упрощает масштабирование

📌 Основные команды Docker:

```
docker pull ubuntu  # Загрузить образ Ubuntu  
docker run -d -p 8080:80 nginx  # Запустить контейнер с Nginx  
docker ps  # Список запущенных контейнеров  
docker stop <container_id>  # Остановить контейнер  
```
❓ Вопрос: В чем отличие контейнера от виртуальной машины?

## 3.4 Kubernetes (k8s)
📌 Kubernetes – платформа для оркестрации контейнеров.

✅ Автоматическое масштабирование и балансировка нагрузки
✅ Самовосстановление контейнеров
✅ Гибкое управление конфигурацией через манифесты

📌 Пример pod.yaml для Kubernetes:

```
apiVersion: v1
kind: Pod
metadata:
  name: my-app
spec:
  containers:
    - name: my-container
      image: nginx
      ports:
        - containerPort: 80
```
📌 Основные команды kubectl:

```
kubectl apply -f pod.yaml  # Развернуть под  
kubectl get pods  # Список подов  
kubectl logs <pod_name>  # Логи пода  
kubectl delete pod <pod_name>  # Удалить под  
```
❓ Вопрос: Зачем использовать Kubernetes, если есть Docker?

## 13.5 Terraform и Ansible (IaC)
📌 Infrastructure as Code (IaC) – управление инфраструктурой через код.

✅ Terraform – декларативный инструмент для создания инфраструктуры
✅ Ansible – инструмент автоматизации конфигурации

📌 Пример Terraform-конфигурации (main.tf) для создания VM в AWS:

```
provider "aws" {
  region = "us-east-1"
}

resource "aws_instance" "example" {
  ami           = "ami-12345678"
  instance_type = "t2.micro"
}
```
📌 Пример Ansible-плейбука (playbook.yml) для установки Nginx:

```
- hosts: all
  tasks:
    - name: Install Nginx
      apt:
        name: nginx
        state: present
```
❓ Вопрос: В чем разница между Terraform и Ansible?

