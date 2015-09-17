package commands

import (
	key "github.com/pdxjohnny/key/commands"
)

const (
	discovery_port = "43089"
	frontend_port  = "25001"
)

var ConfigOptions = map[string]interface{}{
	"discovery": map[string]interface{}{
		"addr": map[string]interface{}{
			"value": "",
			"help":  "Address to send or bind to",
		},
		"port": map[string]interface{}{
			"value": discovery_port,
			"help":  "Port to send or bind to",
		},
		"send": map[string]interface{}{
			"value": "",
			"help":  "A message to send",
		},
		"online": map[string]interface{}{
			"value": "",
			"help":  "A message to send severy interval of time",
		},
		"int": map[string]interface{}{
			"value": 5,
			"help":  "Interval to send messages and detirmine node timeout",
		},
	},
	"proxy": map[string]interface{}{
		"addr": map[string]interface{}{
			"value": "",
			"help":  "Address to bind to",
		},
		"port": map[string]interface{}{
			"value": "8080",
			"help":  "Port to bind to",
		},
		"url": map[string]interface{}{
			"value": "",
			"help":  "Url to reverse proxy to",
		},
		"discover": map[string]interface{}{
			"value": true,
			"help":  "Listen for frontends to proxy to",
		},
		"dAddr": map[string]interface{}{
			"value": "",
			"help":  "Discovery address to bind to",
		},
		"dPort": map[string]interface{}{
			"value": discovery_port,
			"help":  "Discovery port to bind to",
		},
	},
	"frontend": map[string]interface{}{
		"addr": map[string]interface{}{
			"value": "",
			"help":  "Address to bind to",
		},
		"port": map[string]interface{}{
			"value": frontend_port,
			"help":  "Port to bind to",
		},
		"discover": map[string]interface{}{
			"value": "localhost:" + discovery_port,
			"help":  "Proxy host to boardcast online to",
		},
	},
	"key": key.ConfigOptions,
}
