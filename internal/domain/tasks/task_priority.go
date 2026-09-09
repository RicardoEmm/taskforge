package tasks

type TaskPriority string

const (
	LowPriority    TaskPriority = "LOW"
	MediumPriority TaskPriority = "MEDIUM"
	HighPriority   TaskPriority = "High"
)

func (p TaskPriority) IsValid() bool {
	switch p {
	case LowPriority, MediumPriority, HighPriority:
		return true
	}
	return false
}
