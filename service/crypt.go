package service

import (
	"PicEnc/config"
)

type ChaosCrypt struct {
	Key  string
	Mode config.EncryptMode
	Func func(x float64) float64
}

func (c *ChaosCrypt) Encrypt(data [][]uint32) [][]uint32 {
	return nil
}

func (c *ChaosCrypt) Decrypt(data [][]uint32) [][]uint32 {
	return nil
}
