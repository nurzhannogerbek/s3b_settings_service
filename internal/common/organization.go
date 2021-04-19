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
	RootOrganizationLevel   *int    `db:"root_organization_level" json:"rootOrganizationLevel"`
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
		if organizationName > 100 || organizationName < 1 {
			return errors.New("organizationName length should be greater or equal to 1 and less or equal to 100 symbols in length")
		}
	}

	return nil
}

// OrganizationCreateDepartmentInput
type OrganizationDepartmentCreateInput struct {
	OrganizationName     *string `db:"organization_name" json:"organizationName"`
	ParentOrganizationID *string `db:"parent_organization_id" json:"parentOrganizationId"`
}

// Validate
func (o *OrganizationDepartmentCreateInput) Validate() error {
	if o.OrganizationName != nil {
		organizationName := utf8.RuneCountInString(*o.OrganizationName)
		if organizationName > 100 || organizationName < 1 {
			return errors.New("organizationName length should be greater or equal to 1 and less or equal to 100 symbols in length")
		}
	}

	if o.ParentOrganizationID != nil {
		if err := uuid.Validate(o.ParentOrganizationID); err != nil {
			return err
		}
	}

	return nil
}

// OrganizationNameUpdateInput
type OrganizationNameUpdateInput struct {
	OrganizationID   *string `db:"organization_id" json:"organizationId"`
	OrganizationName *string `db:"organization_name" json:"organizationName"`
}

// Validate
// Validates OrganizationNameUpdateInput struct
func (oui OrganizationNameUpdateInput) Validate() error {
	if oui.OrganizationID != nil {
		if err := uuid.Validate(oui.OrganizationID); err != nil {
			return err
		}
	} else {
		return errors.New("must have an ID")
	}

	if oui.OrganizationName != nil && (utf8.RuneCountInString(*oui.OrganizationName) > 100 || utf8.RuneCountInString(*oui.OrganizationName) < 1) {
		return errors.New("organizationName length should be greater or equal to 1 and less or equal to 100 symbols in length")
	}

	return nil
}

// TreeOrganizationNameUpdateInput
type TreeOrganizationNameUpdateInput struct {
	OrganizationID       *string `db:"organization_id" json:"organizationId"`
	TreeOrganizationName *string `db:"tree_organization_name" json:"treeOrganizationName"`
}

// Validate
// Validates TreeOrganizationNameUpdateInput struct
func (oui TreeOrganizationNameUpdateInput) Validate() error {
	if oui.OrganizationID != nil {
		if err := uuid.Validate(oui.OrganizationID); err != nil {
			return err
		}
	} else {
		return errors.New("must have an ID")
	}


	return nil
}
