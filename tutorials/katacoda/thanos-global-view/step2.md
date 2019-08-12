# Step 2 - Starting Prometheus

We'll now start two containers representing our two different Prometheus instances.

Please note the extra flags we're passing to Prometheus:

* `--web.enable-lifecycle` allows Thanos Sidecar to...
* `--web.enable-admin-api` allows Thanos Sidecar to...
* `--storage.tsdb.min-block-duration=2h` and `--storage.tsdb.max-block-duration=2h` are the recommended values to use with Thanos.

## US1

```
docker run -d --net=host \
    -v /root/prometheus_us1.yml:/etc/prometheus/prometheus.yml \
    --name prometheus-us1 \
    prom/prometheus \
    --config.file=/etc/prometheus/prometheus.yml \
    --storage.tsdb.path=/prometheus \
    --web.listen-address=:9090 \
    --web.enable-lifecycle \
    --web.enable-admin-api \
    --storage.tsdb.min-block-duration=2h \
    --storage.tsdb.max-block-duration=2h \
```{{execute}}

Once started, the [dashboard](https://[[HOST_SUBDOMAIN]]-9090-[[KATACODA_HOST]].environments.katacoda.com/) is viewable on port [9090](https://[[HOST_SUBDOMAIN]]-9090-[[KATACODA_HOST]].environments.katacoda.com/).

## EU1

```
docker run -d --net=host \
    -v /root/prometheus_eu1.yml:/etc/prometheus/prometheus.yml \
    --name prometheus-eu1 \
    prom/prometheus \
    --config.file=/etc/prometheus/prometheus.yml \
    --storage.tsdb.path=/prometheus \
    --web.listen-address=:9091 \
    --web.enable-lifecycle \
    --web.enable-admin-api \
    --storage.tsdb.min-block-duration=2h\
    --storage.tsdb.max-block-duration=2h
```{{execute}}

Once started, the [dashboard](https://[[HOST_SUBDOMAIN]]-9091-[[KATACODA_HOST]].environments.katacoda.com/) is viewable on port [9091](https://[[HOST_SUBDOMAIN]]-9091-[[KATACODA_HOST]].environments.katacoda.com/).