# Geoloc-API

## Usage

### Start Server
```bash
java -jar Geoloc-API.jar [-p <port>] <path to ip-locations.txt>
```

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
- `country`
- `coords`
