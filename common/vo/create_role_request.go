package vo

type CreateRoleRequest struct {
	Sid   int    `json:"sid"`
	Name  string `json:"name"`
	Cover string `json:"cover"`
}
