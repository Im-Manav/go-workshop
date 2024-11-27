# github.com/a-h/go-workshop-102

This workshop is separated into the following sections:

- `cmd/app` 
    - hello world example handler
- `./benchmarking`
    - benchmark testing exercise
- `./fuzzing`
    - fuzz testing example
- `./security`
    - Unsecure API with security tooling exercise
    - Container vulnerability scanning exercise
    - Testing handlers exercise
- `./monitoring`
    - example app with monitoring exercise

Each section has it's own `README.md` to explain how to run or interact with the exercise along with the next steps to complete it.

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
