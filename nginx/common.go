package nginx

import "strings"

func GetPropertyValue(source string, property string) string {
	if source != "" && property != "" {
		arr := strings.Split(source, " ")
		var sb strings.Builder
		if len(arr) > 0 {
			propertyName := strings.Trim(arr[0], " ")
			if strings.EqualFold(property, propertyName) {
				for i, element := range arr {
					if i > 0 {
						if i > 1 {
							sb.WriteString(" ")
						}
						sb.WriteString(element)
					}
				}
				propValue := strings.Trim(strings.Trim(strings.Replace(sb.String(), ";", "", -1), " "), "\n")
				return propValue
			}
		}
	}

	return ""
}

func HasProperty(source string, property string) bool {
	return strings.HasPrefix(strings.ToLower(source), property)
}
