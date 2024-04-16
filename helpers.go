package main

import (
	"encoding/binary"
	"errors"
)

func getString(packet []byte, offs *int) (string, error) {
	var str string
	var err error

	ends := *offs
	for {
		if *offs >= len(packet) {
			return "", errors.New("out of bounds")
		}
		if packet[ends] == 0x00 {
			break
		}
		ends++
	}
	str = string(packet[*offs:ends])
	*offs = ends + 1
	return str, err
}

func getUInt32(packet []byte, offs *int) (uint32, error) {

	if (*offs)+3 >= len(packet) {
		return 0, errors.New("out of bounds")
	}
	size := binary.BigEndian.Uint32(packet[*offs : (*offs)+4])
	(*offs) += 4
	return size, nil
}

func getUInt16(packet []byte, offs *int) (uint16, error) {
	if *offs+1 >= len(packet) {
		return 0, errors.New("out of bounds")
	}
	size := binary.BigEndian.Uint16(packet[*offs : (*offs)+2])
	(*offs) += 2
	return size, nil
}

func getByte(packet []byte, offs *int) (byte, error) {
	if *offs >= len(packet) {
		return 0, errors.New("out of bounds")
	}
	b := packet[*offs]
	(*offs) += 1
	return b, nil
}
