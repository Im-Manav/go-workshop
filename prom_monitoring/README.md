# Promethus

As a reminder, the goal here is to familiarise yourself with Promethus and Grafana as it intergates with Go, and to create a new

## Getting started!
- `cd prom_monitoring`
    - get in the right folder
- `docker compose up`
    - This runs the grafana and promethus docker containers. It may take a few moments after the next command for promethus to scrape your metrics, and for grapaha to pick these up
- `go run main.go`
    - run the web server that generates metrics, giving promethus something to pick up

## Links
Use these to to 

- localhost:8080/hello - This is your basic 200 route
- localhost:8080/metrics - Where Promethus metrics are displayed plaintext, scraped by
- localhost:8080/err - An endpoint that produces a 404 error
- localhost:8080/internal-err - An endpoint that produces a 500 error
- http://localhost:9090/query - Promethus
- http://localhost:3001/ - Grafana dashboard

## Getting a dashboard

To see the Grafana dashboard for uptime use the following
- Navigate to `http://localhost:3001/`, the Grafana dashboard
- You'll be presented with a login screen.
    - Username: `admin` 
    - Password: `grafana`
- You should be on the home screen now. Click dashboards on the left or navigate to `http://localhost:3001/dashboards`
- Click the blue `New` button in the top right, and click `New Dashboard`
- `Start your new dashboard by adding a visualization` should be on your screen now. Click the big blue `+ Add visualization` 
- Select `Prometheus` in the data source pop up
- Your dashboard should now be created, make sure you're on the query tab at the bottom, click add query and select `uptimeTotal` from the metric dropdown
- Finally, click `Run Queries` button, you should then see a graph of the uptime show up
