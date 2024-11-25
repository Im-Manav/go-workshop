---
layout: center
---

# Monitoring

---

# What's Monitoring? 

Collecting live data about our program.

---

# Promethus and Graphana

TODO: introduce these
Promethus is an open source monitoring system
Graphana is an observability platform and allows for visualisations

# Promethus in Go

Promethus has an offical go package




// Should this be included somwhere else? Efectivly the conclusion

Red Method:
- Rate (the number of requests per second)
- Errors (the number of those requests that are failing)
- Duration (the amount of time those requests take)

Good for microservices, and records things that directly affect users
Good proxy to how happy your customers will be
Generally, you want to focus on business metrics rather than technical ones (cpu used etc)
Of course these can be useful in certain circumstances but shouldn't be your goal



