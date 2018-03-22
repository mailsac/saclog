package main

import (
	"fmt"
	"strings"
	"time"

	"gopkg.in/mcuadros/go-syslog.v2"
	"gopkg.in/mcuadros/go-syslog.v2/format"
)

func main() {
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	server := syslog.NewServer()
	server.SetFormat(syslog.Automatic)
	server.SetHandler(handler)
	err := server.ListenUDP("0.0.0.0:514")
	if err != nil {
		panic(err)
	}
	fmt.Println("started UDP syslog server")
	err = server.ListenTCP("0.0.0.0:514")
	if err != nil {
		panic(err)
	}
	fmt.Println("started TCP syslog server")

	server.Boot()

	go func(channel syslog.LogPartsChannel) {
		logTime := time.Now().UTC().Format(time.RFC3339)

		for logParts := range channel {
			if logParts["hostname"] == "" { // goofy golang syslog package doesn't parse out content all the way
				contentParts := parseGolangSyslog(fmt.Sprintf("%s", logParts["content"]))
				logParts["hostname"] = contentParts["hostname"]
				logParts["tag"] = contentParts["tag"]
				logParts["content"] = contentParts["content"]
			}

			// RCF 3164:
			// 		map[
			// 			facility:16
			// 			severity:6
			// 			client:23.95.104.109:32799
			// 			tls_peer:
			// 			timestamp:2017-11-27 23:43:26 +0000 UTC
			// 			hostname:acai
			// 			content:crunchberry.mailsac.com nXljMY5iWgLSyggl1QI0kreoYt3rEG2F-0
			// 			tag:uploading priority:134
			// 		]
			//
			// RCF 5424:
			// 		map[
			// 			client:127.0.0.1:49427
			// 			priority:134
			// 			facility:16
			// 			severity:6
			// 			app_name:node
			// 			structured_data:-
			// 			message:frontend message.getFile remote fetch http://gooseberry.mailsac.com/asdf@mailsac.com/2017-11-28/DDCb8I8ifgvh2hzyaHebmLXZDRhNzSAf-0/raw.txt
			// 			version:1
			// 			timestamp:2017-12-05 00:53:37.148 +0000 UTC
			// 			hostname:Jeff-MBP15-Digium.local
			// 			proc_id:56630
			// 			msg_id:-
			// 			tls_peer:
			// 		]

			// RCF 5424
			if fmt.Sprint(logParts["version"]) == "1" {
				fmt.Println(logTime, logParts["app_name"], logParts["hostname"], logParts["message"])
				continue
			}

			// RCF 3164
			if fmt.Sprint(logParts["tag"]) != "<nil>" {
				fmt.Println(logTime, logParts["tag"], logParts["hostname"], logParts["content"])
				continue
			}

			fmt.Println(logTime, "format unknown", logParts)
		}
	}(channel)

	server.Wait()
}

// parseGolangSyslog parses a weirdly formatted golang syslog like:
// 		2017-12-04T19:26:45-08:00 goji.mailsac.com inbound[61504]: Starting server on port 25
func parseGolangSyslog(s string) (logParts format.LogParts) {
	logParts = make(format.LogParts)
	logParts["version"] = "1"

	metaParts := strings.Split(s, " ")
	messageParts := strings.Split(s, ":")

	if len(metaParts) < 2 {
		return logParts
	}
	logParts["hostname"] = metaParts[1]

	if len(messageParts) < 5 {
		return logParts
	}
	logParts["content"] = strings.Join(messageParts[4:], ":")[1:]

	appParts := strings.Split(metaParts[2], "[")
	if len(appParts) < 2 {
		return logParts
	}

	logParts["tag"] = appParts[0]

	return logParts
}
