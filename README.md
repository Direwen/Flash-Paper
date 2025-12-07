# ⚡ FlashPaper

> **Zero-Knowledge, Self-Destructing Message Sharing Service.**
> *Chingu Voyage Tier 3 (Fullstack) Submission*

[![Go](https://img.shields.io/badge/Backend-Go-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Nuxt 3](https://img.shields.io/badge/Frontend-Nuxt_3-00DC82?style=flat&logo=nuxt.js)](https://nuxt.com/)
[![Docker](https://img.shields.io/badge/Deployment-Docker-2496ED?style=flat&logo=docker)](https://www.docker.com/)
[![PostgreSQL](https://img.shields.io/badge/Database-PostgreSQL-336791?style=flat&logo=postgresql)](https://www.postgresql.org/)

## 🚀 Live Demo

* **Deployed App:** [https://flash-paper-zeta.vercel.app/](https://flash-paper-zeta.vercel.app/)
* **Repository:** [https://github.com/Direwen/Flash-Paper](https://github.com/Direwen/Flash-Paper)

## 📸 Screenshots

| Home | Dashboard |
|:----:|:---------:|
| ![Home](docs/design/ui/home_page.png) | ![Dashboard](docs/design/ui/dashboard_page.png) |

| View (Locked) | View (Success) | View (Expired) |
|:-------------:|:--------------:|:--------------:|
| ![Locked](docs/design/ui/view_page_locked.png) | ![Success](docs/design/ui/view_page_success.png) | ![Expired](docs/design/ui/view_page_expired.png) |

| Encryption Flow | About | Mobile |
|:---------------:|:-----:|:------:|
| ![Encryption](docs/design/ui/encryption_flow.png) | ![About](docs/design/ui/about_page.png) | ![Mobile](docs/design/ui/home_page_small_screen.png) |

---

### 🔑 Evaluator Credentials

To test the **Dashboard** and **Protected Routes**, use this account:
* **Email:** `evaluator@chingu.io`
* **Password:** `chingu123`

---

## 📖 Overview

**FlashPaper** is a secure "Pastebin" alternative designed for sending sensitive data (passwords, API keys, config files) over insecure channels.

Unlike standard databases, FlashPaper operates on a **Zero-Knowledge** architecture. Encryption happens entirely in the browser (Client-Side) using **AES-256-GCM**. The server stores only the encrypted blob and never sees the decryption key, which is passed via the URL fragment (`#`) and never sent over the network.

### Key Features
* **🔥 Self-Destruction:** Snippets can be set to "burn" after **1 view** or specific time limits.
* **🔒 Row-Level Locking:** Uses PostgreSQL `FOR UPDATE` locks to strictly enforce view limits, preventing race conditions even under high concurrency.
* **🧹 The Janitor:** A background Go routine that continuously monitors and scrubs expired records from the database.
* **📊 Dashboard:** Authenticated users can track the status of their active secrets (Active vs. Burnt).
* **📱 Responsive UI:** Built with Nuxt 3 and TailwindCSS, fully optimized for mobile and desktop.

---

## 🏗 System Architecture

### 1. Database Schema (ERD)
The system uses **PostgreSQL** hosted on **Neon**. We strictly enforce referential integrity and use indices for the Janitor process.

```mermaid
erDiagram
    USERS ||--o{ SNIPPETS : "owns"
    
    USERS {
        uuid id PK
        string email UK
        string password_hash
        timestamp created_at
    }

    SNIPPETS {
        uuid id PK
        uuid user_id FK "Nullable"
        text content "AES-256 Encrypted Blob"
        string title
        string language
        int max_views
        int current_views
        timestamp expires_at "Indexed for Janitor"
        timestamp created_at
    }
```

### 2\. Request Lifecycle & Concurrency

The core complexity lies in the **"Reveal & Burn"** logic. To ensure a secret with `MaxViews: 1` is viewed exactly once, we wrap the read operation in a transaction with a Row Lock.

📊 [View Sequence Diagram](docs/design/system/sequence_diagram.png)

```mermaid
sequenceDiagram
    autonumber
    actor User
    participant API as Go API (Gin)
    participant DB as Postgres (Neon)

    Note over API, DB: Critical Section: Race Condition Handling

    User->>API: GET /snippets/:id (Reveal Request)
    
    API->>DB: BEGIN TRANSACTION
    API->>DB: SELECT * FROM snippets WHERE id=... FOR UPDATE
    Note right of DB: ROW LOCKED. Concurrent requests wait.
    
    DB-->>API: Snippet Record

    alt Record Not Found
        API->>DB: ROLLBACK
        API-->>User: 404 Not Found
    else Expired or Burnt
        API->>DB: ROLLBACK
        API-->>User: 410 Gone (Self-Destructed)
    else Valid
        API->>API: Increment View Count
        API->>DB: UPDATE snippets SET current_views = +1
        API->>DB: COMMIT
        Note right of DB: Lock Released
        API-->>User: 200 OK (Encrypted Data)
    end
```

-----

## 🛠 Tech Stack

### Backend

  * **Language:** Go (Golang) 1.23
  * **Framework:** Gin Web Framework
  * **ORM:** GORM
  * **Database:** PostgreSQL (Neon Tech)
  * **Authentication:** JWT (JSON Web Tokens)
  * **Deployment:** Docker Container on **Koyeb**

### Frontend

  * **Framework:** Nuxt 3 (Vue.js)
  * **Styling:** TailwindCSS + MazUI
  * **State Management:** Pinia
  * **Deployment:** **Vercel**

-----

## 🧠 Design Decisions

### Why no "Update" feature?

Tier 3 requires CRUD, but FlashPaper implements **CRD** (Create, Read, Delete). The "Update" operation is intentionally omitted for security integrity. Once a secret is encrypted and armed, allowing modifications would break the chain of trust and potentially allow an attacker to swap the ciphertext.

### Client-Side Encryption

We use the Web Crypto API to generate a key in the browser. This key is appended to the URL as a hash (`#key`). Since hash fragments are not sent to the server in HTTP requests, the backend physically cannot decrypt the data even if subpoenaed.

### The "Lazy" Loading Pattern

The dashboard uses Nuxt's `lazy: true` and `dedupe: 'defer'` configuration. This ensures the UI renders immediately without blocking hydration, preventing "infinite loading" states on slower networks.

-----

## 💻 Local Installation

You can run the entire stack locally using Docker Compose.

### Prerequisites

  * Docker & Docker Compose
  * Node.js (for frontend development)

### Environment Variables

**Backend (`flashpaper/.env`):**
```env
# Server
PORT=8080

# Database (for local Docker setup)
DB_HOST=localhost
DB_USER=flashuser
DB_PASSWORD=flashpassword
DB_NAME=flashpaper_local
DB_PORT=5432
DB_SSLMODE=disable

# Or use a connection string for hosted DB (e.g., Neon)
DB_URL=postgresql://user:password@host/dbname?sslmode=require

# Security
JWT_SECRET=your-secret-key-here
ENCRYPTION_KEY=your-32-byte-encryption-key

# App Config
CLIENT_URL=http://localhost:3000
JANITOR_INTERVAL=10s
TOKEN_EXPIRATION=24h
```

**Frontend (`client/.env`):**
```env
NUXT_PUBLIC_API_BASE=http://localhost:8080
```

### 1\. Clone the Repo

```bash
git clone https://github.com/Direwen/Flash-Paper.git
cd Flash-Paper
```

### 2\. Run with Docker (Recommended)

This spins up a local Postgres instance and the Go API.

```bash
cd flashpaper
docker-compose -f docker-compose.local.yml up --build
```

The API will be available at `http://localhost:8080`.

### 3\. Run Frontend

In a new terminal:

```bash
cd client
npm install
npm run dev
```

The App will be available at `http://localhost:3000`.

-----

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

-----

Built with ❤️ by **Direwen**