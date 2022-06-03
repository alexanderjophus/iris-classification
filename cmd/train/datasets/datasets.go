package datasets

import (
	_ "embed"
)

//go:embed iris.csv
var Data []byte
