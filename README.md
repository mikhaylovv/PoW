build 
```bash
docker build -t server -f services/server/Dockerfile .
docker build -t client -f services/client/Dockerfile .
```

run 
```bash
docker run -p 8080:8080 server
docker run --network="host" client
```