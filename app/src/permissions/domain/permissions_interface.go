package domain

type PermissionsInterface interface {
	FindAll() ([]Permissions, error)
}
