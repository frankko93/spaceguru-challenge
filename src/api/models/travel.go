package models

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
