# ☕ coffeeShop-server

A modular, microservices-based backend system for managing a coffee shop. Built in Go using the Echo framework and MongoDB, with a focus on maintainability, scalability, and clean architecture.

---

## 📐 System Overview

### 📁 Project Structure

```text
coffeeShop-server/
│
├── cmd/               # App entry points
│   ├── identity/      # Identity service (user, staff)
│   ├── order/         # Order service (order, shipping, drink, category)
│
├── pkg/               # Shared and core logic
│   ├── identity/      # Identity domain logic
│   ├── order/         # Order domain logic
│   ├── infra/         # Database (MongoDB), Redis, etc.
│   └── util/          # Common utilities
│
└── docker-compose.yml
```

### 🧠 System Architecture

![system_architecture](https://github.com/user-attachments/assets/580a7084-a26c-4009-bf4e-4453073767fd)


---

🗃 **Why MongoDB?**

Although relational databases like PostgreSQL are commonly preferred for transactional systems, **MongoDB was originally used in this project to**:

- ⚡️ Speed up prototyping and early development  
- 🧩 Model embedded/nested data (e.g., order items, user shipping addresses) in a more natural way  
- 🔄 Avoid strict schemas while iterating on business logic  
- 🧱 Fit a microservices model with each service owning its own collections  

While MongoDB has some limitations for complex transactions, the current architecture enforces structure at the application level and isolates domain logic to maintain long-term maintainability.  
**This choice is retained today for compatibility with the legacy system and to continue leveraging its flexibility.**


## 📬 API Collection

You can try out the APIs using the Postman collection below:

👉 [Postman Collection](https://documenter.getpostman.com/view/12048946/2sB2j999uD)

## 🚀 Getting Started

### 🔧 Prerequisites

- [Go 1.21+](https://go.dev/doc/install)
- [Docker & Docker Compose](https://docs.docker.com/compose/)
- [Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli)

### 🌀 Clone the repository

```bash
git clone https://github.com/dothiphuc81299/coffeeShop-server.git
cd coffeeShop-server
```

## 🐳 Run with Docker Compose

First, start the infrastructure services (MongoDB,  etc.):

```bash
docker-compose up -d
```

Then, in separate terminals, run each service manually:

```bash
go run cmd/identity/main.go
```

```bash
go run cmd/order/main.go
```

---

## 🚀 Deploying to Heroku with Docker

### 📆 Steps to Deploy

1. Login to Heroku Container Registry:

    ```bash
    heroku container:login
    ```

2. Navigate to a service (e.g., identity):

    ```bash
    cd cmd/identity
    ```

3. Create a Heroku app:

    ```bash
    heroku create <your-heroku-app-name>
    ```

4. Push the Docker container to Heroku:

    ```bash
    heroku container:push web -a <your-heroku-app-name>
    ```

5. Release the container on Heroku:

    ```bash
    heroku container:release web -a <your-heroku-app-name>
    ```

6. Repeat for other services if needed.

---

## 👤 Author

- [Phúc Đỗ](https://github.com/dothiphuc81299)

---

## 📝 License

MIT License. See [LICENSE](./LICENSE) for details.

