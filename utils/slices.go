package utils

func SliceContainsString(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func ToGenericArray(arr ...interface{}) []interface{} {
    return arr
}