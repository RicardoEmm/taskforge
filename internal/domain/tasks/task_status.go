package tasks

type TaskStatus string

const (
	TodoStatus       TaskStatus = "TODO"
	InProgressStatus TaskStatus = "IN_PROGRESS"
	DoneStatus       TaskStatus = "DONE"
)

func (t TaskStatus) IsValid() bool {
	switch t {
	case TodoStatus, InProgressStatus, DoneStatus:
		return true
	}
	return false
}
