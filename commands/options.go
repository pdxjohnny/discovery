package commands

import (
	key "github.com/pdxjohnny/key/commands"
)

const (
	discovery_addr = "231.213.134.213"
	discovery_port = "43089"
	frontend_addr  = ""
	frontend_port  = "25001"
	proxy_addr  = ""
	proxy_port     = "8080"
)

var ConfigOptions = map[string]interface{}{
	"discovery": map[string]interface{}{
		"addr": map[string]interface{}{
			"value": discovery_addr,
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
		"key": map[string]interface{}{
			"value": "keys/id_rsa.pub",
			"help":  "Key to encrypt with (no .pub if server)",
		},
	},
	"proxy": map[string]interface{}{
		"addr": map[string]interface{}{
			"value": proxy_addr,
			"help":  "Address to bind to",
		},
		"port": map[string]interface{}{
			"value": proxy_port,
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
			"value": discovery_addr,
			"help":  "Discovery address to bind to",
		},
		"dPort": map[string]interface{}{
			"value": discovery_port,
			"help":  "Discovery port to bind to",
		},
	},
	"frontend": map[string]interface{}{
		"addr": map[string]interface{}{
			"value": frontend_addr,
			"help":  "Address to bind to",
		},
		"port": map[string]interface{}{
			"value": frontend_port,
			"help":  "Port to bind to",
		},
		"dAddr": map[string]interface{}{
			"value": discovery_addr,
			"help":  "Discovery address to send to",
		},
		"dPort": map[string]interface{}{
			"value": discovery_port,
			"help":  "Discovery port to send to",
		},
	},
	"key": key.ConfigOptions,
}
