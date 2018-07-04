# tdrug

`tdrug` is **T**ime **D**ependent **R**edirect **U**RL **G**enerator.

## Getting Started

Start server:

    $ docker run --rm -p 8991:8991 t13a/tdrug

Request URL:

    $ curl -D - http://localhost:8991/?f=http://example.com/%25Y/%25m/%25d

Then generate a redirection:

    HTTP/1.1 302 Found
    Content-Type: text/html; charset=utf-8
    Location: https://example.com/2018/07/04
    Date: Wed, 04 Jul 2018 07:40:56 GMT
    Content-Length: 53
    
    <a href="https://example.com/2018/07/04">Found</a>.

It is possible to specify the time zone and offset time:

    $ curl -D - http://localhost:8991/Asia/Tokyo?f=http://example.com/%25Y/%25m/%25d&o=1h

By default, the time zone is assumed as UTC.

## Query parameters

Name | Description | Default value
--- | --- | ---
`f` | Format string | N/A
`o` | Offset time in [Duration](https://golang.org/pkg/time/#Duration) | `0s`

## Format syntax

Similar to `date` UNIX command. [Percent-encoding](https://en.wikipedia.org/wiki/Percent-encoding) is needed.

Placeholder | Description
--- | ---
`%d` | Day of month (`01` .. `31`)
`%H` | Hour (`00` .. `23`)
`%m` | Month (`01` .. `12`)
`%M` | Minute (`00` .. `59`)
`%s` | Seconds since 1970-01-01 00:00:00 UTC
`%S` | Second (`00` .. `60`)
`%Y` | Year (eg: `2018`)
