# github.com/a-h/go-workshop-102

## Topics

#### Monitoring / metrics with Prometheus

Users are reporting that this program is slow. We need to find out why.

Use Prometheus to monitor the program and find out what is slow.

#### Benchmarking with built-in benchmarking tools

Two developers are arguing whether it's faster to use a map or a slice for a particular use case. Benchmark the two implementations to find out.

- Example program: Searches for a list of films by title, does some string processing (e.g. removing punctuation, and converting to lowercase), and then returns the results.

#### Building Docker containers

3 people on the team have different opinions on how to build Docker containers.

One likes `ko` because it's simple, another uses `Dockerfile` because it's standard, and the third uses `Nix` because it provides reproducible builds.

Let's build the same program using all three methods and compare the results.

- Image size
- Build time
- Ease of use / understanding
- Vulnerability assessment with grype
- Reproducibility

- Example program: Hello World style app, just bundled a few different ways.

### Secure coding

The security team are concerned about potential vulnerabilities in your Go apps.

The project has been set up with a number of security tools to help you identify and fix potential issues.

- Supply chain analysis (go.sum, go.mod / grype, govulncheck)
- Secure coding practices (Static analysis, gosec, go vet, golangci-lint)
- Fuzzing (built in)
- Dynamic security testing (OWASP ZAP)

- Example program: A simple web server where you upload a CSV file, and it uses a SQLite DB to store it (so we can demo SQL injection).

#### Profiling

The program is running, but it's slow. The senior engineer has updated the program to include profiling, and has given some instructions on how to access CPU and memory profiles.

Identify the slow parts of the program and suggest improvements.

- Example program: A program that loads a HLS stream (e.g. https://stream-akamai.castr.com/5b9352dbda7b8c769937e459/live_2361c920455111ea85db6911fe397b9e/index.fmp4.m3u8) that creates a thumbnail on disk every 5 seconds.

## Tasks

### gomod2nix-update

```bash
gomod2nix
```

### build

```bash
nix build
```

### run

```bash
nix run
```

### develop

```bash
nix develop
```

### docker-build

```bash
nix build .#docker-image
```

### docker-load

Once you've built the image, you can load it into a local Docker daemon with `docker load`.

```bash
docker load < result
```

### docker-run

```bash
docker run -p 8080:8080 app:latest
```

### slides

dir: workshop-slides
interactive: true

```bash
npm run dev
```
