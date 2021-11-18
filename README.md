# cmdbgo
cmdbgo 不积跬步，无以至千里

# Data type
- int
- string
- enum
- bool
- crypto
- Refer
## int
```go
// model config
{
    ...
    "foo": "int"
    ...
}
// item e.g.
{
    ...
    "foo": 1,
    "bar": 6.66
    ...
}
```
## string
```go
// model config
{
    ...
    "foo": "string"
    ...
}
// item e.g.
{
    ...
    "foo": "bar"
    ...
}
```
## enum
```go
// model config
{
    ...
    "foo": "enum"
    ...
}
// item e.g.
{
    ...
    "foo": [
        "bar1",
        "bar2",
        ...
    ]
    ...
}
```
## bool
```go
// model config
{
    ...
    "foo": "bool"
    ...
}
// item e.g.
{
    ...
    "foo": true
    ...
}
```
## crypto
```go
// model config
{
    ...
    "foo": "crypto"
    ...
}
// item e.g.
{
    ...
    "foo": "Y21kYmdv=="
    ...
}
```
## Refer
```go
// model config
{
    "foo": "Refer:<model_name>:<model_primary_key>:<display_item_key>"
}
// item e.g.
// 1:1
{
    ...
    "foo": "`users`61930d917a1e7253b8c80541310b8b63"
    ...
}
// 1:N
{
    ...
    "foo": [
        "`users`61930d917a1e7253b8c80541310b8b63",
        "`users`61930d917a1e7253b8c80541310b8b64",
        ...
        ]
    ...
}
```
# Default model and item
- users
    ```go
    {
        "id": "1",
        "name": "admin",
        "password": ""
    }
    ```
- groups
    ```go
    {
        "id": "1",
        "name": "admin",
        "users": "`users`:1"
    }
    ```