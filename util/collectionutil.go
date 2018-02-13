package util

import "sort"

func GetIntMapKeys(originMap *map[int]string) []int {

	keys := make([]int, len(*originMap))
	i := 0
	for k := range *originMap {
		keys[i] = k
		i++
	}

	sort.Ints(keys);
	return keys;
}