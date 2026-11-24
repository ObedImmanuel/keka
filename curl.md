```bash
curl -s -X POST localhost:4614/alert \
  -d \"{
 \"serviceId\":  \"svc-payments\",
 "severity":   \"CRITICAL\",          	
 \"title\":  	\"Payment gateway timeout\",
 \"dedupKey\":   \"payments-latency-001\"
 }\"
```
