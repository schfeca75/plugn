package main

import (
	"bytes"
	"compress/gzip"
	"io"
)

// bashenv returns raw, uncompressed file data.
func bashenv() []byte {
	gz, err := gzip.NewReader(bytes.NewBuffer([]byte{
0x1f,0x8b,0x08,0x00,0x00,0x09,0x6e,0x88,0x00,0xff,0x9c,0x58,
0x6d,0x53,0xdb,0xb8,0x13,0x7f,0x9d,0x7c,0x8a,0x45,0x75,0x0b,
0xf9,0xff,0xcf,0xb8,0xc9,0xf5,0x66,0x6e,0xe0,0xd2,0x81,0xb6,
0x5c,0xcb,0x4c,0x8f,0x32,0x40,0x3b,0x77,0x53,0x68,0xc7,0xd8,
0x72,0xe2,0xc1,0x96,0x33,0x96,0xcc,0xc3,0xd1,0x7c,0xf7,0xdb,
0x95,0x64,0x5b,0x0e,0x06,0xda,0xbe,0x01,0x7b,0x25,0xed,0xfe,
0xf6,0xb7,0x0f,0xda,0x78,0x98,0x26,0xf0,0xf9,0x33,0x30,0x6f,
0x83,0x47,0xf3,0x02,0xbc,0x57,0xbb,0xc7,0xef,0xbe,0x7e,0xda,
0x3b,0x3a,0xde,0xff,0x70,0x00,0xdf,0x20,0xaa,0x14,0xf8,0xf1,
0x26,0xf8,0x09,0x8c,0x47,0x0c,0xfc,0x4c,0x01,0x7b,0xc1,0xe0,
0xec,0x6c,0x1b,0xd4,0x9c,0x8b,0xe1,0x40,0x9f,0x62,0x6b,0x6b,
0xf0,0x4f,0x51,0x95,0x20,0x6f,0xa4,0xe2,0x39,0xbc,0x0a,0xe5,
0x1c,0x52,0x09,0x05,0x9e,0x2e,0x12,0x88,0x43,0xc5,0xb7,0xba,
0xaa,0x99,0x73,0xf2,0x30,0xe3,0xa1,0xe4,0x50,0x2d,0x66,0x65,
0x18,0x73,0x50,0x85,0x39,0xff,0x02,0x8a,0x12,0x66,0x25,0xc7,
0xc3,0xe5,0x26,0xee,0x6f,0x80,0x56,0x22,0xcc,0x39,0x82,0x99,
0x4e,0x81,0xbd,0x09,0xcb,0xab,0x54,0xb8,0x80,0xb4,0x5e,0xfb,
0x0f,0xd8,0x07,0x01,0x1f,0x8e,0xe1,0xef,0x5f,0xa0,0x42,0x0b,
0xef,0x8a,0x9c,0x9f,0x97,0xfc,0x8a,0x4c,0xa4,0x42,0xaa,0x30,
0xcb,0x20,0x43,0xf5,0x52,0x69,0x8b,0x5b,0xac,0x3e,0xb6,0x0e,
0x00,0x1e,0xe8,0xad,0xf5,0xbe,0x73,0xdc,0xb0,0xbe,0xaa,0xfd,
0x04,0x2d,0x42,0x18,0xc7,0x90,0x2a,0xd2,0x19,0x70,0x15,0x05,
0x72,0xce,0xb3,0x4c,0x42,0x28,0x62,0x88,0xe6,0xa1,0x98,0x71,
0xb8,0x21,0x66,0xd0,0x7e,0xb9,0x2e,0x41,0xaf,0xae,0x1a,0x92,
0x55,0x5c,0xe0,0x0a,0xf8,0x11,0x30,0x2d,0x0e,0x2a,0x59,0x06,
0x59,0x11,0x85,0x59,0x70,0x9e,0x8a,0x80,0x6c,0xc3,0xcb,0x97,
0xae,0x7e,0xb6,0xde,0x55,0x11,0xcd,0xe9,0xbc,0xec,0x3b,0x79,
0x07,0xf5,0x2e,0x08,0xe2,0x80,0x97,0x79,0x2a,0xc2,0x0c,0x24,
0x97,0x32,0x2d,0x04,0xc5,0x8b,0x18,0xc4,0xc5,0x08,0x25,0x61,
0x79,0x43,0x2e,0xa9,0xf0,0x82,0x03,0x4f,0x12,0x1e,0xa9,0x4d,
0xd6,0x28,0x4a,0x52,0x0c,0xdf,0x35,0x3a,0x3d,0x19,0xe2,0x63,
0xcc,0xa3,0x2c,0x2c,0x39,0xf8,0xbb,0xf0,0xfa,0xaf,0x37,0xc7,
0xc3,0x61,0x94,0xc7,0x7e,0x96,0x4a,0xb5,0x31,0x82,0xdb,0xe1,
0xa0,0x5e,0x8e,0xb9,0x8c,0xa6,0xec,0x3d,0xca,0x91,0x9d,0xcb,
0x30,0xcd,0xc2,0xf3,0x8c,0x43,0x54,0xe4,0x39,0x72,0x25,0x59,
0xbb,0x51,0xc8,0x29,0xf3,0xc6,0x28,0xa8,0xf5,0xf8,0x17,0xfc,
0x46,0x62,0xe4,0x85,0x64,0x98,0x92,0x92,0xc7,0xc0,0x64,0x80,
0x6f,0x5b,0x41,0xc0,0x86,0xcb,0xd6,0x9e,0xde,0xd7,0x35,0xda,
0xe8,0x4a,0x30,0x99,0x2e,0x30,0x9a,0xa8,0xe6,0x76,0x8d,0x60,
0x7e,0xde,0x39,0x5b,0xb2,0x6d,0x88,0x5b,0x5e,0xbc,0x0b,0x02,
0x51,0x08,0x8e,0x46,0x30,0xef,0x16,0xc0,0xbe,0x90,0x11,0x6d,
0xb3,0x28,0x55,0xc7,0x92,0xb0,0x76,0x7e,0x46,0xad,0x7f,0x09,
0x5b,0x2b,0x3a,0xf9,0xf5,0x02,0xdf,0xfa,0xf8,0xda,0xd3,0x2b,
0xc8,0x18,0x24,0x95,0x88,0x14,0x05,0x2a,0xa4,0x37,0xcb,0x9b,
0x43,0x5b,0x22,0xb4,0xab,0xb8,0x8c,0xff,0x6f,0x27,0x5b,0xbe,
0x37,0x5e,0xe2,0xb2,0x4e,0x06,0xcd,0x83,0x65,0x41,0x10,0x5c,
0x6f,0xc3,0x71,0x65,0xd4,0xc5,0x9b,0x08,0xe6,0x32,0xe0,0x33,
0x78,0x86,0xb9,0x17,0xf3,0xcb,0x40,0x54,0x58,0x09,0xcf,0x9e,
0x19,0x56,0x85,0x75,0x6b,0x38,0xd0,0x7e,0x53,0x78,0xb6,0xbc,
0xdb,0x50,0x06,0x4f,0xf0,0xc9,0x0f,0x96,0xec,0x6c,0xaa,0x75,
0x75,0x7d,0x6c,0x98,0x5b,0x89,0x90,0x75,0xd7,0x9b,0x50,0x67,
0xb8,0x44,0xc4,0x28,0xa4,0x8d,0x80,0xb8,0x3a,0x8c,0x9c,0x32,
0x8f,0xfe,0x9f,0x52,0x36,0x92,0x5a,0x21,0xc1,0x1b,0xc3,0x29,
0x3b,0xf5,0x76,0x4e,0x91,0xf8,0xe1,0x60,0x69,0x33,0xc7,0xd8,
0x03,0x13,0x7e,0x0b,0x71,0xac,0x41,0x8d,0x1b,0x4c,0x35,0x98,
0x96,0x24,0x5c,0xdc,0xc6,0x5a,0x4c,0x13,0x05,0xb5,0x18,0x37,
0x76,0xe4,0xdf,0xbe,0x81,0x2a,0x2b,0x5e,0x2f,0x63,0x7f,0x50,
0x95,0x9c,0x3e,0xd7,0x1d,0xaa,0x66,0xb5,0xc9,0x56,0x4d,0xe3,
0x17,0x0f,0xe5,0xa7,0x5e,0x87,0xc7,0xa6,0x5f,0x79,0xb7,0x0e,
0x7f,0xb8,0x8f,0x9d,0x2d,0xf1,0xf4,0x0e,0xf1,0x90,0x49,0xb4,
0xd2,0x34,0x3e,0x5a,0x73,0x1b,0x9d,0x8d,0xd7,0x01,0xf6,0x8e,
0x2a,0x9a,0xd7,0x19,0x81,0x8d,0x96,0x36,0xd2,0xba,0x45,0x36,
0xa1,0xd0,0x66,0xb5,0x16,0x82,0x75,0x57,0x89,0xb7,0x91,0x08,
0x9f,0x78,0x35,0x3b,0x46,0x74,0x9e,0xaa,0x7c,0xa5,0x73,0xdc,
0x29,0x5b,0xdd,0xca,0x28,0xab,0xd0,0x66,0x37,0xaf,0xac,0x22,
0x9b,0x5a,0x83,0x45,0x99,0x0a,0x95,0x00,0x03,0x78,0xea,0x4f,
0x7e,0x97,0xf0,0x54,0x9e,0x62,0x9a,0x59,0xa7,0x3a,0xe6,0xef,
0xd2,0x61,0xe0,0x98,0x5c,0x6b,0x00,0x51,0xfb,0xf1,0x8c,0x87,
0xba,0x21,0xd9,0x88,0x62,0x77,0x5c,0xf4,0xd5,0xd1,0xf1,0xbc,
0xb8,0x92,0x40,0xab,0x08,0x13,0x01,0xe7,0xa1,0xae,0x25,0x82,
0xde,0x57,0x4c,0x61,0x39,0xa3,0x64,0xd8,0x71,0xee,0x1d,0x12,
0x39,0xd4,0x61,0xdb,0x85,0xda,0x73,0x72,0x52,0xe7,0x04,0x3d,
0x60,0xa1,0x3f,0xc1,0x6b,0x05,0x19,0xc0,0x13,0x66,0x9b,0x93,
0x5e,0xb7,0xa4,0x26,0x78,0x4a,0x9e,0x05,0xd4,0x28,0x8c,0x50,
0xa0,0x08,0xba,0xef,0x10,0xf8,0x98,0xc9,0xce,0x71,0x5d,0xdf,
0x3d,0xe4,0x58,0x24,0xc2,0x27,0xbf,0x6c,0xf9,0x8e,0x9b,0xe4,
0xb1,0x25,0xa2,0x8b,0x9f,0x48,0x72,0x2a,0xa3,0xa6,0x4b,0xb3,
0x32,0x1c,0xa2,0x06,0xc2,0xd6,0x47,0xde,0xbe,0x90,0x0b,0xbc,
0x01,0x9c,0x26,0x84,0x57,0x19,0x6e,0xae,0x72,0x2e,0x94,0x6c,
0x9a,0x0c,0x4a,0xb2,0x54,0xf0,0xa9,0xb7,0xa1,0x6e,0x16,0x9c,
0xca,0xd2,0xe6,0x7f,0xad,0xad,0x6d,0x80,0xcc,0x35,0x40,0x75,
0x32,0xe7,0x61,0x0c,0xfe,0x78,0x64,0xe7,0x02,0x9f,0x83,0xe1,
0x8a,0x14,0x06,0x48,0x06,0xc3,0x74,0x59,0xd2,0xc6,0xf0,0xea,
0x02,0xfc,0x3f,0xa7,0xb0,0x1e,0x4c,0x83,0x5b,0x9d,0x55,0xc0,
0xfe,0xa0,0xf2,0x7c,0xc9,0x96,0xeb,0x40,0xa5,0x09,0x4c,0xa7,
0x16,0xe8,0x2a,0xb7,0x79,0xf5,0xdd,0x5e,0xd1,0x52,0x99,0x2e,
0xe8,0x4d,0xa7,0x03,0x6d,0x6c,0x5b,0xd2,0x5d,0xc7,0x30,0x65,
0x5b,0xf0,0x18,0x3f,0x33,0x45,0x91,0xdc,0x5a,0xa7,0xa8,0x3c,
0x60,0xdd,0xed,0xec,0x3d,0xcd,0x5c,0x52,0xda,0xe2,0xfc,0x10,
0xf1,0xba,0x37,0xd6,0x3d,0x1a,0x74,0xcd,0x50,0xc4,0x00,0xdf,
0x46,0xcd,0x12,0x40,0x5b,0x4c,0xce,0x42,0x9b,0xc7,0xad,0xca,
0x4e,0x23,0x30,0x9e,0xa1,0x5e,0xa4,0x10,0xab,0x1c,0x7c,0x01,
0xff,0x9f,0xb8,0x37,0x3e,0xba,0x63,0x87,0xa1,0x7b,0xdc,0xd1,
0x73,0x52,0xa8,0x87,0x8b,0x45,0x56,0xcd,0xb0,0x1b,0x24,0x65,
0x91,0xa3,0xe4,0x2d,0x96,0xea,0xc7,0xa3,0xf7,0x8e,0x7b,0x55,
0x99,0x19,0xff,0x68,0x94,0xb3,0x9e,0x45,0x78,0xad,0x7b,0x87,
0xef,0x3f,0xbe,0xdd,0x3f,0xf8,0x7a,0xb8,0x7b,0xf2,0x2e,0x68,
0x86,0x04,0x5c,0x9c,0xa1,0x8a,0x28,0xa3,0xda,0x62,0x1e,0x1e,
0x66,0xe0,0x4d,0xf4,0x09,0x1f,0x9c,0x96,0x4a,0x10,0x2b,0xf1,
0x00,0xc8,0x23,0x9e,0x17,0x97,0xbc,0x8b,0xae,0x69,0x69,0x46,
0xea,0x0e,0x22,0x46,0x62,0x07,0x88,0x32,0x07,0xbf,0x4c,0x56,
0x20,0x72,0x41,0x47,0xe3,0xc0,0x33,0x5b,0xef,0xdb,0xd6,0x18,
0x69,0x37,0x22,0xd4,0x87,0x46,0x24,0xd0,0xb3,0xa9,0xae,0xaa,
0x16,0x17,0x46,0x6e,0xa1,0x68,0xca,0x23,0x67,0x67,0x59,0x71,
0x6e,0x2e,0xf4,0x45,0xa8,0xe6,0xba,0xf7,0xf6,0x9b,0xfc,0x9f,
0x6d,0xc1,0xae,0x36,0xf4,0x69,0x03,0x67,0x43,0x4e,0xf4,0x83,
0x47,0x0a,0x74,0x8b,0xed,0x5c,0x68,0xb8,0x05,0xf3,0x05,0x0b,
0xf1,0x21,0x8f,0x31,0x83,0x68,0x20,0xd0,0xa9,0x67,0x97,0xe8,
0x7e,0xd4,0xef,0x71,0x2a,0xb5,0xc0,0x51,0x7d,0xc9,0x4b,0x9a,
0x35,0x49,0x77,0x14,0x2a,0x63,0x39,0x30,0x9a,0x36,0x55,0x91,
0x67,0x94,0x7c,0xf8,0xcf,0x9f,0x71,0xac,0x6a,0x4b,0x15,0x30,
0x7b,0x8a,0x39,0x8a,0xec,0xb4,0xf0,0x23,0x5a,0xdc,0xca,0xd6,
0x9a,0x3a,0x57,0xd2,0x73,0xbc,0x92,0xfc,0xdf,0xe8,0xcf,0xf8,
0x79,0x7b,0x3b,0x35,0x87,0xbd,0x1a,0x03,0x15,0x90,0xe6,0x87,
0x9e,0x74,0xf3,0xaa,0xa7,0x20,0x1b,0x9d,0xaa,0x8d,0x0e,0xc6,
0x58,0x95,0xe9,0x6c,0xc6,0xcb,0xbe,0x30,0x9f,0x98,0x25,0xbc,
0x94,0x8a,0x42,0x8f,0x90,0x35,0x7f,0x77,0x13,0x91,0x76,0xb8,
0x23,0xc8,0xfd,0xa9,0x60,0x82,0xb2,0x9a,0x0c,0x75,0xd0,0xea,
0x54,0xa8,0xbb,0x59,0x54,0x88,0x24,0x9d,0xd5,0xf7,0x41,0x27,
0x25,0xb4,0xa2,0xd1,0x88,0x6e,0x97,0xc1,0x80,0xf2,0xe0,0xba,
0xa1,0x23,0xf0,0x08,0x0f,0xb3,0xa1,0xef,0x08,0xed,0xe0,0xf2,
0x00,0x21,0x06,0x4a,0xef,0xa4,0xab,0x57,0xb0,0x5d,0x34,0xb5,
0xd4,0x5b,0x87,0x99,0x20,0xb7,0x1f,0xab,0xaf,0xc7,0xea,0x14,
0x91,0xd8,0xf4,0xec,0x83,0xf2,0xc6,0x2c,0x3d,0x8a,0x05,0x8b,
0xfd,0x71,0x3b,0x96,0x64,0x4c,0xc6,0x3e,0x53,0x6f,0x31,0x47,
0x6d,0xd4,0xcc,0xc6,0xaa,0x0c,0x57,0xee,0x02,0xc7,0x64,0xa7,
0x5f,0x86,0x6a,0xc5,0xba,0x51,0xa0,0x4b,0x80,0x75,0x6a,0xc0,
0x61,0x85,0x14,0xb8,0xb0,0x1e,0xfb,0xe5,0xf1,0x03,0xe0,0xbe,
0x17,0x52,0x33,0x98,0xdf,0x25,0x49,0xf6,0x93,0x74,0xfc,0x93,
0x24,0x01,0xe6,0x79,0x45,0x8f,0xbf,0xd6,0xd7,0xa3,0x4f,0xbf,
0xd4,0xb4,0xb4,0x81,0x23,0xf9,0x83,0xa0,0x7b,0xc9,0x4b,0x45,
0xda,0x0b,0x74,0x1f,0xe5,0x69,0x98,0xa5,0xff,0x62,0xee,0x60,
0x41,0xe7,0x0b,0x75,0x53,0x03,0xa7,0x26,0x85,0x28,0xf2,0x8b,
0x38,0x2d,0xc1,0x5f,0x74,0x4d,0xe2,0x82,0x2a,0x68,0x88,0xbf,
0x1f,0x48,0x7d,0xb4,0x37,0xe1,0xee,0x59,0x75,0x6f,0x50,0x33,
0x1d,0xb4,0x00,0x63,0x0b,0xef,0xb0,0x9a,0x09,0x17,0xe3,0x6a,
0xf3,0xd0,0xfe,0x0e,0xf3,0x30,0x15,0xc6,0x61,0xa2,0xcb,0xe7,
0x05,0x2c,0xd2,0x05,0x4f,0x50,0xfd,0xb6,0x99,0x2a,0x4e,0x8e,
0x76,0x5f,0xef,0xd5,0x3d,0x41,0xef,0xb9,0xa6,0x6f,0x04,0x3a,
0xd2,0xfa,0xc3,0xcf,0xde,0xc1,0xa7,0x69,0x3d,0x84,0xac,0xad,
0x78,0xbf,0xfa,0xf9,0xc6,0x7c,0x16,0x6a,0x37,0xd0,0xb7,0x09,
0x51,0x28,0xad,0x57,0x37,0xca,0xcb,0xb4,0x2c,0x04,0x0d,0x9f,
0xac,0xfe,0x29,0x30,0xd1,0x33,0x4a,0xe7,0x97,0x9f,0x1d,0x04,
0x3a,0xb2,0x66,0x3c,0xe8,0x48,0xe9,0x26,0xee,0x08,0x6c,0xdb,
0xee,0xc8,0x0c,0xd5,0x1d,0x91,0x6d,0x21,0xae,0x8c,0x86,0x6d,
0x13,0x36,0x60,0x87,0xfd,0x29,0xeb,0x8e,0xe0,0x4d,0x6f,0xe8,
0x13,0x9b,0xb7,0xbe,0x15,0xb9,0x72,0x80,0xf2,0x71,0x38,0x18,
0xb6,0xd3,0xbe,0x69,0xc4,0xcb,0xff,0x02,0x00,0x00,0xff,0xff,
0xe9,0xe0,0xba,0x8d,0xde,0x13,0x00,0x00,
	}))

	if err != nil {
		panic("Decompression failed: " + err.Error())
	}

	var b bytes.Buffer
	io.Copy(&b, gz)
	gz.Close()

	return b.Bytes()
}