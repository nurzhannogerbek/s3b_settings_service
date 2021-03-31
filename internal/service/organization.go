package service

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/common"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/repository"
	"bitbucket.org/3beep-workspace/3beep_settings_service/pkg/tool/uuid"
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

// CreateOrganization
func (os *OrganizationService) CreateOrganization(organization *common.OrganizationCreateInput) (*common.Organization, error) {
	return nil, nil
}

// CreateOrganizationDepartment
func (os OrganizationService) CreateOrganizationDepartment(department *common.OrganizationCreateDepartmentInput) (*common.Organization, error) {
	return nil, nil
}

// DeleteOrganizations
func (os *OrganizationService) DeleteOrganizations(organizationsIDs *[]string) error {
	if organizationsIDs == nil {
		return nil
	}

	if err := uuid.ConcurrencyValidate(organizationsIDs); err != nil {
		return err
	}

	err := os.repository.DeleteOrganizations(organizationsIDs)
	if err != nil {
		return err
	}

	return nil
}

// GetOrganizationByID
func (os *OrganizationService) GetOrganizationByID(organizationID *string) (*common.Organization, error) {
	if organizationID == nil {
		return nil, nil
	}

	if err := uuid.Validate(organizationID); err != nil {
		return nil, err
	}

	organization, err := os.repository.GetOrganizationByID(organizationID)
	if err != nil {
		return nil, err
	}

	return organization, err
}


// GetOrganizationsByIDs
func (os *OrganizationService) GetOrganizationsByIDs(organizationsIDs *[]string) (*[]common.Organization, error) {
	if organizationsIDs == nil {
		return nil, nil
	}

	if err := uuid.ConcurrencyValidate(organizationsIDs); err != nil {
		return nil, err
	}

	organizations, err := os.repository.GetOrganizationsByIDs(organizationsIDs)
	if err != nil {
		return nil, err
	}

	return organizations, err
}

// UpdateOrganization
func (os *OrganizationService) UpdateOrganization(organization *common.OrganizationUpdateInput) (*common.Organization, error) {
	return nil, nil
}

// RestoreDeletedOrganizations
func (os *OrganizationService) RestoreDeletedOrganizations(organizationsIDs *[]string) error {
	if organizationsIDs == nil {
		return nil
	}

	if err := uuid.ConcurrencyValidate(organizationsIDs); err != nil {
		return err
	}

	err := os.repository.RestoreDeletedOrganizations(organizationsIDs)
	if err != nil {
		return err
	}

	return nil
}

// GetOrganizationDepartmentsByID
func (os *OrganizationService) GetOrganizationDepartmentsByID(parentOrganizationID *string) (*[]common.Organization, error) {
	if parentOrganizationID == nil {
		return nil, nil
	}

	if err := uuid.Validate(parentOrganizationID); err != nil {
		return nil, err
	}

	organizationDepartments, err := os.repository.GetOrganizationDepartmentsByID(parentOrganizationID)
	if err != nil {
		return nil, err
	}

	return organizationDepartments, err
}

// GetAllOrganizationDepartments
func (os *OrganizationService) GetAllOrganizationDepartments(rootOrganizationID *string) (*[]common.Organization, error) {
	if rootOrganizationID == nil {
		return nil, nil
	}

	if err := uuid.Validate(rootOrganizationID); err != nil {
		return nil, err
	}

	organizationDepartments, err := os.repository.GetAllOrganizationDepartments(rootOrganizationID)
	if err != nil {
		return nil, err
	}

	return organizationDepartments, err
}

// GetArchivedOrganizationDepartmentsByID
func (os *OrganizationService) GetArchivedOrganizationDepartmentsByID(parentOrganizationID *string) (*[]common.Organization, error) {
	if parentOrganizationID == nil {
		return nil, nil
	}

	if err := uuid.Validate(parentOrganizationID); err != nil {
		return nil, err
	}

	archivedOrganizationDepartments, err := os.repository.GetArchivedOrganizationDepartmentsByID(parentOrganizationID)
	if err != nil {
		return nil, err
	}

	return archivedOrganizationDepartments, err
}

// GetAllArchivedOrganizationDepartments
func (os OrganizationService) GetAllArchivedOrganizationDepartments(rootOrganizationID *string) (*[]common.Organization, error) {
	if rootOrganizationID == nil {
		return nil, nil
	}

	if err := uuid.Validate(rootOrganizationID); err != nil {
		return nil, err
	}

	archivedOrganizationDepartments, err := os.repository.GetAllArchivedOrganizationDepartments(rootOrganizationID)
	if err != nil {
		return nil, err
	}

	return archivedOrganizationDepartments, err
}