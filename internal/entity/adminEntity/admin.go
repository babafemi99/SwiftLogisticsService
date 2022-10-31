package adminEntity

// ADMIN AUTHENTICATION

type Admin struct {
	Id             string `json:"_id" bson:"_id"`
	FirstName      string `json:"first_name" bson:"first_name" validate:"required"`
	LastName       string `json:"last_name" bson:"last_name" validate:"required"`
	Email          string `json:"email" bson:"email" validate:"required,email"`
	PhoneNumber    string `json:"phone_number" bson:"phone_number" validate:"required,e164"`
	Gender         string `json:"gender" bson:"gender"`
	Position       string `json:"position"`
	Password       string `json:"password" bson:"password" validate:"required"`
	ProfilePicture string `json:"profile_picture" bson:"profile_picture"`
	AccessLevel    string `json:"access_level" bson:"access_level" validate:"required,alphanum"`
	CreatedAt      string `json:"created_at" bson:"created_at"`
	UpdatedAt      string `json:"updated_at" bson:"updated_at"`
	Role           Role   `json:"role" bson:"role" validate:"required"`
}

type AdminAccess struct {
	Id             string `json:"id" bson:"_id"`
	FirstName      string `json:"first_name" bson:"first_name" validate:"required"`
	LastName       string `json:"last_name" bson:"last_name" validate:"required"`
	Email          string `json:"email" bson:"email" validate:"required,email"`
	PhoneNumber    string `json:"phone_number" bson:"phone_number" validate:"required,e164"`
	Gender         string `json:"gender" bson:"gender"`
	Position       string `json:"position"`
	ProfilePicture string `json:"profile_picture" bson:"profile_picture"`
	AccessLevel    string `json:"access_level" bson:"access_level" validate:"required,alphanum"`
	UpdatedAt      string `json:"updated_at" bson:"updated_at"`
	Role           Role   `json:"role" bson:"role" validate:"required"`
}

type AdminCreated struct {
	FirstName string `json:"first_name" bson:"first_name" validate:"required"`
	LastName  string `json:"last_name" bson:"last_name" validate:"required"`
	Email     string `json:"email" bson:"email" validate:"required,email"`
	Password  string `json:"password" bson:"password" validate:"required"`
}

type AdminList struct {
	FirstName      string `json:"first_name" bson:"first_name" validate:"required"`
	LastName       string `json:"last_name" bson:"last_name" validate:"required"`
	Email          string `json:"email" bson:"email" validate:"required,email"`
	ProfilePicture string `json:"profile_picture" bson:"profile_picture"`
}

type Role struct {
	RoleName        string          `json:"role_name" bson:"role_name" validate:"required"`
	RoleDescription string          `json:"role_description" bson:"role_description"`
	RolePermissions RolePermissions `json:"role_permissions" bson:"role_permissions" validate:"required"`
}

type RolePermissions struct {
	All                    bool `json:"all,omitempty" bson:"all,omitempty"`
	ViewDashboard          bool `json:"view_dashboard,omitempty" bson:"view_dashboard,omitempty"`
	ViewOrders             bool `json:"view_orders,omitempty" bson:"view_orders,omitempty"`
	ViewSales              bool `json:"view_sales,omitempty" bson:"view_sales,omitempty"`
	ViewVisitors           bool `json:"view_visitors,omitempty" bson:"view_visitors,omitempty"`
	ViewActivityLog        bool `json:"view_activity_log,omitempty" bson:"view_activity_log,omitempty"`
	ViewSalesDistributions bool `json:"view_sales_distributions,omitempty" bson:"view_sales_distributions,omitempty"`
	ViewExpense            bool `json:"view_expense,omitempty" bson:"view_expense,omitempty"`
	ViewShipment           bool `json:"view_shipment,omitempty" bson:"view_shipment,omitempty"`
	ViewMessages           bool `json:"view_messages,omitempty" bson:"view_messages,omitempty"`
	ViewRiders             bool `json:"view_riders,omitempty" bson:"view_riders,omitempty"`
	ViewTodo               bool `json:"view_todo,omitempty" bson:"view_todo,omitempty"`
}

type AdminLoginReq struct {
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required,alphanum"`
}

type ChangePasswordReq struct {
	Email       string `json:"email" bson:"email" validate:"required,email"`
	OldPassword string `json:"old_password" bson:"old_password" validate:"required,alphanum"`
	NewPassword string `json:"new_password" bson:"new_password" validate:"required,alphanum"`
}

//////ADMIN SETTINGS///////

type Availability struct {
	Status string `json:"status" bson:"status" validate:"required,oneof=ACTIVE INACTIVE"`
	Start  string `json:"start" bson:"start" validate:"required_if=Status ACTIVE"`
	Stop   string `json:"stop" bson:"stop" validate:"required_if=Status ACTIVE"`
}

type DPR struct {
	Status     string `json:"status" bson:"status"`
	Kilometers int    `json:"kilometers" bson:"kilometers"`
	Seconds    int    `json:"Seconds" bson:"Seconds"`
}

type Category struct {
	Status   string `json:"status" bson:"status" validate:"required,oneof=ACTIVE INACTIVE"`
	Category string `json:"category" bson:"category" validate:"required"`
}

