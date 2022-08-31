package infra

type stringValue interface {
	StringValue(v *string) string
}

func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}
