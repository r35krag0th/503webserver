# Only 503s!

A super pointless example to test how different things deal with the 503
status code.

## Build

```shell
go build
```

## Run

```shell
./503webserver
```

## Test

By default it listens on `:3333`.

```shell
curl http://127.0.0.1:3333/health
# exits 0
```

```shell
curl --fail http://127.0.0.1:3333/health
# exits 22
```

```shell
wget -O- http://127.0.0.1:3333/health
# exits 8
# ```
