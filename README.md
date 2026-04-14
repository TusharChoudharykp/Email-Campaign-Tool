📧 Email Campaign Tool (Go + Concurrency + SMTP)
A production-ready email campaign sender built using:

🟦 Golang

⚡ Goroutines + Worker Pool

📩 SMTP

🐳 Mailpit (Local Email Testing)

This project demonstrates concurrency in Go, bulk email processing, template rendering, and scalable worker-based architecture for sending personalized campaigns.

🚀 Features
🔹 Core Functionality
✅ Send bulk emails concurrently

✅ Worker Pool Architecture

✅ Personalized email templates

✅ Read recipients from CSV file

✅ SMTP email delivery

✅ Fast and scalable processing

✅ Goroutines + Channels + WaitGroup

🔹 Email Testing
🐳 Mailpit for local SMTP testing

📬 View all sent emails in browser UI

🔹 Go Concepts Used
⚡ Goroutines

🔄 Channels

🤝 WaitGroup

📄 Templates

📂 File Handling

📧 net/smtp package

🏗️ Project Structure
Bash
Email-Campaign-Tool/
│
├── Backend/
│   ├── main.go            # Entry point
│   ├── consumer.go        # Email workers / consumers
│   ├── producer.go        # Load recipients from CSV
│   ├── email.tmpl         # Email template
│   ├── recipients.csv      # Recipient data
│   ├── info.md            # Mailpit setup
│   └── go.mod
│
├── .gitignore
└── README.md





👨‍💻 Author
Tushar Choudhary

🏆 SIH 2023 Winner

💼 Backend & Blockchain Developer

⚡ DevOps Enthusiast

⭐ Support
If you like this project:

👉 Star ⭐ the repo 👉 Fork & contribute
