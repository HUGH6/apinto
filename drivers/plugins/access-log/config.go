package access_log

import "github.com/eolinker/eosc"

type Config struct {
	Output []eosc.RequireId `json:"output" skill:"github.com/eolinker/goku/output.output.IOutput"`
}
