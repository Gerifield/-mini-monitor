# Mini monitor [![Build Status](https://travis-ci.org/Gerifield/mini-monitor.svg?branch=master)](https://travis-ci.org/Gerifield/mini-monitor)

This is a small a standalone status page for multiple service checks.
(Like Nagios just much dumber.)

## Build

Simply run the following command:

```
$ go build src/cmd/monitor/monitor.go
```

## Run

```
$ monitor -listen :8080 -config config.json
```

By default the web UI will listen on `localhost:8080`.

You could open the following URLs in your browser:
- `/` - HTML web UI
- `/api` - JSON based status endpoint

Web UI example:
![web ui](misc/pictures/screenshot1.png)

JSON response structure:
```
{
   "checks":{
      "docker-alpine":false,
      "test1":true,
      "test3":true,
      "test4":false
   }
}
```


## Configuration

Put a json file with a similar content like:
```
{
  "checkTime": "5s",
  "configs": [
    {
      "name": "unique-config-name",
      "type": "debug",
      "config": {
        "testKey": "testVal"
      }
    },
    ...
  ]
}
```

Here the fields are the following:
 - `name` is the modules name you'd like to use, it should be unique in the config
 - `type` should be the checkers type, see the available ones below
 - `config` is the given modules configuration

### Available checkers

#### Debug

Obviously for debugging purpose.

Type definition: `debug`

Possible keys:
 - `loadFail` (bool) triggers an error on module load
 - `checkFail` (bool) tiggers and error at the check call


#### Docker

Basic docker container checker using the `docker ps` command.

Type definition: `docker`

Possible keys:
 - `id` (string) is the container id check
 - `nameRegex` (string) is the regular expression to match on a container's name
 - `imageRegex` (string) is the regular expression to match on an image's name
 - `debug` (bool) is for help the regular expression debugging

They are all optional and will match these fields to the `docker ps` output, but be careful if there are multiple matches it'll grab and check the first match only!
(Use the `id` if you'd like to guarantee to have a single match or use unique container names.)

#### HTTP

Type definition: `http`

Possible keys:
- `method` (string) is the HTTP method to use (auto converted to upper-case)
- `url` (string) is the HTTP URL to call
- `body` (string) is the HTTP request body
- `headers` (map[string]string) are the request headers in a key - value format
- `expectedCode` (int) is the HTTP response code for the successful check

The `method`, `url` and `expectedCode` are required.

Example config:
```
{
  "name": "test-http",
   "type": "http",
   "config": {
     "method": "get",
     "url": "http://192.168.0.100:3001/",
     "headers": {
       "Authorization": "Bearer testToken"
     },
     "expectedCode": 200
   }
 }
```