type AdminSettings struct {
	Id                    string       `json:"id" bson:"_id"`
	Status                string       `json:"status" bson:"status" validate:"required,oneof=ACTIVE INACTIVE""`
	Name                  string       `json:"name" bson:"name" validate:"required"`
	NumberOfBikes         int          `json:"number_of_bikes" bson:"number_of_bikes" validate:"required,min=0,max=10"`
	AssigningOrder        string       `json:"assigning_order" bson:"assigning_order" validate:"oneof= AUTO MANUAL"`
	BikeAvailability      Availability `json:"bike_availability" bson:"bike_availability"`
	DeliveryPricingRate   DPR          `json:"delivery_pricing_rate" bson:"delivery_pricing_rate"`
	OrderRejectionReasons []Category   `json:"order_rejection_reasons" bson:"order_rejection_reasons"`
	BusinessCategory      []Category   `json:"business_category" bson:"business_category"`
	TicketCategory        []Category   `json:"ticket_category" bson:"ticket_category"`
	Expenses              []Category   `json:"expenses" bson:"expenses"`
}

type Todo struct {
	Id      string `json:"id" bson:"_id"`
	Starred bool   `json:"starred" bson:"starred"`
	Status  string `json:"status" bson:"status"`
	Title   string `json:"title" bson:"title" validate:"required"`
	Body    string `json:"body" bson:"body"`
}

// ADMIN BIKE SETTINGS

type AdminViewBike struct {
	BikeId       string  `json:"bike_id"`
	NameOfBike   string  `json:"name_of_bike"`
	Colour       string  `json:"colour"`
	EngineNumber string  `json:"engine_number"`
	PlateNumber  string  `json:"plate_number"`
	Rider        *string `json:"rider"`
}

type BikeHistory struct {
	BikeId     string `json:"bike_id"`
	RidersName string `json:"riders_name"`
	Duration   string `json:"duration"`
}

type AdminViewBikeDetails struct {
	BikeId            string         `json:"bike_id"`
	BikeName          string         `json:"bike_name"`
	Colour            string         `json:"colour"`
	EngineNumber      string         `json:"engine_number"`
	PlateNumber       string         `json:"plate_number"`
	Picture           string         `json:"picture"`
	DatePurchased     string         `json:"date_purchased"`
	Duration          string         `json:"duration"`
	NextServicingDate *string        `json:"next_servicing_date"`
	AllocatedDriver   *string        `json:"allocated_driver"`
	History           []*BikeHistory `json:"history"`
}

type AdminViewRider struct {
	RiderId         string  `json:"rider_id"`
	BankAccountNo   *string `json:"bank_account_no"`
	BankName        *string `json:"bank_name"`
	DriversLicense  string  `json:"drivers_license"`
	Passport        string  `json:"passport"`
	Salary          *int    `json:"salary"`
	Bonus           *int    `json:"bonus"`
	NextServiceDate *string `json:"next_service_date"`
	Rating          *int    `json:"rating"`
	NumberOfOrders  *string `json:"number_of_orders"`
	DateJoined      *string `json:"date_joined"`
	PlateNumber     string  `json:"plate_number"`
	Color           string  `json:"color"`
	EngineNumber    string  `json:"engine_number"`
	WorkStatus      string  `json:"work_status"`
}

type AdminViewAllRider struct {
	RiderId        string `json:"rider_id"`
	Name           string `json:"name"`
	NumberOfOrders *int32 `json:"number_of_orders"`
	Ratings        *int   `json:"ratings"`
	Bonus          *int   `json:"bonus"`
	Salary         *int   `json:"salary"`
}

type AdminViewRiderApplication struct {
	RiderId     string `json:"rider_id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type AdminViewRiderApplicationById struct {
	Id                 string       `json:"id"`
	Name               string       `json:"name"`
	Email              string       `json:"email"`
	PhoneNumber        string       `json:"phone_number"`
	DateOfBirth        string       `json:"dob"`
	Gender             string       `json:"gender"`
	MaritalStatus      string       `json:"marital_status"`
	EducationalLevel   string       `json:"educational_level"`
	ResidentialAddress string       `json:"residential_address"`
	DriversLicense     string       `json:"drivers_license"`
	Passport           string       `json:"passport"`
	GuarantorDetails   []*Guarantor `json:"guarantor_details"`
}

type Guarantor struct {
	GuarantorName               string  `json:"guarantor_name"`
	GuarantorEmail              string  `json:"guarantor_email"`
	GuarantorPhone              string  `json:"guarantor_phone"`
	GuarantorResidentialAddress string  `json:"guarantor_residential_address"`
	GuarantorOfficeAddress      string  `json:"guarantor_office_address"`
	GuarantorJob                string  `json:"guarantor_job"`
	GuarantorIdCard             *string `json:"guarantor_id_card"`
}

type AcceptRiderApplication struct {
	AccountStatus string `json:"account_status"`
	Salary        int    `json:"salary"`
	Bonus         int    `json:"bonus"`
	BankName      string `json:"bank_name"`
	BankAccountNo string `json:"bank_account_no"`
	DateJoined    string `json:"date_joined"`
}
