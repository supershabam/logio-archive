# logio

Logio acts somewhat like statsd or rsyslog for forwarding logs to a localhost agent that adds metadata and deals with backpressure and "up"-ness.

Libraries for various languages will send log messages to the emitio agent, preferrably buffered. The connection is made over TCP and exchanges proto back and forth.

TCP is useful for the ability to recognize when a message is flushed and handling larger messages (though the library should be configurable to cap messages to a certain amount).

```query
{
    select: [
        {
            type: "histogram",
            buckets: "auto",
            on: "latency",
            as: "time_seconds",
            width: "15m"
        }
    ],
    filter: [
        {
            type: "service-group",
            eq: "nginx"
        },
        {

        }
    ],
    pivot: [
        {
            type: "url",
            field: "url",
            on: "path[0]"
        }
    ],
    over: '2h'

    }
}
select histogram(Seconds(latency))
where 
```