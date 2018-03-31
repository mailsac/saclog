# saclog

`saclog` implements a syslog server executable for parsing UDP and TCP syslog messages received on port 514.

It is supposed to support the two most common syslog formats, RCF 3164 and RCF 5424, and write them to standard out in a simplified format:

```
LOG_TIME APP_NAME APP_HOSTNAME MESSAGE
```

## installation

View the [latest release](https://github.com/mailsac/saclog/releases/latest)

Run the executable in linux:

```
./saclog
```

## License

MIT

Also uses [mcuadros/go-syslog](https://github.com/mcuadros/go-syslog) which is published under the MIT license.
