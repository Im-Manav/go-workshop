---
layout: center
---

# Security

---

# How to keep our Go apps safe

## Supply chain

- Go modules, lock &amp; cache
- Vendoring
- CodeQL

## SAST

- govulncheck - tool written by the Go security team
- gosec - code security scanner
- grype - vulnerability scanner for containers
- CI/CD integration - https://github.com/golangci/golangci-lint

## DAST

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

---

## Security in Action!

In the `security` folder you will find the `README.md` with instructions on how to run the app and interact with it.

- Detect the security vulnerability with `gosec`.
- Fix the security vulnerability.
- Build the container, and detect issues with `grype`.
- Fix the Docker security issue.
