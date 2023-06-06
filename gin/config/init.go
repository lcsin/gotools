package config

import (
	"flag"
)

var _config Config
var _path *string

func init() {
	_path = flag.String("c", "config/config.yaml", "-c config file path")
}
