package domain

import (
	"time"

	"github.com/google/uuid"
)

type BloodRequest struct {
	ID             uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	HospitalID     *uuid.UUID `json:"hospital_id,omitempty"`
	RequestedBy    string     `json:"requested_by"`
	BloodType      string     `json:"blood_type" gorm:"not null"`
	ProductType    string     `json:"product_type" gorm:"default:PRC"`
	QuantityNeeded int        `json:"quantity_needed" gorm:"not null"`
	UrgencyLevel   string     `json:"urgency_level" gorm:"not null"`
	Notes          *string    `json:"notes,omitempty"`
	BroadcastID    *string    `json:"broadcast_id,omitempty"` // Firestore doc ID
	Status         string     `json:"status" gorm:"default:PENDING"`
	AdminID        *uuid.UUID `json:"admin_id,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	FulfilledAt    *time.Time `json:"fulfilled_at,omitempty"`
}

type EligibleDonor struct {
	ID          uuid.UUID `json:"id"`
	FullName    string    `json:"full_name"`
	Phone       string    `json:"phone"`
	BloodType   string    `json:"blood_type"`
	DeviceToken *string   `json:"-"` // hanya untuk FCM, tidak dikirim ke client
	DistanceKM  float64   `json:"distance_km"`
}

type BloodRequestRepository interface {
	Create(req *BloodRequest) error
	FindByID(id uuid.UUID) (*BloodRequest, error)
	List(filter RequestFilter) ([]*BloodRequest, int64, error)
	UpdateStatus(id uuid.UUID, status string) error
	UpdateBroadcastID(id uuid.UUID, broadcastID string) error
	SetFulfilled(id uuid.UUID) error
	GetEligibleDonors(bloodType string, pmiLat, pmiLng float64) ([]*EligibleDonor,
		error)
}

type BloodRequestUsecase interface {
	Create(req CreateBloodRequestInput) (*BloodRequest, error)
	List(filter RequestFilter) ([]*BloodRequest, int64, error)
	GetEligibleDonors(requestID uuid.UUID) ([]*EligibleDonor, error)
	BroadcastEmergency(requestID uuid.UUID, adminID uuid.UUID) error
	UpdateStatus(id uuid.UUID, status string) error
}

// Request DTOs
type CreateBloodRequestInput struct {
	HospitalID  *uuid.UUID `json:"hospital_id"`
	RequestedBy string     `json:"requested_by" validate:"required"`
	BloodType   string     `json:"blood_type" validate:"required,oneof=A+ A- B+
	B- O+ O- AB+ AB-"`
	ProductType string `json:"product_type" validate:"required,oneof=WB PRC
	FFP THROMBOCYTE"`
	QuantityNeeded int    `json:"quantity_needed" validate:"required,min=1"`
	UrgencyLevel   string `json:"urgency_level"
	validate:"required,oneof=CRITICAL URGENT NORMAL"`
	Notes string `json:"notes"`
}

type RequestFilter struct {
	Status    string
	BloodType string
	Page      int
	Limit     int
}

// Error sentinel
var ErrNotFound = errNotFound("not found")

type errNotFound string

func (e errNotFound) Error() string { return string(e) }
