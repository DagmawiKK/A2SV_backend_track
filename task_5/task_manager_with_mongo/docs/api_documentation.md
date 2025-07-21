# 📘 Task Management API Documentation

This API allows users to manage tasks through a RESTful interface. You can create, read, update, and delete tasks. It is built using Go, Gin, and MongoDB.

---

## 🔗 Published API Documentation

📄 View the full Postman documentation here:  
👉 [https://documenter.getpostman.com/view/39808423/2sB34ijK67](https://documenter.getpostman.com/view/39808423/2sB34ijK67)

---

## 🛠️ How to Run the API

### 1. **Clone the repository**

```bash
git clone https://github.com/tsigemariamzewdu/Task-Management-REST-API-using-Go-and-Gin-Framework
cd Task-Management-REST-API-using-Go-and-Gin-Framework
```

### 2. **Create a `.env` file**

Create a `.env` file in the root directory and add your MongoDB URI like this:

```env
mongo_url=mongodb+srv://<username>:<password>@<cluster>.mongodb.net/?retryWrites=true&w=majority
```

> 🔐 **Note:** Replace `<username>`, `<password>`, and `<cluster>` with your actual MongoDB credentials.

### 3. **Install dependencies**

```bash
go mod tidy
```

### 4. **Run the server**

```bash
go run main.go
```

The server will start on:

```
http://localhost:8080
```

---

## 📬 API Endpoints

### 📥 `GET /tasks`

**Description:** Get a list of all tasks.  
**Response:**

```json
[
  {
    "id": "60b8f2b1f12a2c3b5c8f2b1a",
    "title": "Morning Plan",
    "description": "Finish task 4",
    "due_date": "2025-07-20",
    "status": "in-progress"
  }
]
```

---

### 🔍 `GET /tasks/:id`

**Description:** Get a task by its MongoDB Object ID.  
**Params:**

* `id` (string): MongoDB Object ID  

**Response:**

```json
{
  "id": "60b8f2b1f12a2c3b5c8f2b1a",
  "title": "Morning Plan",
  "description": "Finish task 4",
  "due_date": "2025-07-20",
  "status": "in-progress"
}
```

---

### ➕ `POST /tasks`

**Description:** Create a new task.  
**Request Body:**

```json
{
  "title": "New Task",
  "description": "Complete project",
  "due_date": "2025-07-31",
  "status": "not-started"
}
```

**Response:** Returns the created task with its MongoDB Object ID.

---

### 🔄 `PUT /tasks/:id`

**Description:** Update a task by its MongoDB Object ID.  
**Params:**

* `id` (string): MongoDB Object ID  

**Request Body:**

```json
{
  "title": "Updated Title",
  "description": "Updated Description",
  "due_date": "2025-08-01",
  "status": "completed"
}
```

**Response:** Updated task object.

---

### ❌ `DELETE /tasks/:id`

**Description:** Delete a task by its MongoDB Object ID.  
**Params:**

* `id` (string): MongoDB Object ID  

**Response:**

```json
{
  "message": "Task deleted"
}
```

---

## 📎 Status Enum

The `status` field must be one of the following values:

* `not-started`
* `in-progress`
* `completed`


