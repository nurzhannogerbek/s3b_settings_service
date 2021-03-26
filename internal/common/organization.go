package common

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/pkg/tool/uuid"
	"errors"
	"unicode/utf8"
)

// Organization
// Contains information about organization
type Organization struct {
	OrganizationID          *string `db:"organization_id" json:"organizationId"`
	OrganizationName        *string `db:"organization_name" json:"organizationName"`
	ParentOrganizationID    *string `db:"parent_organization_id" json:"parentOrganizationId"`
	ParentOrganizationName  *string `db:"parent_organization_name" json:"parentOrganizationName"`
	RootOrganizationID      *string `db:"root_organization_id" json:"rootOrganizationId"`
	RootOrganizationName    *string `db:"root_organization_name" json:"rootOrganizationName"`
	OrganizationLevel       *int    `db:"organization_level" json:"organizationLevel"`
	ParentOrganizationLevel *int    `db:"parent_organization_level" json:"parentOrganizationLevel"`
	TreeOrganizationID      *string `db:"tree_organization_id" json:"treeOrganizationId"`
	TreeOrganizationName    *string `db:"tree_organization_name" json:"treeOrganizationName"`
}

// OrganizationCreateInput
type OrganizationCreateInput struct {
	OrganizationName *string `db:"organization_name" json:"organizationName"`
}

// Validate
func (o *OrganizationCreateInput) Validate() error {
	if o.OrganizationName != nil {
		organizationName := utf8.RuneCountInString(*o.OrganizationName)
		if organizationName > 100 || organizationName < 5 {
			return errors.New("organizationName length should be less than 50 and greater than 5")
		}
	}

	return nil
}

// OrganizationCreateDepartmentInput
type OrganizationCreateDepartmentInput struct {
	OrganizationName       *string `db:"organization_name" json:"organizationName"`
	ParentOrganizationID *string `dbg:"parent_organization_id" json:"parentOrganizationId"`
}

// Validate
func (o *OrganizationCreateDepartmentInput) Validate() error {
	if o.OrganizationName != nil {
		organizationName := utf8.RuneCountInString(*o.OrganizationName)
		if organizationName > 100 || organizationName < 5 {
			return errors.New("organizationName length should be less than 50 and greater than 5")
		}
	}

	if o.ParentOrganizationID != nil {
		if err := uuid.Validate(o.ParentOrganizationID); err != nil {
			return err
		}
	}

	return nil
}

// UpdateOrganization
type OrganizationUpdateInput struct {
	OrganizationID       *string `db:"organization_id" json:"organizationId"`
	OrganizationName     *string `db:"organization_name" json:"organizationName"`
	ParentOrganizationID *string `db:"parent_organization_id" json:"parentOrganizationId"`
}
