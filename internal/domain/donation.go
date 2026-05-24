package domain

import (
	"time"

	"github.com/google/uuid"
)

type DonationHistory struct {
	ID               uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	DonorID          uuid.UUID  `json:"donor_id" gorm:"not null"`
	RequestID        *uuid.UUID `json:"request_id,omitempty"`
	DonationDate     time.Time  `json:"donation_date" gorm:"not null"`
	PMILocation      string     `json:"pmi_location"`
	BloodPressure    string     `json:"blood_pressure"`
	Hemoglobin       float64    `json:"hemoglobin"`
	Weight           float64    `json:"weight"`
	VolumeMl         int        `json:"volume_ml" gorm:"default:350"`
	IsEligible       *bool      `json:"is_eligible,omitempty"`
	DisqualifyReason *string    `json:"disqualify_reason,omitempty"`
	Status           string     `json:"status" gorm:"default:CHECKED_IN"`
	AdminID          *uuid.UUID `json:"admin_id,omitempty"`
	CreatedAt        time.Time  `json:"created_at"`
	// Joined field (tidak di DB)
	DonorName string `json:"donor_name,omitempty" gorm:"-"`
}

type DonationRepository interface {
	Create(d *DonationHistory) error
	FindByID(id uuid.UUID) (*DonationHistory, error)
	Complete(id uuid.UUID) error
	ListByDonor(donorID uuid.UUID) ([]*DonationHistory, error)
	ListByDonorAdmin(donorID uuid.UUID) ([]*DonationHistory, error)
}

type DonationUsecase interface {
	CheckIn(req CheckInRequest, adminID uuid.UUID) (*DonationHistory, error)
	Complete(donationID uuid.UUID, adminID uuid.UUID) error
	GetDonorHistory(donorID uuid.UUID) ([]*DonationHistory, error)
}

// Request DTOs
type CheckInRequest struct {
	DonorUUID     string     `json:"donor_uuid" validate:"required"`
	RequestID     *uuid.UUID `json:"request_id"`
	BloodPressure string     `json:"blood_pressure" validate:"required"`
	Hemoglobin    float64    `json:"hemoglobin" validate:"required,min=5,max=25"`
	Weight        float64    `json:"weight" validate:"required,min=20,max=200"`
	PMILocation   string     `json:"pmi_location"`
}

// Hasil screening medis
type MedicalScreeningResult struct {
	IsEligible       bool
	DisqualifyReason string
}
