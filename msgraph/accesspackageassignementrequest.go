package msgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/manicminer/hamilton/odata"
)

type AccessPackageAssignementRequestClient struct {
	BaseClient Client
}

func NewAccessPackageAssignmentRequestClient(tenantId string) *AccessPackageAssignementRequestClient {
	return &AccessPackageAssignementRequestClient{
		BaseClient: NewClient(Version10, tenantId),
	}
}

// List will list all access package assignment requests
func (c *AccessPackageAssignementRequestClient) List(ctx context.Context, query odata.Query) (*[]AccessPackageAssignmentRequest, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		DisablePaging:    query.Top > 0,
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      "/identityGovernance/entitlementManagement/assignmentRequests",
			Params:      query.Values(),
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("AccessPackageAssignementRequestClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var data struct {
		AccessPacakgeAssignmentRequest []AccessPackageAssignmentRequest `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &data.AccessPacakgeAssignmentRequest, status, nil
}

// Get will get an Access Package request
func (c *AccessPackageAssignementRequestClient) Get(ctx context.Context, id string) (*AccessPackageAssignmentRequest, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
		ValidStatusCodes:       []int{http.StatusOK},
		Uri: Uri{
			Entity:      fmt.Sprintf("/identityGovernance/entitlementManagement/assignmentRequests/%s", id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("AccessPackageAssignementRequestClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var accessPackageAssignmentRequest AccessPackageAssignmentRequest
	if err := json.Unmarshal(respBody, &accessPackageAssignmentRequest); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &accessPackageAssignmentRequest, status, nil
}

// Create will create an access package request
func (c *AccessPackageAssignementRequestClient) Create(ctx context.Context, accessPackageAssignementRequest AccessPackageAssignmentRequest) (*AccessPackageAssignmentRequest, int, error) {
	var status int
	body, err := json.Marshal(accessPackageAssignementRequest)
	if err != nil {
		return nil, status, fmt.Errorf("json.Marshal(): %v", err)
	}

	resp, status, _, err := c.BaseClient.Post(ctx, PostHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      "/identityGovernance/entitlementManagement/assignmentRequests",
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("AccessPackageAssignementRequestClient.BaseClient.Post(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var newAccessPackageAssignmentRequest AccessPackageAssignmentRequest
	if err := json.Unmarshal(respBody, &newAccessPackageAssignmentRequest); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &newAccessPackageAssignmentRequest, status, nil
}

// Delete will delete an access package request
func (c *AccessPackageAssignementRequestClient) Delete(ctx context.Context, id string) (int, error) {
	_, status, _, err := c.BaseClient.Delete(ctx, DeleteHttpRequestInput{
		ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
		ValidStatusCodes:       []int{http.StatusNoContent},
		Uri: Uri{
			Entity:      fmt.Sprintf("/identityGovernance/entitlementManagement/assignmentRequests/%s", id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return status, fmt.Errorf("AccessPackageAssignmentPolicyClient.BaseClient.Delete(): %v", err)
	}

	return status, nil

}

// Cancel will cancel a request is in a cancellable state
func (c *AccessPackageAssignementRequestClient) Cancel(ctx context.Context, id string) (int, error) {
	var status int

	_, status, _, err := c.BaseClient.Post(ctx, PostHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      fmt.Sprintf("/identityGovernance/entitlementManagement/assignmentRequests/%s/cancel", id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return status, fmt.Errorf("AccessPackageAssignementRequestClient.BaseClient.Post(): %v", err)
	}

	return status, nil
}

// Reprocess re-processes an access package assignment request
func (c *AccessPackageAssignementRequestClient) Reprocess(ctx context.Context, id string) (int, error) {
	var status int

	_, status, _, err := c.BaseClient.Post(ctx, PostHttpRequestInput{
		ValidStatusCodes: []int{http.StatusAccepted},
		Uri: Uri{
			Entity:      fmt.Sprintf("/identityGovernance/entitlementManagement/assignmentRequests/%s/reprocess", id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return status, fmt.Errorf("AccessPackageAssignementRequestClient.BaseClient.Post(): %v", err)
	}

	return status, nil
}
