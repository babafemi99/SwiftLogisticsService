package riderEntity

type Rider struct {
	RiderId            string       `json:"rider_id"`
	FirstName          string       `json:"first_name"`
	LastName           string       `json:"last_name"`
	Phone              string       `json:"phone"`
	Email              string       `json:"email"`
	PhoneNumber        string       `json:"phone_number"`
	Password           string       `json:"password"`
	ProfilePicture     string       `json:"profile_picture"`
	DOB                string       `json:"dob"`
	Gender             string       `json:"gender"`
	MaritalStatus      string       `json:"marital_status"`
	EducationLevel     string       `json:"education_level"`
	ResidentialAddress string       `json:"residential_address"`
	Guarantor          []*Guarantor `json:"guarantor"`
	DriverLicense      string       `json:"driver_license"`
	Passport           string       `json:"passport"`
	AccountStatus      string       `json:"account_status"`
	WorkStatus         string       `json:"work_status"`
	Salary             string       `json:"salary"`
	BankName           string       `json:"bank_name"`
	BankAccountNo      string       `json:"bank_account_no"`
	Bonus              string       `json:"bonus"`
	Rating             int          `json:"rating"`
	DateCreated        string       `json:"date_created"`
	DateUpdated        string       `json:"date_updated"`
	DateJoined         string       `json:"date_joined"`
}

type AdminEditReq struct {
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Phone              string `json:"phone"`
	Email              string `json:"email"`
	PhoneNumber        string `json:"phone_number"`
	Password           string `json:"password"`
	ProfilePicture     string `json:"profile_picture"`
	DOB                string `json:"dob"`
	Gender             string `json:"gender"`
	MaritalStatus      string `json:"marital_status"`
	EducationLevel     string `json:"education_level"`
	ResidentialAddress string `json:"residential_address"`
	DriverLicense      string `json:"driver_license"`
	Passport           string `json:"passport"`
	AccountStatus      string `json:"account_status"`
	WorkStatus         string `json:"work_status"`
	Salary             string `json:"salary"`
	BankName           string `json:"bank_name"`
	BankAccountNo      string `json:"bank_account_no"`
	Bonus              string `json:"bonus"`
	Rating             int    `json:"rating"`
	DateUpdated        string `json:"date_updated"`
}

type AdminEditRes struct {
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Phone              string `json:"phone"`
	Email              string `json:"email"`
	PhoneNumber        string `json:"phone_number"`
	Password           string `json:"password"`
	ProfilePicture     string `json:"profile_picture"`
	DOB                string `json:"dob"`
	Gender             string `json:"gender"`
	MaritalStatus      string `json:"marital_status"`
	EducationLevel     string `json:"education_level"`
	ResidentialAddress string `json:"residential_address"`
	DriverLicense      string `json:"driver_license"`
	Passport           string `json:"passport"`
	AccountStatus      string `json:"account_status"`
	WorkStatus         string `json:"work_status"`
	Salary             string `json:"salary"`
	BankName           string `json:"bank_name"`
	BankAccountNo      string `json:"bank_account_no"`
	Bonus              string `json:"bonus"`
	Rating             int    `json:"rating"`
	DateUpdated        string `json:"date_updated"`
}

type UpdateRiderReqAdmin struct {
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Email              string `json:"email"`
	PhoneNumber        string `json:"phone_number"`
	ProfilePicture     string `json:"profile_picture"`
	MaritalStatus      string `json:"marital_status"`
	EducationLevel     string `json:"education_level"`
	ResidentialAddress string `json:"residential_address"`
	Salary             int    `json:"salary"`
	BankName           string `json:"bank_name"`
	BankAccountNo      string `json:"bank_account_no"`
	Bonus              int    `json:"bonus"`
	Rating             int    `json:"rating"`
	DateUpdated        string `json:"date_updated"`
}

type CreateRiderReq struct {
	RiderId            string       `json:"rider_id"`
	FirstName          string       `json:"first_name"`
	LastName           string       `json:"last_name"`
	Email              string       `json:"email"`
	Password           string       `json:"password"`
	PhoneNumber        string       `json:"phone_number"`
	DOB                string       `json:"dob"`
	Gender             string       `json:"gender"`
	MaritalStatus      string       `json:"marital_status"`
	EducationLevel     string       `json:"education_level"`
	ResidentialAddress string       `json:"residential_address"`
	Guarantor          []*Guarantor `json:"guarantor"`
	DriverLicense      string       `json:"driver_license"`
	Passport           string       `json:"passport"`
	ProfilePicture     string       `json:"profile_picture"`
	AccountStatus      string       `json:"account_status"`
	DateCreated        string       `json:"date_created"`
}

type CreateRiderRes struct {
	RiderId        string `json:"rider_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	Password       string `json:"password"`
	ProfilePicture string `json:"profile_picture"`
	AccountStatus  string `json:"account_status"`
}

type Guarantor struct {
	GuarantorId                 string `json:"guarantor_id"`
	RiderId                     string `json:"rider_id"`
	GuarantorFirstName          string `json:"guarantor_first_name"`
	GuarantorLastName           string `json:"guarantor_last_name"`
	GuarantorPhone              string `json:"guarantor_phone"`
	GuarantorEmail              string `json:"guarantor_email"`
	GuarantorResidentialAddress string `json:"guarantor_residential_address"`
	GuarantorOfficeAddress      string `json:"guarantor_office_address"`
	GuarantorJob                string `json:"guarantor_job"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateRiderReq struct {
	RiderId        string `json:"rider_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	ProfilePicture string `json:"profile_picture"`
	DateUpdated    string `json:"date_updated"`
}

type UpdateRiderRes struct {
	RiderId        string `json:"rider_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	ProfilePicture string `json:"profile_picture"`
}

type DeleteUser struct {
	RiderId     string `json:"rider_id"`
	GuarantorId string `json:"guarantor_id"`
}

type ResetPassword struct {
	Id       string `json:"id" validate:"required,id"`
	Password string `json:"password" validate:"required"`
}

type ChangePassword struct {
	RiderId     string `json:"-"`
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}
