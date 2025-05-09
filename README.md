# 🕔 5minpastebin.com

**A fast, ephemeral Pastebin — no configs, no accounts, no nonsense (or is it?)**

Paste your code or text and get a link that self-destructs in 5 minutes. Designed for simplicity, speed, and privacy. Ideal for quick sharing without worrying about cleanup or storage.

![Screenshot](https://your-screenshot-link-if-any.png)

---

## 🚀 Features

- ⏳ Pasted content expires after 5 minutes.
- 🧼 No signups, no settings, no clutter.
- 🔗 Shareable, one-time-use paste links.
- 🌐 Clean and minimal web UI.
- ⚡ Built with Go, HTMX, Templ, and Tailwind CSS.

---

## 🛠️ Getting Started

These instructions will help you run `5minpastebin.com` locally for development and testing.

### 📦 Prerequisites

- Go 1.24+
- Docker (for local redis)
- Make

---

## 📂 Clone the Repo

```bash
git clone https://github.com/YOUR_USERNAME/5minpastebin.com.git
cd 5minpastebin.com
```

## 🧰 Makefile Commands

### Build, Run, Test

| Command            | Description                          |
|--------------------|--------------------------------------|
| `make all`         | Run tests and build the project      |
| `make build`       | Build the application binary         |
| `make run`         | Run the server locally               |
| `make watch`       | Watch for file changes and reload    |
| `make test`        | Run unit tests                       |
| `make itest`       | Run integration tests                |
| `make clean`       | Clean up compiled binary             |

### Docker for DB

| Command              | Description                      |
|----------------------|----------------------------------|
| `make docker-run`    | Start DB container (PostgreSQL)  |
| `make docker-down`   | Stop DB container                |

---

## 🧪 Example

1. Paste any code/text in the home page.
2. Get a unique link like:
   `https://5minpastebin.com/a1b2c3`
3. That paste will auto-expire in 5 minutes.

---

## 🧑‍💻 Tech Stack

- **Go** – Backend
- **HTMX** – Dynamic frontend
- **Templ** – HTML templates with Go
- **Tailwind CSS** – Styles
- **Redis** – Storage

---

## 🔒 Privacy

Your data is never stored permanently. All pastes are deleted after 5 minutes — automatically. There is no logging, no cookies, and no tracking.

---

## 🌐 Live Demo

➡️ [https://5minpastebin.com](https://5minpastebin.com) *(Coming Soon)*

---

## 🤝 Contributing

Pull requests are welcome! For major changes, open an issue first to discuss what you’d like to change.

Please make sure to update tests as appropriate.

---

## 📜 License

This project is licensed under a custom **Non-Commercial License**.

You are free to use, modify, and deploy it for personal or educational use.
**Commercial use (including offering it as a hosted service) is strictly prohibited.**

For commercial licenses, contact the author.

---

## ⭐️ Show your support

If you like this project, consider giving it a ⭐️ on GitHub!

---

## 🔗 Follow the Dev

Made with ❤️ by [@dhruvsachdev](https://github.com/Dhruv-Sachdev1313)
