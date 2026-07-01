# Ovo Gaming Platform Backend Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build a Go + PostgreSQL backend and a Vben admin platform for the Ovo Gaming mini program, including database schema, auth, package/banner/notice management, and order/withdrawal foundations.

**Architecture:** Use a single Go API service with Gin, GORM, and PostgreSQL as the shared data layer for both the WeChat mini program and the admin web app. Keep business logic in service/domain layers so order state transitions, role switching, and admin actions stay testable and reusable.

**Tech Stack:** Go, Gin, GORM, PostgreSQL, JWT, Vue 3, Vben Admin, Ant Design Vue

---

### Task 1: Create PostgreSQL schema

**Files:**
- Create: `server/sql/001_init.sql`
- Create: `server/sql/002_seed.sql`

- [ ] **Step 1: Write the schema SQL**

```sql
CREATE DATABASE ovo_gaming_platform;
```

- [ ] **Step 2: Add the core tables**

```sql
CREATE TABLE IF NOT EXISTS users (...);
```

- [ ] **Step 3: Add seed data for banners, notices, and hot packages**

```sql
INSERT INTO banners (...);
```

- [ ] **Step 4: Verify the SQL is complete**

Run: `go test ./...`
Expected: backend still compiles after SQL is added

- [ ] **Step 5: Commit**

```bash
git add server/sql
git commit -m "feat: add initial platform schema"
```

### Task 2: Scaffold Go backend

**Files:**
- Create: `server/go.mod`
- Create: `server/cmd/api/main.go`
- Create: `server/internal/config/config.go`
- Create: `server/internal/db/db.go`
- Create: `server/internal/model/*.go`
- Create: `server/internal/handler/*.go`
- Create: `server/internal/service/*.go`
- Create: `server/internal/repository/*.go`
- Create: `server/internal/middleware/*.go`

- [ ] **Step 1: Define the module and package layout**

```go
module ovo-gaming/server
```

- [ ] **Step 2: Wire config, DB, router, and health check**

```go
func main() { ... }
```

- [ ] **Step 3: Add auth and admin middleware**

```go
func JWTAuth() gin.HandlerFunc { ... }
```

- [ ] **Step 4: Add a basic API route group**

```go
api := r.Group("/api")
```

- [ ] **Step 5: Verify the backend builds**

Run: `go test ./...`
Expected: packages compile

- [ ] **Step 6: Commit**

```bash
git add server
git commit -m "feat: scaffold backend service"
```

### Task 3: Scaffold Vben admin app

**Files:**
- Create: `admin-web/package.json`
- Create: `admin-web/src/main.ts`
- Create: `admin-web/src/router/index.ts`
- Create: `admin-web/src/layouts/*.vue`
- Create: `admin-web/src/views/dashboard/*.vue`
- Create: `admin-web/src/views/content/*.vue`
- Create: `admin-web/src/views/orders/*.vue`

- [ ] **Step 1: Initialize the Vite app structure**

```ts
import { createApp } from 'vue'
```

- [ ] **Step 2: Add route and layout shells**

```ts
const routes = [...]
```

- [ ] **Step 3: Add placeholder pages for banners, notices, packages, orders**

```vue
<template>轮播图管理</template>
```

- [ ] **Step 4: Verify the frontend boots**

Run: `npm run dev`
Expected: Vben shell loads

- [ ] **Step 5: Commit**

```bash
git add admin-web
git commit -m "feat: scaffold admin web"
```

### Task 4: Connect backend and admin endpoints

**Files:**
- Modify: `server/internal/handler/*.go`
- Modify: `admin-web/src/api/*.ts`
- Modify: `admin-web/src/views/**/*.vue`

- [ ] **Step 1: Add banner, notice, package, and order APIs**

```go
r.GET("/admin/banners", ...)
```

- [ ] **Step 2: Bind admin pages to API data**

```ts
await api.getBanners()
```

- [ ] **Step 3: Verify CRUD flows**

Run: backend tests and frontend smoke check
Expected: list/create/update/delete routes work

- [ ] **Step 4: Commit**

```bash
git add server admin-web
git commit -m "feat: connect admin platform modules"
```
