# 📘 Todo API with Golang + SQLite (No MySQL required)

This project is a simple Todo REST API built using **Golang**, **Gin** framework, and **SQLite** (via native SQL driver) — with no need for MySQL or other external database setups.

---

## 🚀 Features

- CRUD API for managing todos
- Simple SQLite backend (local `.db` file)
- Uses native `database/sql` (no ORM)
- `.env`-based configuration
- Modular folder structure (`controllers`, `models`, `routes`, `config`)

---

## 📦 Project Structure

```
.
├── config/            # DB connection
├── controllers/       # Business logic
├── models/            # Data model
├── routes/            # API routing
├── main.go            # App entrypoint
├── .env.example       # Env variable example
├── go.mod             # Go modules
```

---

## 🛠️ Setup Instructions

### 1. Clone the project
```bash
git clone <repo_url>
cd todo-api-sqlite
```

### 2. Create `.env` file
```bash
cp .env.example .env
```

Default `.env`:
```
PORT=8080
DB_PATH=./todos.db
```

### 3. Run the app
```bash
go mod tidy
go run main.go
```

Server will run at:
```
http://localhost:8080
```

---

## 🧪 API Endpoints

### ✅ Create Todo
```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"My Task"}'
```

### 📄 Get All Todos
```bash
curl http://localhost:8080/todos
```

### ✏️ Update Todo
```bash
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Task","completed":true}'
```

### ❌ Delete Todo
```bash
curl -X DELETE http://localhost:8080/todos/1
```

---

## 📌 Notes

- Database file will be automatically created at the path in `.env` (default: `./todos.db`)
- If DB connection fails, error will be logged in terminal.

---

## 📮 License

MIT
