package repository

type Organization interface {
	Create() error
	Update() error
	Get() error
	GetByID() error
	GetArchived() error
	Delete() error
	RestoreDeleted() error
	AssignInternalUser() error
}