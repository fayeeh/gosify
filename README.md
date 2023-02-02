Gosify is a Redis-like(?) database.

Key and value are strings.

Basic usage:
```
SET a b 5s
GET a
RM a
```

With the 3rd parameter we specify how many seconds it will be deleted. (It can be minutes or hours)

For time durations: [Click](https://pkg.go.dev/time#ParseDuration)

Note: If the 3rd parameter is empty, the data will not be deleted.

# Commands
```
SET(set): set <key> <value> <duration> # set a b 5h2s
GET(get): get <key>
RM(rm, remove, delete, DELETE): rm <key>
```
