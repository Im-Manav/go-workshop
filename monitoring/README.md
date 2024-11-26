# Monitoring

## API endpoints

The app has a few endpoints that generate metrics

`/hello` - Just returns a hello back and counts as a succesful  Prometheus:
To query prometheus go to `http://localhost:9090/query` on your browser.

Grafana:
Go to `http://localhost:3001/` on your browser and logging with username `admin` and password `grafana`.

## Tasks

### docker-up

Start the grafana and prometheus docker containers. Grafana port 3001 and Prometheus port 9090

```bash
docker compose up
```

### run-app

Start the go app. It will serve on port 8080 

```bash
go run .
```

### get-hello

Call to the `/hello` endpoint.

```bash
curl --verbose "http://localhost:8080/hello"
```

### get-metrics

Call to the `metrics` endpoint.

```bash
curl --verbose "http://localhost:8080/metrics"
```
### notfound-err

Call to the `/err` endpoint.

```bash
curl --verbose "http://localhost:8080/err"
```

### internal-err

Call to the `/internal-err` endpoint.

```bash
curl --verbose "http://localhost:8080/internal-err"
```

## Next steps
- Add another metric
- Query them in Prometheus
- Create a dashboard in grafana
