# Geoloc-API

## Usage

### Start Server
```bash
java -jar Geoloc-API.jar [-p <port>] <path to ip-locations.txt>
```

### Request
```bash
curl -X POST -d "<IP>" http://<serverIP>:<port>/country
```
