package utils


func GenerateOrderString(order []string) string {
	orderStr := ""

	for idx, o := range order {
		if idx == 0 {
			orderStr += "ORDER BY "
		}

		if ask := o[0]; ask == '-' {
			orderStr += ToSnakeCase(o[1:]) + " DESC"
		} else {
			orderStr += ToSnakeCase(o) + " ASC"
		}

		if idx + 1 != len(order) {
			orderStr += ", "
		}
	}

	return orderStr
}