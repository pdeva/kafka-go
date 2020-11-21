package protocol

import (
	"reflect"
)

func sizeOf(typ reflect.Type) int {
	switch typ.Kind() {
	case reflect.Bool, reflect.Int8:
		return 1
	case reflect.Int16:
		return 2
	case reflect.Int32:
		return 4
	case reflect.Int64:
		return 8
	default:
		return 0
	}
}

func sizeOfString(s string) int {
	return 2 + len(s)
}

func sizeOfBytes(b []byte) int {
	return 4 + len(b)
}

func sizeOfVarString(s string) int {
	return sizeOfVarInt(int64(len(s))) + len(s)
}

func sizeOfCompactString(s string) int {
	return sizeOfUnsignedVarInt(int64(len(s)+1)) + len(s)
}

func sizeOfVarBytes(b []byte) int {
	return sizeOfVarInt(int64(len(b))) + len(b)
}

func sizeOfCompactBytes(b []byte) int {
	return sizeOfUnsignedVarInt(int64(len(b)+1)) + len(b)
}

func sizeOfVarNullString(s string) int {
	n := len(s)
	if n == 0 {
		return sizeOfVarInt(-1)
	}
	return sizeOfVarInt(int64(n)) + n
}

func sizeOfCompactNullString(s string) int {
	n := len(s)
	if n == 0 {
		return sizeOfUnsignedVarInt(0)
	}
	return sizeOfUnsignedVarInt(int64(n)+1) + n
}

func sizeOfVarNullBytes(b []byte) int {
	if b == nil {
		return sizeOfVarInt(-1)
	}
	n := len(b)
	return sizeOfVarInt(int64(n)) + n
}

func sizeOfVarNullBytesIface(b Bytes) int {
	if b == nil {
		return sizeOfVarInt(-1)
	}
	n := b.Len()
	return sizeOfVarInt(int64(n)) + n
}

func sizeOfCompactNullBytes(b []byte) int {
	if b == nil {
		return sizeOfUnsignedVarInt(0)
	}
	n := len(b)
	return sizeOfUnsignedVarInt(int64(n)+1) + n
}

func sizeOfVarInt(i int64) int {
	u := uint64((i << 1) ^ (i >> 63)) // zig-zag encoding
	n := 0

	for u >= 0x80 {
		u >>= 7
		n++
	}

	return n + 1
}

func sizeOfUnsignedVarInt(i int64) int {
	n := 0

	for i >= 0x80 {
		i >>= 7
		n++
	}

	return n + 1
}
