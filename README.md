# saclog

`saclog` implements a syslog server executable for parsing UDP and TCP syslog messages received on port 514.

It is supposed to support the two most common syslog formats, RCF 3164 and RCF 5424, and write them to standard out in a simplified format:

```
LOG_TIME APP_NAME APP_HOSTNAME MESSAGE
```

## Installation

View the [**latest release**](https://github.com/mailsac/saclog/releases/latest)

Run the executable in linux:

```
./saclog
```

Now you can send **TCP** and **UDP** syslog messages to port `514`.

## Security 

There is no security - a syslog server should be internal to your network. If that is not possible, using iptables to restrict traffic may be a minimal security layer.


## License

MIT

Also uses [mcuadros/go-syslog](https://github.com/mcuadros/go-syslog) which is published under the MIT license.
