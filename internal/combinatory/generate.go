package combinatory

/*
Generate generates all combinations of parameters passed as a slice
of slices of possible values for a parameter
*/
func Generate(arr [][]any) [][]any {
	// Return empty slice if nothing to combine
	if len(arr) == 0 {
		return [][]any{}
	}

	// Result is a slice of slices that will hold each parameter combination
	result := [][]any{{}}

	// traverse the slice of parameters value slice
	for _, parameter := range arr {
		var temp [][]any
		// traverse the values in this parameter value slice
		for _, value := range parameter {
			// append the value to each combination
			for _, combination := range result {
				temp = append(temp, append(combination, value))
			}
		}
		// update result
		result = temp
	}

	return result
}
