# Security

Go has a number of useful security features that you can use to secure your applications. In this workshop, we will explore some of these features.

## govulncheck

`govulncheck` is a tool that checks your Go code for known supply chain vulnerabilities. It uses the [Go vulnerability database](https://pkg.go.dev/golang.org/x/vuln) to check for vulnerabilities in your code and its dependencies.

It's similar to `npm audit` for Node.js, except with an important difference. `npm audit` tells you if your dependencies have known vulnerabilities. `govulncheck` goes a step further and tells you if your code is actually using the vulnerable code.

To install `govulncheck`, run `go install golang.org/x/vuln/cmd/govulncheck@latest`.

At the moment, `govulncheck` is an experimental tool, which will change over time. When it stabilises, it will be included in the Go toolchain where the compatibility promise applies.

## gosec

`gosec` is a static analysis tool that checks your Go code for security issues. It looks for common security issues such as SQL injection, cross-site scripting (XSS), and hardcoded credentials.

To install `gosec`, run `go install github.com/securego/gosec/v2/cmd/gosec@latest`.

## grype

`grype` is a vulnerability scanner for container images and filesystems. It scans your Docker images for known vulnerabilities in the packages they contain.

Despite being a security tool, it recommends installing it via a shell script, complete with `sudo`.

To install `grype`, check the instructions at https://github.com/anchore/grype?tab=readme-ov-file#installation

## Instructions

The tasks below guide you through installing and using these tools to check the application in this repository for security issues. Copy and paste the commands into your terminal to run them, or use the `xc` command if you have it installed see github.com/joerdav/xc for more details.

## Tasks

### run-govulncheck

Run the govulncheck tool to see if there are any known vulnerabilities in the supply chain.

```bash
govulncheck ./...
```

### run-gosec

Run the gosec tool.

```bash
gosec ./...
```

### export-path

If you get a "command not found" error, Go installs binaries in the $GOPATH/bin.

```bash
export PATH=$PATH:$GOPATH/bin
```

### run

Run the application. It will serve on port 8005.

```bash
go run .
```

### create-customer

interactive: true

First, let's create a customer.

```bash
curl --verbose -d '{"Id": 1, "name": "John", "surname": "Doe", "company": "JDoe LTD"}' http://localhost:8005/customer
```

### get-customer

interactive: true

We can now get the customer we just created.

```bash
curl http://localhost:8005/customer/1
```

### get-customer-hacker

interactive: true

But a hacker can craft a URL to delete all customers. This is the URL encoded version of `1;delete from customer;`.

```bash
curl 'http://localhost:8005/customer/1%3Bdelete%20from%20customers%3B'
```

### create-docker-container

```bash
docker build -t go-workshop-200:security .
```

### run-docker-container

```bash
docker run -p 8005:8005 go-workshop-200:security
```

### scan-docker-container

```bash
grype go-workshop-200:security
```

## Next steps

- Detect security vulnerabilities with `govulncheck`.
- Detect additional security vulnerabilities with `gosec`.
- Use `grype` to check the Docker container for vulnerabilities and fix those!

There's no "solution" for this workshop, as the tools will find different issues over time. The important thing is to get familiar with the tools and how to use them.
