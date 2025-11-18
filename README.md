# 🔗 URL Shortener (Go + Redis)

A fast and lightweight **URL Shortener** built with **Go (Golang)** and **Redis**.
This project exposes two APIs:

* **POST /shorten** — Create a short URL
* **GET /{shortCode}** — Redirect to the original long URL

The application ensures **idempotency**, handles **collisions**, and stores mappings in Redis using a simple, efficient key structure.

---

## 🚀 Features

* ⚡ Ultra-fast key-value storage using Redis
* 🔐 Idempotent shortening (same long URL → same short code)
* 🌀 Collision handling using salted hash generation
* 🔁 URL redirection via short code
* 🧹 Clean and modular project structure (handlers, config, utils)
* 🧪 Easy to extend

---

## 🛠 Technology Stack

* **Programming Language:** Go (Golang)
* **Database:** Redis
* **Dependencies:**

  * `github.com/redis/go-redis/v9`

---

## 📦 Installation & Setup

### 1️⃣ Clone the repository

```bash
git clone https://github.com/YOUR_USERNAME/url-shortner.git
cd url-shortner
```

### 2️⃣ Install dependencies

```bash
go mod tidy
```

### 3️⃣ Start Redis locally

If you have Redis installed:

```bash
redis-server
```

### 4️⃣ Run the Go server

```bash
go run main.go
```

Server runs at:

```
http://localhost:8080
```

---

## 📡 API Endpoints

### **1. Shorten a URL**

**POST /shorten**

#### Request Body:

```json
{
  "url": "https://example.com/very-long-url"
}
```

#### Response:

```json
{
  "short_url": "http://localhost:8080/aB93ks"
}
```

---

### **2. Redirect to the Original URL**

**GET /{shortCode}**

Example:

```
GET http://localhost:8080/aB93ks
```

Redirects to:

```
https://example.com/very-long-url
```

---

## 🗂 Project Structure

```
url-shortener/
│── config/
│   └── redis.go        # Redis client setup
│
│── handlers/
│   ├── redirect.go     # GET /{shortCode} handler
│   └── shorten.go      # POST /shorten handler
│
│── utils/
│   └── hash.go         # Short code generator functions
│
│── main.go             # Server setup & routing
│── go.mod
│── go.sum
```

---

## 🔍 How It Works (Internally)

### ✔ Idempotency

Before generating a short code, the app checks if the long URL already has one:

```
GET long:<url> → shortCode
```

If found → returns the existing short URL.

---

### ✔ Collision Handling

If a generated short code already exists, a **salted hash** is generated until a unique one is found.

---

### ✔ Redis Key Schema

| Key Format       | Value     | Purpose                    |
| ---------------- | --------- | -------------------------- |
| `{shortCode}`    | long URL  | Redirection mapping        |
| `long:{longURL}` | shortCode | Idempotency reverse lookup |

---

