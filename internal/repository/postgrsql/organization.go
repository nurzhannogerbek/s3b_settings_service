package postgrsql

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/common"

	"github.com/jmoiron/sqlx"
)

// OrganizationRepository
// Contains information about organization repository.
type OrganizationRepository struct {
	db *sqlx.DB
}

// NewOrganizationRepository
// Creates new OrganizationRepository.
func NewOrganizationRepository(db *sqlx.DB) *OrganizationRepository {
	return &OrganizationRepository{
		db: db,
	}
}

// Create
func (or *OrganizationRepository) Create(organization *common.OrganizationCreateInput) (*common.Organization, error) {
	return nil, nil
}

// CreateDepartment
func (or OrganizationRepository) CreateDepartment(department *common.OrganizationCreateDepartmentInput) (*common.Organization, error) {
	return nil, nil
}

// Delete
func (or *OrganizationRepository) Delete(organizationsIDs *[]string) error {
	return nil
}

// GetByID
func (or *OrganizationRepository) GetByID(organizationID *string) (*common.Organization, error) {
	return nil, nil
}

// Update
func (or *OrganizationRepository) Update(organization *common.OrganizationUpdateInput) (*common.Organization, error) {
	return nil, nil
}

// RestoreDeleted
func (or *OrganizationRepository) RestoreDeleted(organizationsIDs *[]string) error {
	return nil
}

// GetByIDDepartments
func (or *OrganizationRepository) GetByIDDepartments(parentOrganizationID *string) (*[]common.Organization, error) {
	return nil, nil
}

// GetAllDepartments
func (or *OrganizationRepository) GetAllDepartments(rootOrganizationID *string) (*[]common.Organization, error) {
	var organizations []common.Organization
	err := or.db.Select(&organizations, `
		select
			organization_id,
			organization_name,
			parent_organization_id,
			parent_organization_name,
			root_organization_id,
			root_organization_name,
			organization_level,
			parent_organization_level,
			root_organization_level,
			tree_organization_id,
			tree_organization_name
		from
			organizations
		where
			root_organization_id = $1
			and entry_deleted_date_time is null;`, *rootOrganizationID)
	if err != nil {
		return nil, err
	}

	return &organizations, nil
}

// GetByIDArchived
func (or *OrganizationRepository) GetByIDArchived(parentOrganizationID *string) (*[]common.Organization, error) {
	return nil, nil
}

// GetAllArchived
func (or OrganizationRepository) GetAllArchived(rootOrganizationID *string) (*[]common.Organization, error) {
	return nil, nil
}
