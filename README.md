# TaskTracer CLI  / CLI для управления задачами 

[![English](https://img.shields.io/badge/English-🇬🇧-blue?style=flat-square)](README.md#english-version)
[![Русский](https://img.shields.io/badge/Русский-🇷🇺-red?style=flat-square)](README.md#русская-версия)

## Tech Stack
![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat-square&logo=go)
![JSON](https://img.shields.io/badge/JSON-Data_Storage-000000?style=flat-square&logo=json)
![CLI](https://img.shields.io/badge/CLI-Interface-4EAA25?style=flat-square&logo=terminal)


<a id="english-version"></a>
## 🇬🇧 English Version

### Description
TaskTracer CLI is a simple command-line task manager that helps you organize your daily tasks efficiently.

### Features
-  Add, update, delete tasks
- Mark tasks as todo/in progress/done
- Filter tasks by status
- Persistent storage in JSON format
- Colorful console interface

### Installation
1. Make sure you have [Go installed](https://golang.org/dl/)
2. Clone the repository:
   ```bash
   git clone https://github.com/Vladik22611/TaskTracerCLI.git
   cd TaskTracerCLI
   ```
3. Build and run:
    ```bash
   go build -o tasktracercli
   ./tasktracercli
   ```
### Quick start 
   ``` bash
   # Add a new task
   task-cli add "Buy groceries"
   
   # List all tasks
   task-cli list
   
   # Mark task as in progress
   task-cli mark-in-progress 1
   
   # Delete a task
   task-cli delete 1
   ```
### Full Command List
``` text
add <description>      Add new task
list [status]         List tasks (filter by status)
update <ID> <desc>    Update task description
delete <ID>           Delete a task
mark-in-progress <ID> Change status to 'in progress'
mark-done <ID>        Mark task as done
help                  Show help
exit                  Exit program
```
<a id="русская-версия"></a>
## 🇷🇺 Русская версия
### Описание
TaskTracer CLI - это простой менеджер задач для командной строки, помогающий эффективно организовывать ежедневные задачи.

### Возможности
- Добавление, изменение, удаление задач
- Изменение статуса (todo/in progress/done)
- Фильтрация задач по статусу
- Хранение данных в JSON-файле
- Цветной интерфейс в консоли

## Установка
1. Установите [Go](https://golang.org/dl/)
2. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/yourusername/TaskTracerCLI.git
   cd TaskTracerCLI
   ```
3. Запустите программу:
    ```bash
   go build -o tasktracercli
   ./tasktracercli
   ```
## Быстрый старт
```bash
# Добавить задачу
task-cli add "Купить продукты"

# Показать все задачи
task-cli list

# Изменить статус задачи
task-cli mark-in-progress 1

# Удалить задачу
task-cli delete 1
```
## Все команды
```text
add <описание>       Добавить задачу
list [статус]       Список задач (фильтр по статусу)
update <ID> <опис>  Изменить описание
delete <ID>         Удалить задачу
mark-in-progress <ID> Статус "в работе"
mark-done <ID>      Статус "завершено"
help                Справка
exit                Выход
```
## Author / Автор:

Grechkin Vladislav  
Гречкин Владсислав  

[![Telegram](https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white)](https://t.me/g_vladislav22)
[![GitHub](https://img.shields.io/badge/GitHub-181717?style=for-the-badge&logo=github&logoColor=white)](https://github.com/Vladik22611)

   
