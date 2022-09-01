package riderEntity

import "github.com/google/uuid"

type CreateRiderReq struct {
	RiderId            uuid.UUID  `json:"rider_id"`
	FirstName          string     `json:"first_name"`
	LastName           string     `json:"last_name"`
	Phone              string     `json:"phone"`
	Email              string     `json:"email"`
	Password           string     `json:"password"`
	ProfilePicture     string     `json:"profile_picture"`
	DOB                string     `json:"dob"`
	Gender             string     `json:"gender"`
	MaritalStatus      string     `json:"marital_status"`
	EducationLevel     string     `json:"education_level"`
	ResidentialAddress string     `json:"residential_address"`
	Guarantor          *Guarantor `json:"guarantor"`
	DriverLicense      string     `json:"driver_license"`
	IdentityCard       string     `json:"identity_card"`
	VerificationStatus string     `json:"verification_status"`
	AccountStatus      string     `json:"account_status"`
	CreatedAt          string     `json:"created_at"`
	UpdatedAt          string     `json:"updated_at"`
}

type CreateRiderRes struct {
	RiderId            uuid.UUID `json:"rider_id"`
	FirstName          string    `json:"first_name"`
	LastName           string    `json:"last_name"`
	Phone              string    `json:"phone"`
	Email              string    `json:"email"`
	Password           string    `json:"password"`
	ProfilePicture     string    `json:"profile_picture"`
	VerificationStatus string    `json:"verification_status"`
	AccountStatus      string    `json:"account_status"`
}

type Guarantor struct {
	GuarantorId      uuid.UUID `json:"guarantor_id"`
	RiderId          uuid.UUID `json:"rider_id"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	Phone            string    `json:"phone"`
	Email            string    `json:"email"`
	GuarantorAddress string    `json:"guarantor_address"`
	GuarantorJob     string    `json:"guarantor_job"`
}

type LoginReq struct {
	Credentials string `json:"credentials"`
	Password    string `json:"password"`
}

type UpdateRiderReq struct {
	RiderId        uuid.UUID `json:"rider_id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Phone          string    `json:"phone"`
	Email          string    `json:"email"`
	ProfilePicture string    `json:"profile_picture"`
	UpdatedAt      string    `json:"updated_at"`
}

type UpdateRiderRes struct {
	RiderId        uuid.UUID `json:"rider_id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Phone          string    `json:"phone"`
	Email          string    `json:"email"`
	ProfilePicture string    `json:"profile_picture"`
}

type DeleteUser struct {
	RiderId     uuid.UUID `json:"rider_id"`
	GuarantorId uuid.UUID `json:"guarantor_id"`
}
