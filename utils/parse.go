package utils

type TableContents struct {
	selections []Selection
}
type Selection struct {
	label  string
	qlinks []Qlink
}
type Qlink struct {
	number string
	link   string
}
