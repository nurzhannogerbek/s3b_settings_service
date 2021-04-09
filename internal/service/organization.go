package service

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/common"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/repository"
	"bitbucket.org/3beep-workspace/3beep_settings_service/pkg/tool/uuid"
	"strings"
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
	if organization.OrganizationName == nil {
		return nil, nil
	}

	if err := organization.Validate(); err != nil {
		return nil, err
	}

	newOrganization, err := os.repository.CreateOrganization(*organization)
	if err != nil {
		return nil, err
	}

	return newOrganization, nil
}

// CreateOrganizationDepartment
func (os OrganizationService) CreateOrganizationDepartment(organization *common.OrganizationDepartmentCreateInput) (*common.Organization, error) {
	if organization.OrganizationName == nil || organization.ParentOrganizationID == nil {
		return nil, nil
	}

	if err := uuid.Validate(organization.ParentOrganizationID); err != nil {
		return nil, err
	}

	department, err := os.repository.CreateOrganizationDepartment(*organization)

	if err != nil {
		return nil, err
	}

	return department, nil
}

// DeleteOrganizations
func (os *OrganizationService) DeleteOrganizations(organizationsIDs *[]string) error {
	if organizationsIDs == nil {
		return nil
	}

	if err := uuid.ConcurrencyValidate(organizationsIDs); err != nil {
		return err
	}

	err := os.repository.DeleteOrganizations(*organizationsIDs)
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

	organization, err := os.repository.GetOrganizationByID(*organizationID)
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

	organizations, err := os.repository.GetOrganizationsByIDs(*organizationsIDs)
	if err != nil {
		return nil, err
	}

	return organizations, err
}

// UpdateOrganization
func (os *OrganizationService) UpdateOrganization(organization *common.OrganizationNameUpdateInput) (*common.Organization, error) {
	if err := organization.Validate(); err != nil {
		return nil, err
	}

	organizationInformation, err := os.repository.GetOrganizationByID(*organization.OrganizationID)
	if err != nil {
		return nil, err
	}

	if err := os.repository.UpdateOrganizationName(*organization.OrganizationID, *organization.OrganizationName); err != nil {
		return nil, err
	}

	organizations, err := os.repository.GetUpdateTreeOrganizations(*organization.OrganizationID, *organization.OrganizationName)
	if err != nil {
		return nil, err
	}

	organizationLevel := *organizationInformation.OrganizationLevel

	for _, org := range *organizations {
		a := strings.Split(*org.TreeOrganizationName, "\\")
		a[organizationLevel] = strings.ReplaceAll(a[organizationLevel], a[organizationLevel], *organization.OrganizationName)
		newTreeOrgName := strings.Join(a, "\\")
		if err := os.repository.UpdateTreeOrganizationName(*org.OrganizationID, newTreeOrgName); err != nil {
			return nil, err
		}
	}

	a := strings.Split(*organizationInformation.TreeOrganizationName, "\\")
	a[organizationLevel] = strings.ReplaceAll(a[organizationLevel], a[organizationLevel], *organization.OrganizationName)
	newTreeOrgName := strings.Join(a, "\\")
	organizationInformation.OrganizationName = organization.OrganizationName
	organizationInformation.TreeOrganizationName = &newTreeOrgName

	return organizationInformation, nil
}

// RestoreDeletedOrganizations
func (os *OrganizationService) RestoreDeletedOrganizations(organizationsIDs *[]string) error {
	if organizationsIDs == nil {
		return nil
	}

	if err := uuid.ConcurrencyValidate(organizationsIDs); err != nil {
		return err
	}

	err := os.repository.RestoreDeletedOrganizations(*organizationsIDs)
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

	organizationDepartments, err := os.repository.GetOrganizationDepartmentsByID(*parentOrganizationID)
	if err != nil {
		return nil, err
	}

	return organizationDepartments, err
}

// GetAllOrganizationDepartments
func (os *OrganizationService) GetAllOrganizationDepartments(rootOrganizationID *string) ([]common.Organization, error) {
	if rootOrganizationID == nil {
		return nil, nil
	}

	if err := uuid.Validate(rootOrganizationID); err != nil {
		return nil, err
	}

	organizationDepartments, err := os.repository.GetAllOrganizationDepartments(*rootOrganizationID)
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

	archivedOrganizationDepartments, err := os.repository.GetArchivedOrganizationDepartmentsByID(*parentOrganizationID)
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

	archivedOrganizationDepartments, err := os.repository.GetAllArchivedOrganizationDepartments(*rootOrganizationID)
	if err != nil {
		return nil, err
	}

	return archivedOrganizationDepartments, err
}