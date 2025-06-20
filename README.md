<h1 align="center">📌 GoTaskify — Simple Task REST API in Go</h1>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.18+-00ADD8?style=for-the-badge&logo=go" alt="Go Version">
  <img src="https://img.shields.io/badge/SQLite-Embedded_DB-003B57?style=for-the-badge&logo=sqlite" alt="SQLite">
  <img src="https://img.shields.io/badge/Mux-Gorilla-blue?style=for-the-badge" alt="Gorilla Mux">
  <img src="https://img.shields.io/badge/GORM-ORM-orange?style=for-the-badge" alt="GORM">
</p>

<p align="center">
  A clean and minimalistic RESTful API for managing Tasks using <strong>Go</strong>, <strong>Gorilla Mux</strong>, <strong>GORM</strong>, and <strong>SQLite</strong>.<br>
  Easily create, read, update, and delete task items with a simple HTTP interface.
</p>

---

<p align="center">
  <img src="https://github.com/sidharthchauhan/GoTaskify/raw/main/demo.gif" alt="GoTaskify demo" width="600"/>
  <br>
  <em>Demo: Creating and listing tasks</em>
</p>

---

## ⚡ Quick Start

### 📥 1. Clone the Repository

```bash
git clone https://github.com/sidharthchauhan/GoTaskify.git
cd GoTaskify
```

### 📦 2. Install Dependencies

```bash
go mod tidy
```

### ▶️ 3. Run the Application

```bash
go run .
```

Now your API will be running at 👉 [http://localhost:8080](http://localhost:8080)

---

### 🐳 Run with Docker Compose (Optional)

```bash
docker-compose up --build
```

---

## ✨ Features

- ✅ Full CRUD for Task items  
- 🧠 Automatic database migration (no manual setup)  
- 🗂 Lightweight SQLite storage  
- 🧪 Effortless testing with REST clients like Postman or curl  
- 🐳 Docker support for easy deployment  
- 🧱 Modular, beginner-friendly project structure  
- 🚀 Easily extensible with new fields or logic  

---

## 🌐 API Endpoints

| Method | Endpoint         | Description                  |
|--------|------------------|------------------------------|
| GET    | `/tasks`         | List all tasks               |
| GET    | `/tasks/{id}`    | Get a single task by ID      |
| POST   | `/tasks`         | Create a new task            |
| PUT    | `/tasks/{id}`    | Update an existing task      |
| DELETE | `/tasks/{id}`    | Delete a task                |
| GET    | `/status`        | Health check endpoint        |

---

### 🛠️ Example Usage

**Create a Task**
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Buy groceries","completed":false}'
```

**List Tasks**
```bash
curl http://localhost:8080/tasks
```

**Update a Task**
```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Buy groceries and milk","completed":true}'
```

---

## 🧪 Sample Task JSON

```json
{
  "title": "Buy groceries",
  "completed": false
}
```

---

## 🧱 Project Structure

```
GoTaskify/
├── main.go                 # Application entry point
├── models/
│   └── task.go             # Task model definition
├── routes/
│   └── taskRoutes.go       # API route definitions
├── controllers/
│   └── taskController.go   # Handler logic for endpoints
└── data/
    └── tasks.db            # SQLite database (auto-created)
```

---

## 🔧 Environment Variable

| Variable  | Description                    | Default           |
|-----------|--------------------------------|-------------------|
| `DB_PATH` | Path to SQLite database file   | `data/tasks.db`   |

To use a custom DB path:

```bash
export DB_PATH=your/custom/path.db
go run .
```

---

## 👨‍💻 Development Tips

- 💾 Schema auto-migrates on startup (no SQL setup needed)  
- 🔁 Restart the app to apply code changes  
- 🧪 Use Postman or `curl` to test endpoints quickly  
- ➕ Add more fields in `models/task.go` as needed  

---

## 🤝 Contributing

Pull requests are welcome!  
Feel free to fork the repo, suggest improvements, or open issues.

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).

---

## Running Tests

To run all tests:

```
go test ./test/...
```

---

<p align="center">
  Made with 💙 by <a href="https://github.com/sidharthchauhan">Sidharth Chauhan</a>
</p>
