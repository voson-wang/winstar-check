package crc32

func MPEG2(data []byte) uint32 {
	crc := uint32(0xffffffff)

	for _, v := range data {
		crc ^= uint32(v) << 24
		for i := 0; i < 8; i++ {
			if (crc & 0x80000000) != 0 {
				crc = (crc << 1) ^ 0x04C11DB7
			} else {
				crc <<= 1
			}
		}
	}
	return crc
}
