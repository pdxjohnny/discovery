package commands

var ConfigOptions = map[string]interface{}{
	"discovery": map[string]interface{}{
		"addr": map[string]interface{}{
			"value": "",
			"help":  "Address to send or bind to",
		},
		"port": map[string]interface{}{
			"value": "43089",
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
}
