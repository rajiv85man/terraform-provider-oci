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

// PlanJobOperationDetails Job details that are specific to plan operations.
type PlanJobOperationDetails struct {
}

func (m PlanJobOperationDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m PlanJobOperationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePlanJobOperationDetails PlanJobOperationDetails
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypePlanJobOperationDetails
	}{
		"PLAN",
		(MarshalTypePlanJobOperationDetails)(m),
	}

	return json.Marshal(&s)
}
