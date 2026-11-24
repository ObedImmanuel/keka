```bash
curl -s -X POST localhost:4614/alert \
  -d "{
 \"serviceId":  "svc-payments",
 "severity":   "CRITICAL",          	// INFO | WARNING | CRITICAL
 "title":  	"Payment gateway timeout",
 "dedupKey":   "payments-latency-001"
 }"
```

```bash
curl -s -X POST localhost:4614/hello \
  -d "{\"userId\":1,\"idempotencyKey\":\"abc\"}"
```