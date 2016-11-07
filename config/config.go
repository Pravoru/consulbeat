// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period          time.Duration `config:"period"`
	ConsulURL       string        `config:"consul_url"`
	ServicesTags    []string        `config:"tags"`
	FailOnHttpError bool        `config:"fail_on_http_error"`
}

var DefaultConfig = Config{
	1 * time.Second,
	"",
	nil,
	true,
}
