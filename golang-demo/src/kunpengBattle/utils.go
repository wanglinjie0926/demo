package kunpengBattle

import (
	"strconv"
)

func splitPackage(data []byte, atEOF bool) (int, []byte, error) {
	if atEOF && len(data) < 5 {
		return 0, nil, nil
	}

	size, err := strconv.ParseInt(string(data[:5]), 10, 32)

	if err != nil {
		return 0, nil, err
	}

	totalSize := int(size + 5)
	if totalSize > len(data) {
		return 0, nil, nil
	}

	msgBytes := make([]byte, size, size)
	copy(msgBytes, data[5:totalSize])

	return totalSize, msgBytes, nil
}

// func decode(msgBytes []byte, {}interface) error {

// }

// func encode({}interface) []byte{}
