package service

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/common"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/repository"
)

// OrganizationService
// Contains all dependencies for organization service.
type OrganizationService struct {
	repository repository.Organization
}

// NewOrganizationService
// Creates new organization service (Organization).
func NewOrganizationService(os repository.Organization) *OrganizationService {
	return &OrganizationService{
		repository: os,
	}
}

// Create
func (os *OrganizationService) Create(organization *common.OrganizationCreateInput) (*common.Organization, error) {
	return nil, nil
}

// CreateDepartment
func (os OrganizationService) CreateDepartment(department *common.OrganizationCreateDepartmentInput) (*common.Organization, error) {
	return nil, nil
}

// Delete
func (os *OrganizationService) Delete(organizationsIDs *[]string) error {
	return nil
}

// GetByID
func (os *OrganizationService) GetByID(organizationID *string) (*common.Organization, error) {
	return nil, nil
}

// Update
func (os *OrganizationService) Update(organization *common.OrganizationUpdateInput) (*common.Organization, error) {
	return nil, nil
}

// RestoreDeleted
func (os *OrganizationService) RestoreDeleted(organizationsIDs *[]string) error {
	return nil
}

// GetByIDDepartments
func (os *OrganizationService) GetByIDDepartments(parentOrganizationID *string) (*[]common.Organization, error) {
	return nil, nil
}

// GetAllDepartments
func (os *OrganizationService) GetAllDepartments(rootOrganizationID *string) (*[]common.Organization, error) {
	return nil, nil
}

// GetByIDArchived
func (os *OrganizationService) GetByIDArchived(parentOrganizationID *string) (*[]common.Organization, error) {
	return nil, nil
}

// GetAllArchived
func (os OrganizationService) GetAllArchived(rootOrganizationID *string) (*[]common.Organization, error) {
	return nil, nil
}