package gocd

import (
	"context"
)

// RoleService describes Actions which can be performed on roles
type RoleService service

// Role represents a type of agent/actor who can access resources perform operations
type Role struct {
	Name       string              `json:"name"`
	Type       string              `json:"type"`
	Attributes *RoleAttributesGoCD `json:"attributes"`
}

// RoleAttributesGoCD are attributes describing a role, in this cae, which users are present in the role.
type RoleAttributesGoCD struct {
	Users        []string                     `json:"users"`
	AuthConfigId string                       `json:"auth_config_id"`
	Properties   []*PluginConfigurationKVPair `json:"properties"`
}

// RoleCollection is a collection of roles
type RoleCollection struct {
	Links *HALLinks `json:"_links,omitempty"`
	Embedded *struct {
		Roles []*Role `json:"roles"`
	} `json:"_embedded,omitempty"`
}

// Create a role
func (rs *RoleService) Create(ctx context.Context, role *Role) (r *Role, resp *APIResponse, err error) {
	r = &Role{}
	_, resp, err = rs.client.postAction(ctx, &APIClientRequest{
		APIVersion:   apiV1,
		Path:         "admin/security/roles",
		RequestBody:  role,
		ResponseBody: r,
	})

	return
}

// List roles available
func (rs *RoleService) List(ctx context.Context, roleType *string) (r []*Role, resp *APIResponse, err error) {
	rc := RoleCollection{}

	if roleType != nil {
		u := rs.client.BaseURL
		q := u.Query()
		q.Set("type", *roleType)
		u.RawQuery = q.Encode()
	}

	_, resp, err = rs.client.getAction(ctx, &APIClientRequest{
		APIVersion:   apiV1,
		Path:         "admin/security/roles",
		ResponseBody: &rc,
	})

	return rc.Embedded.Roles, resp, err
}
