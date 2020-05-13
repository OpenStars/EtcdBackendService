package format

func I2Byte(u int64) []byte {
	return []byte(I2S(u))
}
