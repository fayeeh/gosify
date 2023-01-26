Gosify is a Redis-like(?) database.

Key and value are strings. (It was easy for me)

Basic usage:
```
SET a b 5s
GET a
RM a
```

With the 3rd parameter we specify how many seconds it will be deleted. (It can be minutes or hours)
For time durations: (Click)[https://pkg.go.dev/time#ParseDuration]