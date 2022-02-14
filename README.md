# beacon-server
very simple server for browser sendBeacon 

## Local build

```bash
docker build . -t beacon-server
docker run -d --rm -p 8090:8090 --name beacon-server  beacon-server

```

## Testing

Watch logs enhanced and colored by `jq`

```bash

docker logs -f beacon-server | jq &

```

### Test with curl

```bash

curl -X POST localhost:8090/response -H 'Content-Type: application/json'  -d '{"file":"xyz.js:123","condition":"red","version": "0.8.0"}'

```

### Basic error handling

```bash

curl -X POST localhost:8090/response -H 'Content-Type: application/json'  -d '{"file":"xyz.js:123","condition":"red","'
curl -X POST localhost:8090/response -H 'Content-Type: application/json'  -d ''

```

### Output

The command `docker logs beacon-server` show the JSON logging

```
{"time": "2022-02-14 08:27:06.320446191 +0000 UTC","Content-Type": "application/json","Content-Length": "36","User-Agent": "curl/7.68.0","Accept": "*/*","error:": "JSON parse error: unexpected end of JSON input {\"login\":\"my_login\",\"password\":\"my_p"}{"time": "2022-02-14 08:27:06.320446191 +0000 UTC","Content-Type": "application/json","Content-Length": "36","User-Agent": "curl/7.68.0","Accept": "*/*","error:": "JSON parse error: unexpected end of JSON input {\"login\":\"my_login\",\"password\":\"my_p"}
{"time": "2022-02-14 08:27:11.59989297 +0000 UTC","User-Agent": "curl/7.68.0","Accept": "*/*","Content-Type": "application/json","Content-Length": "45","login": "my_login","password": "my_password"}
{"time": "2022-02-14 08:28:15.921562051 +0000 UTC","User-Agent": "curl/7.68.0","Accept": "*/*","Content-Type": "application/json","Content-Length": "0","error": "no body supplied"}{"time": "2022-02-14 08:28:15.921562051 +0000 UTC","User-Agent": "curl/7.68.0","Accept": "*/*","Content-Type": "application/json","Content-Length": "0","error": "no body supplied"}
{"time": "2022-02-14 08:30:01.322720114 +0000 UTC","User-Agent": "curl/7.68.0","Accept": "*/*","Content-Type": "application/json","Content-Length": "0","error": "no body supplied"}{"time": "2022-02-14 08:30:01.322720114 +0000 UTC","User-Agent": "curl/7.68.0","Accept": "*/*","Content-Type": "application/json","Content-Length": "0","error": "no body supplied"}
{"time": "2022-02-14 08:30:11.187699296 +0000 UTC","User-Agent": "curl/7.68.0","Accept": "*/*","Content-Type": "application/json","Content-Length": "45","login": "my_login","password": "my_password"}
{"time": "2022-02-14 08:30:31.299761777 +0000 UTC","User-Agent": "curl/7.68.0","Accept": "*/*","Content-Length": "45","Content-Type": "application/x-www-form-urlencoded","login": "my_login","password": "my_password"}
{"time": "2022-02-14 08:30:41.060121739 +0000 UTC","User-Agent": "curl/7.68.0","Accept": "*/*","Content-Length": "6","Content-Type": "application/x-www-form-urlencoded","error:": "JSON parse error: invalid character 'a' looking for beginning of value adadad"}{"time": "2022-02-14 08:30:41.060121739 +0000 UTC","User-Agent": "curl/7.68.0","Accept": "*/*","Content-Length": "6","Content-Type": "application/x-www-form-urlencoded","error:": "JSON parse error: invalid character 'a' looking for beginning of value adadad"}
{"time": "2022-02-14 08:31:05.717856866 +0000 UTC","User-Agent": "curl/7.68.0","Accept": "*/*","Content-Length": "6","Content-Type": "application/x-www-form-urlencoded","error:": "JSON parse error: invalid character 'a' looking for beginning of value adadad"}{"time": "2022-02-14 08:31:05.717856866 +0000 UTC","User-Agent": "curl/7.68.0","Accept": "*/*","Content-Length": "6","Content-Type": "application/x-www-form-urlencoded","error:": "JSON parse error: invalid character 'a' looking for beginning of value adadad"}
{"time": "2022-02-14 08:40:00.780884108 +0000 UTC","User-Agent": "curl/7.68.0","Accept": "*/*","Content-Type": "application/json","Content-Length": "40","error:": "JSON parse error: unexpected end of JSON input {\"file\":\"xyz.js:123\",\"condition\":\"red\",\""}{"time": "2022-02-14 08:40:00.780884108 +0000 UTC","User-Agent": "curl/7.68.0","Accept": "*/*","Content-Type": "application/json","Content-Length": "40","error:": "JSON parse error: unexpected end of JSON input {\"file\":\"xyz.js:123\",\"condition\":\"red\",\""}
{"time": "2022-02-14 08:40:00.792837858 +0000 UTC","Accept": "*/*","Content-Type": "application/json","Content-Length": "0","User-Agent": "curl/7.68.0","error": "no body supplied"}{"time": "2022-02-14 08:40:00.792837858 +0000 UTC","Accept": "*/*","Content-Type": "application/json","Content-Length": "0","User-Agent": "curl/7.68.0","error": "no body supplied"}
{"time": "2022-02-14 08:40:02.989042811 +0000 UTC","Content-Length": "40","User-Agent": "curl/7.68.0","Accept": "*/*","Content-Type": "application/json","error:": "JSON parse error: unexpected end of JSON input {\"file\":\"xyz.js:123\",\"condition\":\"red\",\""}{"time": "2022-02-14 08:40:02.989042811 +0000 UTC","Content-Length": "40","User-Agent": "curl/7.68.0","Accept": "*/*","Content-Type": "application/json","error:": "JSON parse error: unexpected end of JSON input {\"file\":\"xyz.js:123\",\"condition\":\"red\",\""}
{"time": "2022-02-14 08:40:03.000862602 +0000 UTC","User-Agent": "curl/7.68.0","Accept": "*/*","Content-Type": "application/json","Content-Length": "0","error": "no body supplied"}{"time": "2022-02-14 08:40:03.000862602 +0000 UTC","User-Agent": "curl/7.68.0","Accept": "*/*","Content-Type": "application/json","Content-Length": "0","error": "no body supplied"}
{"time": "2022-02-14 08:40:03.732770516 +0000 UTC","User-Agent": "curl/7.68.0","Accept": "*/*","Content-Type": "application/json","Content-Length": "40","error:": "JSON parse error: unexpected end of JSON input {\"file\":\"xyz.js:123\",\"condition\":\"red\",\""}{"time": "2022-02-14 08:40:03.732770516 +0000 UTC","User-Agent": "curl/7.68.0","Accept": "*/*","Content-Type": "application/json","Content-Length": "40","error:": "JSON parse error: unexpected end of JSON input {\"file\":\"xyz.js:123\",\"condition\":\"red\",\""}
{"time": "2022-02-14 08:40:03.745281961 +0000 UTC","Content-Length": "0","User-Agent": "curl/7.68.0","Accept": "*/*","Content-Type": "application/json","error": "no body supplied"}{"time": "2022-02-14 08:40:03.745281961 +0000 UTC","Content-Length": "0","User-Agent": "curl/7.68.0","Accept": "*/*","Content-Type": "application/json","error": "no body supplied"}
{"time": "2022-02-14 08:40:19.914386005 +0000 UTC","Content-Length": "0","Content-Type": "application/x-www-form-urlencoded","User-Agent": "curl/7.68.0","Accept": "*/*","error": "no body supplied"}{"time": "2022-02-14 08:40:19.914386005 +0000 UTC","Content-Length": "0","Content-Type": "application/x-www-form-urlencoded","User-Agent": "curl/7.68.0","Accept": "*/*","error": "no body supplied"}
{"time": "2022-02-14 08:40:48.522727641 +0000 UTC","Content-Length": "0","Content-Type": "application/x-www-form-urlencoded","User-Agent": "curl/7.68.0","Accept": "*/*","error": "no body supplied"}{"time": "2022-02-14 08:40:48.522727641 +0000 UTC","Content-Length": "0","Content-Type": "application/x-www-form-urlencoded","User-Agent": "curl/7.68.0","Accept": "*/*","error": "no body supplied"}
{"time": "2022-02-14 08:41:06.082859306 +0000 UTC","User-Agent": "curl/7.68.0","Accept": "*/*","X-Asdadad": "adadad","Content-Length": "0","Content-Type": "application/x-www-form-urlencoded","error": "no body supplied"}
```

