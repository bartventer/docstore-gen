package generate

import "strings"

func getPackageName(fullName string) string {
	return strings.Split(delPointerSym(fullName), ".")[0]
}

func delPointerSym(name string) string {
	return strings.TrimLeft(name, "*")
}

func getPureName(s string) string {
	return string(strings.ToLower(delPointerSym(s))[0])
}

// not need capitalize
func getStructName(t string) string {
	list := strings.Split(t, ".")
	return list[len(list)-1]
}

func uncaptialize(s string) string {
	if s == "" {
		return ""
	}

	return strings.ToLower(s[:1]) + s[1:]
}

func isCapitalize(s string) bool {
	if len(s) < 1 {
		return false
	}
	b := s[0]
	if b >= 'A' && b <= 'Z' {
		return true
	}
	return false
}
