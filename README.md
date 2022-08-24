# Geoloc-API

## Usage

### Start Server
```bash
java -jar [-p <port>] Geoloc-API.jar <path to ip-locations.txt>
```

### Request
```bash
curl -X POST -d "<IP>" http://<serverIP>:<port>/country
```
