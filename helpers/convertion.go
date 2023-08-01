package helpers

import "strconv"

func StringToUint(str string) (uint, error) {
	if str == "" {
		return 0, nil
	}

	uint_64, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(uint_64), nil
}

func StringToInt64(str string) (int64, error) {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return num, nil
}
