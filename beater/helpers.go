package beater

func stringArrayToStringBoolMap(array []string) map[string]bool  {
	stringBoolMap := make(map[string]bool)
	for _, tag := range array {
		stringBoolMap[tag] = true
	}
	return stringBoolMap
}
