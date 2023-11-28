package logging

// logParamsToZapParams converts ExtraKey to string and returns a slice of interface{}
func logParamsToZapParams(keys map[ExtraKey]interface{}) []interface{} {
	params := make([]interface{}, 0, len(keys))

	for k, v := range keys {
		params = append(params, string(k)) // ExtraKey to string
		params = append(params, v)         // value
	}

	return params
}

// logParamsToZeroParams converts ExtraKey to string and returns a map[string]interface{}
func logParamsToZeroParams(keys map[ExtraKey]interface{}) map[string]interface{} {
	params := map[string]interface{}{}

	for k, v := range keys {
		params[string(k)] = v // ExtraKey to string
	}

	return params
}
