// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// API for the Resource Manager service.
// Use this API to install, configure, and manage resources via the "infrastructure-as-code" model.
// For more information, see
// Overview of Resource Manager (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm).
//

package resourcemanager

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v43/common"
)

// UpdateGithubAccessTokenConfigurationSourceProviderDetails The details for updating a configuration source provider of the type `GITHUB_ACCESS_TOKEN`.
// This type corresponds to a configuration source provider in GitHub that is authenticated with a personal access token.
type UpdateGithubAccessTokenConfigurationSourceProviderDetails struct {

	// Human-readable name of the configuration source provider. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the configuration source provider. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags associated with the resource. Each tag is a key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The GitHub service endpoint.
	// Example: `https://github.com/`
	ApiEndpoint *string `mandatory:"false" json:"apiEndpoint"`

	// The personal access token to be configured on the GitHub repository.
	AccessToken *string `mandatory:"false" json:"accessToken"`
}

//GetDisplayName returns DisplayName
func (m UpdateGithubAccessTokenConfigurationSourceProviderDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m UpdateGithubAccessTokenConfigurationSourceProviderDetails) GetDescription() *string {
	return m.Description
}

//GetFreeformTags returns FreeformTags
func (m UpdateGithubAccessTokenConfigurationSourceProviderDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateGithubAccessTokenConfigurationSourceProviderDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateGithubAccessTokenConfigurationSourceProviderDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateGithubAccessTokenConfigurationSourceProviderDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateGithubAccessTokenConfigurationSourceProviderDetails UpdateGithubAccessTokenConfigurationSourceProviderDetails
	s := struct {
		DiscriminatorParam string `json:"configSourceProviderType"`
		MarshalTypeUpdateGithubAccessTokenConfigurationSourceProviderDetails
	}{
		"GITHUB_ACCESS_TOKEN",
		(MarshalTypeUpdateGithubAccessTokenConfigurationSourceProviderDetails)(m),
	}

	return json.Marshal(&s)
}
