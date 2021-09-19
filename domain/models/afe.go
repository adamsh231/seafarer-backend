package models

import "time"

type AFE struct {
	UserID                   string                      `bson:"_id"`
	PersonalInformation      afePersonalInformation      `bson:"personal_information,omitempty" json:"personal_information,omitempty"`
	ContactInformation       afeContactInformation       `bson:"contact_information,omitempty" json:"contact_information,omitempty"`
	DependantInformation     afeDependantInformation     `bson:"dependant_information,omitempty" json:"dependant_information,omitempty"`
	PositionDesired          afePositionDesired          `bson:"position_desired,omitempty" json:"position_desired,omitempty"`
	DocumentationInformation afeDocumentationInformation `bson:"documentation_information,omitempty" json:"documentation_information,omitempty"`
	EmploymentHistory        []afeEmploymentHistory      `bson:"employment_history,omitempty" json:"employment_history,omitempty"`
	Education                []afeEducation              `bson:"education,omitempty" json:"education,omitempty"`
	Languages                []afeLanguages              `bson:"languages,omitempty" json:"languages,omitempty"`
}

type afePersonalInformation struct {
	LastName        string    `bson:"last_name,omitempty" json:"last_name,omitempty"`
	FirstName       string    `bson:"first_name,omitempty" json:"first_name,omitempty"`
	MiddleName      string    `bson:"middle_name,omitempty" json:"middle_name,omitempty"`
	DateOfBirth     time.Time `bson:"date_of_birth,omitempty" json:"date_of_birth,omitempty"`
	BirthPlace      string    `bson:"birth_place,omitempty" json:"birth_place,omitempty"`
	CountryOfBirth  string    `bson:"country_of_birth,omitempty" json:"country_of_birth,omitempty"`
	Nationality     string    `bson:"nationality,omitempty" json:"nationality,omitempty"`
	Gender          string    `bson:"gender,omitempty" json:"gender,omitempty"`
	HairColor       string    `bson:"hair_color,omitempty" json:"hair_color,omitempty"`
	Weight          float64   `bson:"weight,omitempty" json:"weight,omitempty"`
	Height          float64   `bson:"height,omitempty" json:"height,omitempty"`
	HasTattoo       bool      `bson:"has_tattoo,omitempty" json:"has_tattoo,omitempty"`
	IsTattooVisible bool      `bson:"is_tattoo_visible,omitempty" json:"is_tattoo_visible,omitempty"`
}

type afeContactInformation struct {
	Street1     string `bson:"street_1,omitempty" json:"street_1,omitempty"`
	Street2     string `bson:"street_2,omitempty" json:"street_2,omitempty"`
	City        string `bson:"city,omitempty" json:"city,omitempty"`
	State       string `bson:"state,omitempty" json:"state,omitempty"`
	Zip         string `bson:"zip,omitempty" json:"zip,omitempty"`
	Country     string `bson:"country,omitempty" json:"country,omitempty"`
	HomePhone   string `bson:"home_phone,omitempty" json:"home_phone,omitempty"`
	MobilePhone string `bson:"mobile_phone,omitempty" json:"mobile_phone,omitempty"`
	Email       string `bson:"email,omitempty" json:"email,omitempty"`
}

type afeDependantInformation struct {
	MaritalStatus           string                           `bson:"marital_status,omitempty" json:"marital_status,omitempty"`
	NumberOfChildrenUnder18 int                              `bson:"number_of_children_under_18,omitempty" json:"number_of_children_under_18,omitempty"`
	Persons                 []afeEmergencyContactInformation `bson:"persons,omitempty" json:"persons,omitempty"`
}

type afeEmergencyContactInformation struct {
	Relationship string `bson:"relationship,omitempty" json:"relationship,omitempty"`
	LastName     string `bson:"last_name,omitempty" json:"last_name,omitempty"`
	FirstName    string `bson:"first_name,omitempty" json:"first_name,omitempty"`
	MiddleName   string `bson:"middle_name,omitempty" json:"middle_name,omitempty"`
	HomePhone    string `bson:"home_phone,omitempty" json:"home_phone,omitempty"`
	MobilePhone  string `bson:"mobile_phone,omitempty" json:"mobile_phone,omitempty"`
	Email        string `bson:"email,omitempty" json:"email,omitempty"`
}

type afePositionDesired struct {
	PositionDesired string  `bson:"position_desired,omitempty" json:"position_desired,omitempty"`
	SalaryDesired   float64 `bson:"salary_desired,omitempty" json:"salary_desired,omitempty"`
	HasCruiseShip   bool    `bson:"has_cruise_ship,omitempty" json:"has_cruise_ship,omitempty"`
	ListLastCompany string  `bson:"list_last_company,omitempty" json:"list_last_company,omitempty"`
}

