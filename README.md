🗂️ Daily Worker Roster Management System
A mini backend system built with Go + PostgreSQL + SvelteKit to manage daily worker scheduling, shift requests, approvals, and assignments.

This project is suitable for companies managing shift-based workers such as warehouses, retail, or temporary staffing.

🚀 Features
✅ Worker shift request system
✅ Admin approval and assignment flow
✅ Role-based authorization using JWT
✅ Clean UI using SvelteKit + Tailwind + daisyUI
✅ REST API with Swagger/Postman support
✅ Dockerized for fast local setup

🛠️ Tech Stack
Layer	Tech
Backend	Go (Gin), PostgreSQL, JWT
Frontend	SvelteKit, TailwindCSS, daisyUI
API Docs	Postman Collection
Auth	JWT with role_id context
Containerized	Docker & Docker Compose

📦 Getting Started
1. Clone the repo
```
git clone https://github.com/your-user/daily-worker-roster-management-system.git
cd daily-worker-roster-management-system
```
2. Run with Docker
```
docker-compose up --build
```
Backend: http://localhost:9400

Frontend: http://localhost:5173

📂 Folder Structure
```
├── backend/
│   ├── main.go
│   ├── models/
│   ├── handlers/
│   ├── usecase/
│   ├── repository/
│   └── middleware/
├── frontend/
│   ├── src/
│   ├── routes/
│   ├── lib/
│   └── components/
├── docker-compose.yml
└── README.md
```

🔐 Sample Login Credentials
```
Role	Username	Password
Admin	admin	password
Worker	any existing worker name	password
```

For example: use "Alice" as username and "password" as password to log in as a worker.

📮 API Documentation
Access the full API reference via Postman:
👉 [Postman Collection](https://api.postman.com/collections/13310151-3f1784f2-04de-45a3-ac43-a49f8a0fbf60?access_key=PMAT-01JVJKVHAF5TW4GJF3DZFC3V3E).

Includes examples for login, shift requests, approvals, assignments, etc.
