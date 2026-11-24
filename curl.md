```bash
curl -s -X POST localhost:4614/alert \
  -H 'Content-Type: application/json' \
  -d '{"serviceId":"svc-payments","severity":"CRITICAL","title":"Payment gateway timeout","dedupKey":"payments-latency-001"}'
```
