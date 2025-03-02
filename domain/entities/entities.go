package entities

type ChangeOperation struct {
	TargetId      string
	OperationType string
	Args          map[string]string
}
