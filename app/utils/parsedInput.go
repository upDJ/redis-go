package utils

type ParsedInput struct {
	args       string
	key        string
	val        string
	valLength  string
	expiryFlag string
	expiryTime string
	nullBulk   string
}

func (p ParsedInput) toGetArr() []string {
	return []string{p.val}
}
