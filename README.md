# <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" width="25px" height="25px"/> Todo List Service - GDGoC 

Todo List Service adalah layanan backend sederhana untuk mengelola daftar tugas (To-Do List). Layanan ini dikembangkan menggunakan Golang dengan Gin framework.


## Fitur
- ✅ **Menambahkan tugas baru** (`POST /tasks`)
- 📋 **Mengambil semua tugas** (`GET /tasks`)
- 🔍 **Mengambil tugas berdasarkan ID** (`GET /tasks/{id}`)
- ✏ **Memperbarui tugas berdasarkan ID** (`PUT /tasks/{id}`)
- 🗑 **Menghapus tugas berdasarkan ID** (`DELETE /tasks/{id}`)

## Teknologi yang Digunakan
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)

## Instalasi dan Menjalankan Project
### 1. Clone Repository
```sh
git clone https://github.com/kennethwn/todo-list-service-gdgoc.git
cd todo-list-service-gdgoc
```

### 2. Install Dependencies
```sh
go mod tidy
```

### 3. Konfigurasi Environment
Buat file `.env` berdasarkan `.env.example` dan sesuaikan konfigurasi database jika diperlukan.

### 4. Menjalankan Server
```sh
go run cmd/web/main.go
```

Server akan berjalan di `http://localhost:8080`

## API Endpoints

### 1. Create Task
- **Endpoint**: `POST /tasks`
- **Request Body**:
  ```json
  {
    "title": "Belajar pointer",
    "description": null,
    "due_date": "2025-02-25T19:50:01.838343Z"
  }
  ```
- **Response**:
  ```json
  {
    "code": 200,
    "data": {
      "id": 1,
      "title": "Belajar pointer",
      "description": null,
      "status": 0,
      "due_date": "2025-02-25T19:50:01.838343Z",
      "created_at": "2025-02-09T12:43:39.567891Z",
      "UpdatedAt": null
    },
    "message": "success create new task"
  }
  ```

### 2. Get All Tasks
- **Endpoint**: `GET /tasks`
- **Response**:
  ```json
  {
    "code": 200,
    "data": [
      {
        "id": 7,
        "title": "Belajar concurrency",
        "description": "Belajar goroutine, channel di golang",
        "status": 1,
        "due_date": "2025-02-05T19:50:01.838343Z",
        "created_at": "2025-02-08T23:06:34.88851Z",
        "UpdatedAt": "2025-02-09T09:50:30.776051Z"
      }
    ],
    "message": "success get all data"
  }
  ```

### 3. Get Task by ID
- **Endpoint**: `GET /tasks/{id}`
- **Response**:
  ```json
  {
    "code": 200,
    "data": {
      "id": 7,
      "title": "Belajar concurrency",
      "description": "Belajar goroutine, channel di golang",
      "status": 1,
      "due_date": "2025-02-05T19:50:01.838343Z",
      "created_at": "2025-02-08T23:06:34.88851Z",
      "UpdatedAt": "2025-02-09T09:50:30.776051Z"
    },
    "message": "success get data by id"
  }
  ```

### 4. Update Task
- **Endpoint**: `PUT /tasks/{id}`
- **Request Body**:
  ```json
  {
    "title": "Belajar Golang Lanjutan",
    "description": "Menyelesaikan tutorial Golang tingkat lanjut",
    "status": 99,
    "due_date": "2024-02-15T12:00:00Z"
  }
  ```
- **Response**:
  ```json
  {
    "code": 200,
    "message": "success update new task"
  }
  ```

### 5. Delete Task
- **Endpoint**: `DELETE /tasks/{id}`
- **Response**:
  ```json
  {
    "code": 200,
    "message": "success delete task"
  }
  ```

## Lisensi
Proyek ini menggunakan lisensi **MIT**.

---
✨ Happy Coding! 🚀