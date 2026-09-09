package users

type UserRole string

const (
	AdminRole  UserRole = "ADMIN"
	MemberRole UserRole = "MEMBER"
)

func (r UserRole) IsValid() bool {
	switch r {
	case AdminRole, MemberRole:
		return true
	}
	return false
}
