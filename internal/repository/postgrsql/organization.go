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

// CreateOrganization
func (or *OrganizationRepository) CreateOrganization(organization common.OrganizationCreateInput) (*common.Organization, error) {
	rows, err := or.db.NamedQuery(`
		insert into organizations (
			organization_name,
			parent_organization_id)
		values (
			:organization_name,
			null)
		returning organization_id,
		    organization_name,
			parent_organization_id,
			parent_organization_name,
			root_organization_id,
			root_organization_name,
			organization_level,
			tree_organization_id,
			tree_organization_name;`, organization)
	if err != nil {
		return nil, err
	}

	var newOrganization common.Organization
	for rows.Next() {
		if err := rows.StructScan(&newOrganization); err != nil {
			return nil, err
		}
	}

	return &newOrganization, nil
}

// CreateOrganizationDepartment
func (or OrganizationRepository) CreateOrganizationDepartment(department common.OrganizationCreateDepartmentInput) (*common.Organization, error) {
	rows, err := or.db.NamedQuery(`
		insert into organizations (
			organization_name,
			parent_organization_id)
		values (
			:organization_name,
			:parent_organization_id)
		returning organization_id,
		    organization_name,
			parent_organization_id,
			parent_organization_name,
			root_organization_id,
			root_organization_name,
			organization_level,
			tree_organization_id,
			tree_organization_name;`, department)
	if err != nil {
		return nil, err
	}

	var newDepartment common.Organization
	for rows.Next() {
		if err := rows.Scan(&newDepartment.OrganizationID,
			&newDepartment.OrganizationName,
			&newDepartment.ParentOrganizationID,
			&newDepartment.ParentOrganizationName,
			&newDepartment.RootOrganizationID,
			&newDepartment.RootOrganizationName,
			&newDepartment.RootOrganizationLevel,
			&newDepartment.OrganizationLevel,
			&newDepartment.TreeOrganizationID,
			&newDepartment.TreeOrganizationName); err != nil {
			return nil, err
		}
	}

	return &newDepartment, nil
}

// DeleteOrganizations
func (or *OrganizationRepository) DeleteOrganizations(organizationsIDs []string) error {
	query, args, err := sqlx.In(`
		update
			organizations 
		set
			entry_deleted_date_time = now()
		where
			organization_id in (?);`, organizationsIDs)
	if err != nil {
		return err
	}

	query = or.db.Rebind(query)
	_, err = or.db.Query(query, args...)
	if err != nil {
		return err
	}

	return nil
}

// GetOrganizationByID
func (or *OrganizationRepository) GetOrganizationByID(organizationID string) (*common.Organization, error) {
	var organization common.Organization
	err := or.db.Get(&organization, `
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
			organization_id = $1;`, organizationID)
	if err != nil {
		return nil, err
	}

	return &organization, nil
}

// GetOrganizationsByIDs
// Queries organizations by IDs from database.
func (or *OrganizationRepository) GetOrganizationsByIDs(organizationsIDs []string) (*[]common.Organization, error) {
	query, args, err := sqlx.In(`
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
			organization_id in (?);`, organizationsIDs)
	if err != nil {
		return nil, err
	}

	query = or.db.Rebind(query)
	rows, err := or.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}

	var organizations []common.Organization
	for rows.Next() {
		var organization common.Organization
		if err := rows.Scan(&organization.OrganizationID,
			&organization.OrganizationName,
			&organization.ParentOrganizationID,
			&organization.ParentOrganizationName,
			&organization.RootOrganizationID,
			&organization.RootOrganizationName,
			&organization.OrganizationLevel,
			&organization.RootOrganizationLevel,
			&organization.ParentOrganizationLevel,
			&organization.TreeOrganizationID,
			&organization.TreeOrganizationName); err != nil {
			return nil, err
		}
		organizations = append(organizations, organization)
	}

	return &organizations, nil
}

// UpdateOrganization
func (or *OrganizationRepository) UpdateOrganization(organizationID, organizationName string) error {
	_, err := or.db.Query(`
		update 
		    organizations
		set entry_updated_date_time = now(),
		    organization_name = $1
		where organization_id = $2;`, organizationName, organizationID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateOrganizations
func (or *OrganizationRepository) UpdateTreeOrganizationName(organizationID, treeOrganizationName string) error {
	_, err := or.db.Query(`
		update 
		    organizations
		set entry_updated_date_time = now(),
		    tree_organization_name = $1
		where organization_id = $2;`, treeOrganizationName, organizationID)
	if err != nil {
		return err
	}

	return nil
}

func (or *OrganizationRepository) GetUpdateTreeOrganizations(organizationID, organizationName string) (*[]common.Organization, error) {
	var organizations []common.Organization
	err := or.db.Select(&organizations, `
		select organization_id,
			tree_organization_id
		from organizations
		where root_organization_id = (select root_organization_id 
							  		  from organizations 
							          where organization_id = $1::uuid)
		and tree_organization_name like '%' || '\' || $2::text || '%';`, organizationID, organizationName)
	if err != nil {
		return nil, err
	}

	return &organizations, nil
}

// RestoreDeletedOrganizations
func (or *OrganizationRepository) RestoreDeletedOrganizations(organizationsIDs []string) error {
	query, args, err := sqlx.In(`
		update
			organizations 
		set
			entry_deleted_date_time = null
		where
			organization_id in (?);`, organizationsIDs)
	if err != nil {
		return err
	}

	query = or.db.Rebind(query)
	_, err = or.db.Query(query, args...)
	if err != nil {
		return err
	}

	return nil
}

// GetOrganizationDepartmentsByID
func (or *OrganizationRepository) GetOrganizationDepartmentsByID(parentOrganizationID string) (*[]common.Organization, error) {
	var organizationDepartments []common.Organization
	err := or.db.Select(&organizationDepartments, `
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
			parent_organization_id = $1
		and
			entry_deleted_date_time is null;`, parentOrganizationID)
	if err != nil {
		return nil, err
	}

	return &organizationDepartments, nil
}

// GetAllOrganizationDepartments
func (or *OrganizationRepository) GetAllOrganizationDepartments(rootOrganizationID string) ([]common.Organization, error) {
	var organizationDepartments []common.Organization
	err := or.db.Select(&organizationDepartments, `
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
		and
			entry_deleted_date_time is null;`, rootOrganizationID)
	if err != nil {
		return nil, err
	}

	return organizationDepartments, nil
}

// GetArchivedOrganizationDepartmentsByID
func (or *OrganizationRepository) GetArchivedOrganizationDepartmentsByID(parentOrganizationID string) (*[]common.Organization, error) {
	var archivedOrganizationDepartments []common.Organization
	err := or.db.Select(&archivedOrganizationDepartments, `
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
			parent_organization_id = $1
		and
			entry_deleted_date_time is not null;`, parentOrganizationID)
	if err != nil {
		return nil, err
	}

	return &archivedOrganizationDepartments, nil
}

// GetAllArchivedOrganizationDepartments
func (or OrganizationRepository) GetAllArchivedOrganizationDepartments(rootOrganizationID string) (*[]common.Organization, error) {
	var archivedOrganizationDepartments []common.Organization
	err := or.db.Select(&archivedOrganizationDepartments, `
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
		and
		    entry_deleted_date_time is not null;`, rootOrganizationID)
	if err != nil {
		return nil, err
	}

	return &archivedOrganizationDepartments, nil
}
