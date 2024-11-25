# Security

## Tasks

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
docker build -t go-workshop-102:security .
```

### run-docker-container

```bash
docker run -p 8005:8005 go-workshop-102:security
```

### scan-docker-container

```bash
grype go-workshop-102:security
```

## Next steps

- Detect the security vulnerability with `gosec`.
- Fix the security vulnerability.
- Use `grype` to check the Docker container for vulnerabilities.
