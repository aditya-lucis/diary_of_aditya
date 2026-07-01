# Diary Of Aditya

![Diary Of Aditya Logo](assets/logo.png)

**IT Consultant Management System**  
*Integrated platform for CRM, Project Management, Timesheet, and Invoicing*

---

## ✨ About the Application

**Diary Of Aditya** is an internal system specially designed for IT consultants and enterprise teams. It integrates the entire business workflow — from daily activity logging (*diary*), project management, to automated invoicing.

Built to solve key challenges faced by IT consultants:
- Data fragmentation across platforms
- Billing leakage (unbilled working hours)
- Lack of project transparency for clients
- Difficulty in calculating project profitability

## 🎯 Key Features

### Core Diary
- **Kanban Board** with drag & drop
- **Timesheet Logging** + Start/Stop Timer
- **Daily Activity Log**
- **Issue & Bug Tracker** integrated with tasks

### Project & Client Management
- **Full CRM** (Leads → Active Clients)
- **Project & Milestone Management**
- **Work Breakdown Structure (WBS)**
- **Task Dependencies**

### Estimation & Finance
- **Project Complexity Matrix** (Architecture, Integration, Security)
- **Automated BoQ / Cost Estimation** + Risk Buffer
- **Automatic Invoicing** based on milestones or timesheets
- **Expense Tracking** per project
- **Dynamic Payment Status**

### Analytics & Dashboard
- **Burn Rate Analysis** (Actual vs Estimate)
- **Early Warning System**
- **Profitability Dashboard**
- **Real-time KPI Overview**

## 🛠️ Tech Stack

- **Backend**: Golang + Gin Framework (Clean Architecture)
- **Frontend**: Vue.js 3 + Pinia + TailwindCSS
- **Database**: PostgreSQL
- **Authentication**: JWT
- **API Documentation**: OpenAPI 3.0 (Swagger)
- **Deployment**: Docker Ready

## 🚀 Quick Start

```bash
# Clone repository
git clone https://github.com/aditya-lucis/diary-of-aditya.git
cd diary-of-aditya

# Backend
cd backend
cp .env.example .env
go mod tidy
go run main.go

# Frontend (new terminal)
cd frontend
npm install
npm run dev