# Mini monitor

This is a small a standalone status page for multiple service checks.
(Like Nagios just much dumber.)

## Build

Simply run the following command:

```
$ go build src/cmd/monitor/monitor.go
```

## Configuration

Put a json file with a similar content like:
```
{
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


#### Docker

Basic docker container checker using the `docker ps` command.

Type definition: `docker`

Possible keys:
