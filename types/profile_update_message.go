package types

type ProfileUpdateMessage struct {
	EmployeeId string `json:"employee_id"`
	Roles []string `json:"roles"`
}
