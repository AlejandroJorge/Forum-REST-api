package util

func MergeAND(conditions []bool) bool {
	totalValue := true
	for _, v := range conditions {
		totalValue = totalValue && v
	}
	return totalValue
}

func MergeOR(conditions []bool) bool {
	totalValue := false
	for _, v := range conditions {
		totalValue = totalValue || v
	}
	return totalValue
}
