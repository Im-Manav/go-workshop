# Monitoring

## API endpoints

Use these to access endpoints

- localhost:8080/hello - This is your basic 200 route
- localhost:8080/metrics - Where Prometheus metrics are displayed plaintext, scraped by
- localhost:8080/err - An endpoint that produces a 404 error
- localhost:8080/internal-err - An endpoint that produces a 500 error
- http://localhost:9090/query - Prometheus
- http://localhost:3001/ - Grafana dashboard

The app has a few endpoints that generate metrics

`/hello` - Just returns a hello back and counts as a succesful  Prometheus:
To query prometheus go to `http://localhost:9090/query` on your browser.

Grafana:

Go to `http://localhost:3001/` on your browser and logging with username `admin` and password `grafana`.

## Dashboard guide

To see the Grafana dashboard for uptime use the following
- Navigate to `http://localhost:3001/`, the Grafana dashboard
- You'll be presented with a login screen.
    - Username: `admin` 
    - Password: `grafana`
- You should be on the home screen now. Click dashboards on the left or navigate to `http://localhost:3001/dashboards`
- Click the blue `New` button in the top right, and click `New Dashboard`
- `Start your new dashboard by adding a visualization` should be on your screen now. Click the big blue `+ Add visualization` 
- Select `Prometheus` in the data source pop up
- Your dashboard should now be created, make sure you're on the query tab at the bottom, click add query and select `uptimeTotal` from the metric dropdown (If you're adding your own metric, then change this)
- Finally, click `Run Queries` button, you should then see a graph of the uptime show up

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
