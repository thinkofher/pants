## pants-app

`pants-app` is a backend part of the pants web application, which main purpose is to provide simple REST APIs for shortening links.

## Routes

### `/api/short`

##### POST:

Assigns seven character long random string to given link in a value field of request.

Sample request body:

```
{
    "value": "www.duckduckgo.com"
}
```

Sample respond body:
```
{
    "key": "RPLyq8h",
    "url": "www.duckduckgo.com"
}
```

Example of usage with `curl`:

    $ curl \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"value":"www.duckduckgo.com"}' \
    http://localhost:8080/api/short
    {"key":"RPLyq8h","url":"www.duckduckgo.com"}

### `/api/short/{key}`

##### GET:

Where `{key}` is a random string generted by application when shortening some link.

Responds with a JSON body with given `{key}` and the link assigned to it.

Sample respond body:
```
{
    "key": "RPLyq8h",
    "url": "www.duckduckgo.com"
}
```

Example of usage with `curl`:

    $ curl \
    -X GET \
    -H "Content-Type: application/json" \
    http://localhost:8080/api/short/RPLyq8h
    {"key":"RPLyq8h","url":"www.duckduckgo.com"}
