# Monitoring

## Getting started!
- docker compose up
- go run main.go


## Links
- localhost:8080/hello - This is your basic 200 route
- localhost:8080/metrics - Where Promethus is displayed
- localhost:8080/err - 404
- localhost:8080/internal-err - 500
- http://localhost:9090/query - prometheus
- http://localhost:3001/ - grafana

## Getting a dashboard
    - dashboards
    - create dashboard
    - add visualisation
    - promethus
    - set up a metric from dropdown (uptimeTotal)
    - run query
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
