# Week 8

## 字节

### 10字节

```
[root@test--01 ~]$ redis-benchmark  -d 10 -r 10000 -n 50000 -t get,set --csv
"test","rps","avg_latency_ms","min_latency_ms","p50_latency_ms","p95_latency_ms","p99_latency_ms","max_latency_ms"
"SET","64020.48","0.442","0.120","0.423","0.679","0.951","3.223"
"GET","56242.97","0.533","0.104","0.455","1.023","1.319","3.831"
```

### 20字节

```
[root@test--01 ~]$ redis-benchmark  -d 20 -r 10000 -n 50000 -t get,set --csv
"test","rps","avg_latency_ms","min_latency_ms","p50_latency_ms","p95_latency_ms","p99_latency_ms","max_latency_ms"
"SET","71942.45","0.372","0.096","0.343","0.519","0.671","4.551"
"GET","64350.06","0.410","0.096","0.391","0.575","0.783","1.167"
```

### 50字节

```
[root@test--01 ~]$ redis-benchmark  -d 50 -r 10000 -n 50000 -t get,set --csv
"test","rps","avg_latency_ms","min_latency_ms","p50_latency_ms","p95_latency_ms","p99_latency_ms","max_latency_ms"
"SET","74074.07","0.374","0.104","0.351","0.575","0.735","1.079"
"GET","60313.63","0.438","0.128","0.439","0.623","0.743","1.279"
```

### 100字节

```
[root@test--01 ~]$ redis-benchmark  -d 100 -r 10000 -n 50000 -t get,set --csv
"test","rps","avg_latency_ms","min_latency_ms","p50_latency_ms","p95_latency_ms","p99_latency_ms","max_latency_ms"
"SET","75987.84","0.359","0.080","0.335","0.519","0.783","1.567"
"GET","74850.30","0.367","0.112","0.343","0.575","0.799","1.583"
```

### 1k字节

```
[root@test--01 ~]$ redis-benchmark  -d 1000 -r 10000 -n 50000 -t get,set --csv
"test","rps","avg_latency_ms","min_latency_ms","p50_latency_ms","p95_latency_ms","p99_latency_ms","max_latency_ms"
"SET","73964.50","0.361","0.104","0.343","0.479","0.631","1.183"
"GET","69541.03","0.387","0.104","0.359","0.567","0.775","1.087"
```

### 5k字节

```
[root@test--01 ~]$ redis-benchmark  -d 5000 -r 10000 -n 50000 -t get,set --csv
"test","rps","avg_latency_ms","min_latency_ms","p50_latency_ms","p95_latency_ms","p99_latency_ms","max_latency_ms"
"SET","58411.21","0.470","0.112","0.463","0.663","0.903","1.415"
"GET","59808.61","0.448","0.152","0.415","0.631","0.815","1.383"
```

## 插入kv

### 插入1w条，value长度为2字节

```shell
127.0.0.1:6379> info memory
# 插入前
# Memory
used_memory:27054864
used_memory_human:25.80M
used_memory_rss:18706432
used_memory_rss_human:17.84M
used_memory_peak:82368672
used_memory_peak_human:78.55M
used_memory_peak_perc:32.85%
used_memory_overhead:21632736
used_memory_startup:531544
used_memory_dataset:5422128
used_memory_dataset_perc:20.44%
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
# 插入后
127.0.0.1:6379> info memory
# Memory
used_memory:27371576
used_memory_human:26.10M
used_memory_rss:19521536
used_memory_rss_human:18.62M
used_memory_peak:82368672
used_memory_peak_human:78.55M
used_memory_peak_perc:33.23%
used_memory_overhead:22163808
used_memory_startup:531544
used_memory_dataset:5207768
used_memory_dataset_perc:19.40%
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
```

### 插入1w条，value长度为99字节

```shell
127.0.0.1:6379> info memory
# 插入前
# Memory
used_memory:25629760
used_memory_human:24.44M
used_memory_rss:19492864
used_memory_rss_human:18.59M
used_memory_peak:82368672
used_memory_peak_human:78.55M
used_memory_peak_perc:31.12%
used_memory_overhead:21763808
used_memory_startup:531544
used_memory_dataset:3865952
used_memory_dataset_perc:15.40%
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
127.0.0.1:6379> info memory
# 插入后
# Memory
used_memory:29680976
used_memory_human:28.31M
used_memory_rss:20656128
used_memory_rss_human:19.70M
used_memory_peak:82368672
used_memory_peak_human:78.55M
used_memory_peak_perc:36.03%
used_memory_overhead:22163744
used_memory_startup:531544
used_memory_dataset:7517232
used_memory_dataset_perc:25.79%
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B

```

### 测试结论

1. Key总数不变时，value越大，平均占用内存空间越多。