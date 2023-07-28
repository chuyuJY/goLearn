package api_service

import (
	_ "embed"
)

//go:embed chatgpt_preset/preQ.txt
var PreQ string

//go:embed chatgpt_preset/preA.txt
var PreA string
