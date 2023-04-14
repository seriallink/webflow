package webflow

func In(value string, list ...string) bool {

	for _, element := range list {
		if value == element {
			return true
		}
	}

	return false
}

func NotIn(value interface{}, list ...interface{}) bool {

	for _, element := range list {
		if value == element {
			return false
		}
	}

	return true

}

func StringPtr(value string) *string {
	return &value
}

func IntPtr(value int) *int {
	return &value
}

func FloatPtr(value float64) *float64 {
	return &value
}

func BoolPtr(value bool) *bool {
	return &value
}

func Ternary(condition bool, x, y interface{}) interface{} {

	if condition {
		return x
	}

	return y

}

func TernaryStr(condition bool, x, y string) string {
	return Ternary(condition, x, y).(string)
}

func TernaryInt(condition bool, x, y int) int {
	return Ternary(condition, x, y).(int)
}
