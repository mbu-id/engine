package common

import "encoding/json"

// CreatedBySnapshot is the standard denormalized creator info embedded as JSONB
// in every table that tracks document creation. Populated once from SessionClaims
// at creation time — survives downstream user changes (rename, delete, transfer).
type CreatedBySnapshot struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	CompanyID      string `json:"company_id"`
	BranchID       string `json:"branch_id"`
	BranchRegionID string `json:"branch_region_id"`
	DepartmentID   string `json:"department_id"`
}

// NewCreatedBySnapshot marshals SessionClaims into a *json.RawMessage.
// Returns nil if claims is nil.
func NewCreatedBySnapshot(claims *SessionClaims) *json.RawMessage {
	if claims == nil {
		return nil
	}
	snap := CreatedBySnapshot{
		ID:             claims.UserID,
		Name:           claims.DisplayName,
		Email:          claims.Email,
		CompanyID:      claims.CompanyID,
		BranchID:       claims.BranchID,
		BranchRegionID: claims.BranchRegionID,
		DepartmentID:   claims.DepartmentID,
	}
	b, _ := json.Marshal(snap)
	raw := json.RawMessage(b)
	return &raw
}
