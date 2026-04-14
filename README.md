# 📧 Email Campaign Tool (Golang + Concurrency + Docker)

A **production-ready email campaign sender** built using:

* 🟦 **Golang**
* ⚡ **Goroutines & Worker Pool**
* 📩 **SMTP**
* 🐳 **Docker (Mailpit for testing)**

This project demonstrates real-world concepts like **concurrency, channels, worker pools, template rendering, bulk email sending, and SMTP integration**.

---

## 🚀 Features

### 🔹 Backend (Golang)

* ✅ Send bulk emails concurrently
* ✅ Worker Pool implementation
* ✅ Goroutines + Channels
* ✅ WaitGroup synchronization
* ✅ Personalized email templates
* ✅ CSV-based recipient import
* ✅ SMTP integration
* ✅ Scalable architecture

### 🔹 Email Testing

* 📩 Local SMTP testing using **Mailpit**
* 🌐 Web UI to preview sent emails

### 🔹 DevOps

* 🐳 Run Mailpit using Docker

---

## 🏗️ Project Structure

```bash
Email-Campaign-Tool/
│
├── Backend/
│   ├── main.go              # Entry point
│   ├── consumer.go          # Email worker logic
│   ├── producer.go          # Load recipients from CSV
│   ├── email.tmpl          # Email template
│   ├── recipients.csv      # Contact list
│   ├── go.mod
│   └── info.md             # Mailpit setup
│
├── .gitignore
└── README.md

```

⚙️ Tech Stack
Backend
Golang
Goroutines
Channels
WaitGroup
net/smtp
text/template
Email Testing
Mailpit
DevOps
Docker


🔧 Setup & Installation
1️⃣ Clone Repo
git clone https://github.com/your-username/Email-Campaign-Tool.git
cd Email-Campaign-Tool

▶️ Run Project Locally
1️⃣ Start Mailpit
  docker run -d \
  --name=mailpit \
  --restart unless-stopped \
  -p 8025:8025 \
  -p 1025:1025 \
  axllent/mailpit

  2️⃣ Run Golang App
      cd Backend
     go run .

---

## 👨‍💻 Author

**Tushar Choudhary**

* 🏆 SIH 2023 Winner
* 💼 Backend & Blockchain Developer
* ⚡ DevOps Enthusiast

---

## ⭐ Support

If you like this project:

👉 Star ⭐ the repo
👉 Fork & contribute

---   provide me readme like this 
