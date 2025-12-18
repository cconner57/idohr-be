# IDOHR Backend (Pet Adoption Platform)

This repository contains the backend API for the **I Dream of Home Rescue (IDOHR)** platform. It is a dedicated server application built with **Go (Golang)** that manages pet data, adoption workflows, and shelter operations.

The application is designed to run on a **Raspberry Pi 5** (Production) and macOS (Development), connecting to a **PostgreSQL** database.

---

## 🏗 Architecture

The project follows a **Modular Monolith** architecture to support scaling from a simple intake system to a full SaaS-ready ERP.

- **Language:** Go 1.24+
- **Database:** PostgreSQL 16
- **Server:** Standard Library `net/http` (No heavy frameworks)
- **Deployment:** Raspberry Pi 5 (Headless) via SSH

### Folder Structure

```text
idohr-be/
├── cmd/
│   ├── api/            # REST API Server (Main Application)
│   ├── seeder/         # CSV Data Importer
│   └── worker/         # Background Tasks (Email, Image Processing)
├── internal/
│   ├── data/           # Database Models (SQL Queries)
│   ├── models/         # Shared Structs (JSON Shapes)
│   ├── services/       # Business Logic (Stripe, PDF, IoT)
│   ├── middleware/     # Auth, Rate Limiting, Logging
│   └── validator/      # Form Validation Helpers
├── migrations/         # SQL Database Schema Changes
└── uploads/            # Local storage for pet photos
```
