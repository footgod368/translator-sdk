package utils

import "encoding/json"

func MarshalJSONToString(v any) string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}
