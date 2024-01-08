# Geoloc-API

## Usage

### Run (binary)

- download binary from releases or compile (see Makefile in src folder)
```bash
./Geoloc-API "path/to/ip-locations.txt"
```

### Run (Docker)

- `docker build . -t geolocapi`
- `docker run -d -p 9001:9001 geolocapi`

### ENV Variables

- `GEOAPI_RAM_OPT`: enables low memory mode if set to `"1"`
- `GEOAPI_PORT`: port HTTP server will listen to

### Request
#### CURL
```bash
curl -X POST -d "<IP>" http://<serverIP>:<port>/<endpoint>
```

#### PHP
```php
<?php

$opts = array('http' =>
    array(
        'method' => 'POST',
        'header' => 'Content-type: application/x-www-form-urlencoded',
        'content' => "<IP>"
    )
);
$context = stream_context_create($opts);
$result = file_get_contents('http://<serverIP>:<port>/<endpoint>', false, $context);
echo $result;

?>
```

### Endpoints
- `POST /country`
- `POST /coords`
- `POST /all/json` example respone:
```json
{
  "Range": {
    "StartIP": 3758088960,
    "EndIP": 3758089215
  },
  "Country": "Indonesia",
  "Location": {
    "Latitude": -0.22019,
    "Longitude": 100.63078
  }
}
```