type afeDocumentationInformation struct {
	PassportNumber      string                 `bson:"passport_number,omitempty" json:"passport_number,omitempty"`
	PassportNationality string                 `bson:"passport_nationality,omitempty" json:"passport_nationality,omitempty"`
	DateOfIssue         time.Time              `bson:"date_of_issue,omitempty" json:"date_of_issue,omitempty"`
	PlaceOfIssue        string                 `bson:"place_of_issue,omitempty" json:"place_of_issue,omitempty"`
	DateOfExpiration    time.Time              `bson:"date_of_expiration,omitempty" json:"date_of_expiration,omitempty"`
	CrewVisas           []afeCrewVisas         `bson:"crew_visas,omitempty" json:"crew_visas,omitempty"`
	STCWCertification   []afeSTCWCertification `bson:"stcw_certification,omitempty" json:"stcw_certification,omitempty"`
	SeamansBook         []afeSeamansBook       `bson:"seamans_book,omitempty" json:"seamans_book,omitempty"`
	OtherCertificates   []afeOtherCertificates `bson:"other_certificates,omitempty" json:"other_certificates,omitempty"`
}

type afeCrewVisas struct {
	Name             string    `bson:"name,omitempty" json:"name,omitempty"`
	YesOrNo          bool      `bson:"yes_or_no,omitempty" json:"yes_or_no,omitempty"`
	DateOfExpiration time.Time `bson:"date_of_expiration,omitempty" json:"date_of_expiration,omitempty"`
	VisaNo           string    `bson:"visa_no,omitempty" json:"visa_no,omitempty"`
	Type             string    `bson:"type,omitempty" json:"type,omitempty"`
}

type afeSTCWCertification struct {
	Type              string    `bson:"type,omitempty" json:"type,omitempty"`
	YesOrNo           bool      `bson:"yes_or_no,omitempty" json:"yes_or_no,omitempty"`
	DateOfExpiration  time.Time `bson:"date_of_expiration,omitempty" json:"date_of_expiration,omitempty"`
	CertificateNumber string    `bson:"certificate_number,omitempty" json:"certificate_number,omitempty"`
}

type afeSeamansBook struct {
	Type             string    `bson:"type,omitempty" json:"type,omitempty"`
	YesOrNo          bool      `bson:"yes_or_no,omitempty" json:"yes_or_no,omitempty"`
	DateOfExpiration time.Time `bson:"date_of_expiration,omitempty" json:"date_of_expiration,omitempty"`
	Number           string    `bson:"number,omitempty" json:"number,omitempty"`
	Nationality      string    `bson:"nationality,omitempty" json:"nationality,omitempty"`
}

type afeOtherCertificates struct {
	Type             string    `bson:"type,omitempty" json:"type,omitempty"`
	YesOrNoOrNa      string    `bson:"yes_or_no,omitempty" json:"yes_or_no_or_na,omitempty"`
	DateOfIssue      time.Time `bson:"date_of_issue,omitempty" json:"date_of_issue,omitempty"`
	DateOfExpiration time.Time `bson:"date_of_expiration,omitempty" json:"date_of_expiration,omitempty"`
	Comments         string    `bson:"comments,omitempty" json:"comments,omitempty"`
}

type afeEmploymentHistory struct {
	EmployerName     string    `bson:"employer_name,omitempty" json:"employer_name,omitempty"`
	CompanyPhoneNo   string    `bson:"company_phone_no,omitempty" json:"company_phone_no,omitempty"`
	PositionHeld     string    `bson:"position_held,omitempty" json:"position_held,omitempty"`
	SupervisorName   string    `bson:"supervisor_name,omitempty" json:"supervisor_name,omitempty"`
	From             time.Time `bson:"from,omitempty" json:"from,omitempty"`
	To               time.Time `bson:"to,omitempty" json:"to,omitempty"`
	StartingSalary   float64   `bson:"starting_salary,omitempty" json:"starting_salary,omitempty"`
	EndSalary        float64   `bson:"end_salary,omitempty" json:"end_salary,omitempty"`
	ReasonForLeaving string    `bson:"reason_for_leaving,omitempty" json:"reason_for_leaving,omitempty"`
}

type afeEducation struct {
	Type              string    `bson:"type,omitempty" json:"type,omitempty"`
	SchoolNameAndCity string    `bson:"school_name_and_city,omitempty" json:"school_name_and_city,omitempty"`
	NoOfYears         int       `bson:"no_of_years,omitempty" json:"no_of_years,omitempty"`
	From              time.Time `bson:"from,omitempty" json:"from,omitempty"`
	To                time.Time `bson:"to,omitempty" json:"to,omitempty"`
	Major             string    `bson:"major,omitempty" json:"major,omitempty"`
}

type afeLanguages struct {
	Language              string `bson:"language,omitempty" json:"language,omitempty"`
	ProviciencyLevelSpeak string `bson:"proviciency_level_speak,omitempty" json:"proviciency_level_speak,omitempty"`
	ProviciencyLevelWrite string `bson:"proviciency_level_write,omitempty" json:"proviciency_level_write,omitempty"`
}

func NewAFE() *AFE {
	return &AFE{}
}

func (model AFE) GetCollection() string {
	return "afe"
}
