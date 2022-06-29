package super_reader

import (
	SUPER_CONSTANTS "main/modules/go-main/SUPER-MODS/consts"
)

func Rem(data []string) []string {
	for _, file_data := range data {
		if _, value := SUPER_CONSTANTS.SUPER_KEY_STR_BOOL[file_data]; !value {
			SUPER_CONSTANTS.SUPER_KEY_STR_BOOL[file_data] = true
			SUPER_CONSTANTS.SUPER_DATA_STRING = append(SUPER_CONSTANTS.SUPER_DATA_STRING, file_data)
		}
	}
	return SUPER_CONSTANTS.SUPER_DATA_STRING
}
