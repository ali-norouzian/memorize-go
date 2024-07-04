package reflection

import "reflect"

// StructToMap converts a struct to a map
func StructToMap(data any) map[string]any {
	result := make(map[string]any)
	value := reflect.ValueOf(data)
	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		result[field.Name] = value.Field(i).Interface()
	}
	return result
}
