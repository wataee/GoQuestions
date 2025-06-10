
# 🧠 GoQuestions

**REST API сервис на языке **Go**

## 🚀 Возможности

- Регистрация и авторизация с помощью JWT (access + refresh токены)
- Получение пользовательского профиля
- Просмотр списка вопросов
- Административные функции:
  - Добавление вопроса
  - Удаление пользователя
  - Получение списка всех пользователей
- Swagger-документация

## 🛠 Технологии

- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [JWT](https://github.com/golang-jwt/jwt)
- [Swagger (Swaggo)](https://github.com/swaggo/swag)
- [Validator.v10](https://github.com/go-playground/validator)

## 📁 Структура проекта

```

GoQuestions/
├── cmd/
│   └── main.go                  # Точка входа
├── config/
│   └── config.go                # Конфигурации (например, JWT ключ)
├── internal/
│   ├── admin/                   # Логика и хендлеры администратора
│   ├── user/                    # Логика и хендлеры пользователей
│   ├── questions/               # Работа с вопросами
│   ├── middleware/              # Middleware (Auth, CORS)
│   ├── models/                  # DTO и структуры
│   └── database/repository/    # Работа с базой данных
├── router/
│   └── router.go                # Настройка всех маршрутов
└── docs/                        # Сгенерированные Swagger-документы

````

## 🔐 Аутентификация

- **Access-токен** — действует 2 часа  
- **Refresh-токен** — действует 7 дней  
- Используются роли `user` и `admin` для разграничения доступа

## 📦 Установка и запуск

```bash
git clone https://github.com/wataee/GoQuestions.git
cd GoQuestions
go mod tidy
go run cmd/main.go
````

> Перед запуском укажите переменные окружения

## 📘 Swagger-документация

После запуска доступна по адресу:

```
http://localhost:8080/swagger/index.html
```

## 🔗 API Эндпоинты

### 🧑 Пользователь

| Метод | Путь         | Описание                        |
| ----- | ------------ | ------------------------------- |
| POST  | `/login`     | Регистрация или вход            |
| POST  | `/refresh`   | Обновление токенов              |
| GET   | `/profile`   | Получение профиля (авторизация) |
| GET   | `/questions` | Получение списка вопросов (JWT) |

### 🔐 Админ

| Метод  | Путь                     | Описание                      |
| ------ | ------------------------ | ----------------------------- |
| GET    | `/admin/user_list`       | Получить список пользователей |
| DELETE | `/admin/delete_user/:id` | Удалить пользователя          |
| POST   | `/admin/addquestion`     | Добавить вопрос               |

