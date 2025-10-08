---
theme: dracula
title: Go Workshop
---

# Go Workshop

---

# People

## Me (Adrian Hesketh)

- Using Go since 2016
- "Famous" for the `templ` package
  - `https://github.com/a-h/templ`
- Built and operated production systems in Go
  - UK's 3rd largest online pharmacy
  - Vehicle insurance company

## Contributors

- Arman
- Mia

---

# Setup

- Clone https://github.com/a-h/go-workshop
- Install Go

The 2xx monitoring and security sessions use some container images, so it's a good idea to pull them down now if you can, since the downloads might be slow during the session.

- `docker pull ubuntu:jammy`
- `docker pull golang:1.23`
- `docker pull prom/prometheus`
- `docker pull grafana/grafana`

---

# Workshop structure

- Divided into two sections
  - 100 sessions: Go basics
  - 200 sessions: Advanced Go

---

# 100 sessions: Go basics

- Hello World
- CLI flags
- Testing and table-driven tests
- Web servers
- Type system
- Concurrency features of Go

---

# 200 sessions: Advanced Go

- Dependency management with Go modules
- Error handling
- Benchmarking
- Security features of Go and the wider ecosystem
- Monitoring with Prometheus and Grafana

---
src: ./pages/setup.md
---

---
src: ./pages/testing.md
---

---
src: ./pages/benchmark.md
---

---
src: ./pages/security.md
---

---
src: ./pages/monitoring.md
---
