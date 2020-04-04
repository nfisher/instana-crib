# Instana API Cribsheet

## Requirements

 * Go 1.14+.
 * Make

## Building

`make`

## Execution

```
export INSTANA_URL={URL}
export INSTANA_TOKEN={TOKEN}

./infraq -query='entity.zone:k8s-demo' -plugin=host -metric=cpu.system -window=300000
./infraq -query='entity.zone:k8s-demo' -plugin=kubernetesPod -metric=cpuRequests -window=300000
```

## Relevant API URLs

* `/api/infrastructure-monitoring/catalog/plugins` - list  plugins in the system.
* `/api/infrastructure-monitoring/catalog/search` - list search fields.
* `/api/infrastructure-monitoring/catalog/metrics/{plugin}` - list available metrics.

## Metrics of Interest

### Host

* `cpu.sys`
* `cpu.user`
* `cpu.wait`
* `load.1m`
* `memory.buffers`
* `memory.cached`
* `memory.free`
* `memory.used`

# Namespace

* `alloc_pods_percentage`
* `cap_limits_cpu`
* `cap_pods`
* `cap_limits_memory`
* `cap_requests_cpu`
* `cap_requests_memory`
* `limit_cpu_percentage`
* `limit_mem_percentage`
* `required_cpu_percentage`
* `required_mem_percentage`
* `used_limits_cpu`
* `used_limits_memory`
* `used_pods`
* `used_pods_percentage`
* `used_requests_cpu`
* `used_requests_memory`
