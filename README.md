# 🛍️ Go E-Commerce API (go-ecommerce-with-claude)

[![Go](https://img.shields.io/badge/Go-1.x-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Build](https://img.shields.io/badge/build-passing-brightgreen.svg)](#)
[![sqlc](https://img.shields.io/badge/sqlc-queries-auto-generated-yellow.svg)](https://sqlc.dev/)

> A simple, modular **e-commerce backend API** built with **Go**, **Makefile**, **SQLC**, **Paseto**, and **database migrations** — for learning and practice purposes.  
> Created by [@ranggaAdiPratama](https://github.com/ranggaAdiPratama).

---

## 📖 Table of Contents
- [About](#about)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Environment Variables](#environment-variables)
  - [Database Migration](#database-migration)
  - [Run the Server](#run-the-server)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

---

## 💡 About

This project implements a **backend API for an e-commerce system** using Go.  
It demonstrates key backend concepts — authentication, database handling, migrations, and modular design — with a focus on learning best practices.

> 🧠 **Note:** This is a **practice project**, not a production-ready app.

---

## ✨ Features

✅ User registration and login (Paseto token-based authentication)  
✅ CRUD operations for products  
✅ Order management endpoints  
✅ PostgreSQL database with versioned migrations  
✅ Strongly-typed SQL queries via `sqlc`  
✅ Automated development tasks using `Makefile`  
✅ `.env` configuration support  
✅ Includes Postman collection for quick testing  

---

## 🧰 Tech Stack

| Tool | Description |
|------|--------------|
| **Go** | Primary programming language |
| **PostgreSQL** | Database |
| **sqlc** | Generate type-safe Go code from SQL |
| **Paseto** | Token-based authentication (alternative to JWT) |
| **Makefile** | Command automation |
| **Migrate** | Database schema management |
| **dotenv** | Environment variable loading |

---

## 🚀 Getting Started

### 🧩 Prerequisites
Make sure you have installed:
- [Go](https://go.dev/dl/) ≥ 1.20  
- [PostgreSQL](https://www.postgresql.org/download/)  
- [make](https://www.gnu.org/software/make/)  
- [sqlc](https://docs.sqlc.dev/en/latest/overview/install.html)

### ⚙️ Installation

```bash
# Clone the repository
git clone https://github.com/ranggaAdiPratama/go-ecommerce-with-claude.git

# Move into the project directory
cd go-ecommerce-with-claude

# Copy and edit environment variables
cp .env.example .env
