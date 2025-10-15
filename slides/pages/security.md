---
layout: section
---

# Security

---

# Protecting supply chain

- `go.sum` &amp; package cache
- Vendoring
- CodeQL
- `govulncheck`

---

# Static Analysis (SAST)

- `gosec` - code security scanner
- `grype` - vulnerability scanner for containers
- CI/CD integration - https://github.com/golangci/golangci-lint

---

# Dynamic Analysis (DAST)

- OWASP ZAP - web application security scanner

---

# Docker build

```docker {|1-10|11-16|1|5-6|8-10|11|13|15}
FROM golang:1.23 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -ldflags="-s -w" -o /app/api .

FROM ubuntu:jammy AS runtime-stage

COPY --from=build-stage /app/api /usr/local/bin/api

ENTRYPOINT [ "/usr/local/bin/api" ]
```
