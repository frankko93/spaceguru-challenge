package models

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name""`
	Surname   string `json:"surname""`
	Type      string `json:"type""` // admin/driver
	Email     string `json:"email""`
	Password  string `json:"password"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Users struct {
	Page       int64  `json:"page"`
	PageSize   int64  `json:"pageSize"`
	TotalPages int64  `json:"totalPages"`
	Total      int64  `json:"total"`
	Data       []User `json:"data"`
}

type DriversSearchParams struct {
	Status   string `json:"status,omitempty"`
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize`
}

//rutas
type Travel struct {
	ID        int64  `json:"id"`
	UserID    string `json:"user_id"`
	VehicleID string `json:"vehicle_id"`
	Status    string `json:"status"`
	Route     string `json:"route"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}
