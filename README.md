# Go Observability: OpenTracing with Zap Logger

Example for context propagation with OpenTelemetry trace and zap logging for
logging in distributed applications.

structured output

```
{"level":"info","ts":1648373086.577307,"caller":"opentelemetry-logger/main.go:20","msg":"start test without traces"}
{"level":"info","ts":1648373086.577437,"caller":"opentelemetry-logger/main.go:34","msg":"span test 1 with trace","trace-id":"505fd06ccf32b4904fc60b916623283a","span-id":"d0c7aaed0d60a5ab"}
{"level":"info","ts":1648373086.577467,"caller":"opentelemetry-logger/main.go:35","msg":"span test 1 repeats trace and span id","trace-id":"505fd06ccf32b4904fc60b916623283a","span-id":"d0c7aaed0d60a5ab"}
{"level":"info","ts":1648373086.577496,"caller":"opentelemetry-logger/main.go:43","msg":"span test 2 with new trace","trace-id":"9c6e95aa4d6254183108ddb4d62dfa8b","span-id":"65c6fb4b4489f025"}
{"level":"info","ts":1648373086.577538,"caller":"opentelemetry-logger/main.go:52","msg":"span test 3 repeates trace id from 2 with new span id","trace-id":"9c6e95aa4d6254183108ddb4d62dfa8b","span-id":"2a7190e5e0ca8f16"}
{"level":"info","ts":1648373086.577607,"caller":"opentelemetry-logger/main.go:26","msg":"exit test without traces"}
```

unstructured output

```
2022-03-27T17:25:22.191+0800    INFO    opentelemetry-logger/main.go:20 start test without traces
2022-03-27T17:25:22.191+0800    INFO    opentelemetry-logger/main.go:34 span test 1 with trace  {"trace-id": "3282793f68a156aabf48335d7ba1f468", "span-id": "469843b9d0d1275a"}
2022-03-27T17:25:22.191+0800    INFO    opentelemetry-logger/main.go:35 span test 1 repeats trace and span id   {"trace-id": "3282793f68a156aabf48335d7ba1f468", "span-id": "469843b9d0d1275a"}
2022-03-27T17:25:22.191+0800    INFO    opentelemetry-logger/main.go:43 span test 2 with new trace      {"trace-id": "05525adea5a16be4be049bd336d604b3", "span-id": "339d84828277142f"}
2022-03-27T17:25:22.191+0800    INFO    opentelemetry-logger/main.go:52 span test 3 repeates trace id from 2 with new span id   {"trace-id": "05525adea5a16be4be049bd336d604b3", "span-id": "3cce817c9c297453"}
2022-03-27T17:25:22.191+0800    INFO    opentelemetry-logger/main.go:26 exit test without traces
```
