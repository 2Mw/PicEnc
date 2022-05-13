package config

import "time"

const (
	WaitToExit = time.Second * 2
)

type EncryptMode int

const (
	Row        EncryptMode = 1
	RowAndLine EncryptMode = 2
)
