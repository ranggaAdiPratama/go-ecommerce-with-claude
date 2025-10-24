# Go Ecommerce with Claude  
_A practice project using Go, Makefile, migrations, `sqlc`, PASETO, etc._

## Overview  
This repository is a **practice implementation** of an ecommerce backend built with Go. The goal is to explore and learn various tools and patterns, including:

- Go (golang) as the primary language  
- Makefile to orchestrate tasks  
- Database migrations  
- `sqlc` for generating type-safe database access  
- PASETO (Platform-Agnostic SEcurity TOkens) for authentication/authorization  
- (Possibly) a REST API for ecommerce operations  
- Structured, maintainable code base — primarily for learning and experimentation  
- FE (coming soon)

⚠️ *Note:* This project is for learning and practicing — it might **not** be production-ready.

## Features  
Here are some core features/characteristics:

- Modular structure using Go packages  
- Database schema and migration files (via Makefile)  
- Auto-generation of Go types from SQL using `sqlc`  
- Authentication using PASETO tokens  
- CRUD operations for typical ecommerce domain entities (e.g., users, products, orders)  
- Makefile targets to simplify workflow (e.g., build, run, migrate)  
- Emphasis on clean code, clarity, and learning toolchains  

## Getting Started  

### Prerequisites  
Make sure you have the following installed on your machine:  
- Go (version 1.x)  
- Make  
- A SQL database (e.g., PostgreSQL)  
- `sqlc` tool (for generating code from SQL)  
- (Optional) any environment-variable manager or `.env` file tool  

### Setup  
1. Clone the repository  
   ```bash
   git clone https://github.com/ranggaAdiPratama/go-ecommerce-with-claude.git  
   cd go-ecommerce-with-claude  
