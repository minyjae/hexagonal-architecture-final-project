package entities

import "time"

type Faculty string

const (
	FacultyOfEducation                             Faculty = "คณะศึกษาศาสตร์"
	FacultyOfHumanities                            Faculty = "คณะมนุษยศาสตร์"
	FacultyOfFineArts                              Faculty = "คณะวิจิตรศิลป์"
	FacultyOfSocialSciences                        Faculty = "คณะสังคมศาสตร์"
	FacultyOfScience                               Faculty = "คณะวิทยาศาสตร์"
	FacultyOfEngineering                           Faculty = "คณะวิศวกรรมศาสตร์"
	FacultyOfMedicine                              Faculty = "คณะแพทยศาสตร์"
	FacultyOfAgriculture                           Faculty = "คณะเกษตรศาสตร์"
	FacultyOfDentistry                             Faculty = "คณะทันตแพทยศาสตร์"
	FacultyOfPharmacy                              Faculty = "คณะเภสัชศาสตร์"
	FacultyOfMedTechnology                         Faculty = "คณะเทคนิคการแพทย์"
	FacultyOfNursing                               Faculty = "คณะพยาบาลศาสตร์"
	FacultyOfAgroIndustry                          Faculty = "คณะอุตสาหกรรมเกษตร"
	FacultyOfVeterinaryMedicine                    Faculty = "คณะสัตวแพทยศาสตร์"
	FacultyOfArchitecture                          Faculty = "คณะสถาปัตยกรรมศาสตร์"
	BachelorOfBussinessAdministration              Faculty = "คณะบริหารธุรกิจ"
	FacultyOfEconomics                             Faculty = "คณะเศรษฐศาสตร์"
	FacultyOfLaw                                   Faculty = "คณะนิติศาสตร์"
	FacultyOfMassCommunication                     Faculty = "คณะสื่อสารมวลชน"
	FacultyOfPoliticalScience                      Faculty = "คณะรัฐศาสตร์"
	CollegeOfArtMediaAndTechnology                 Faculty = "วิทยาลัยศิลปะ สื่อ และเทคโนโลยี"
	FacultyOfPublicHealth                          Faculty = "คณะสาธารณสุขศาสตร์"
	CollegeOfEducationAndSeaManagement             Faculty = "วิทยาลัยการศึกษาและการจัดการทะเล"
	InternationalCollegeOfDigitalTechnology        Faculty = "วิทยาลัยนานาชาติเทคโนโลยีดิจิทัล"
	InstituteOfPublic                              Faculty = "สถาบันนโยบายสาธารณะ"
	BiomedicalEngineeringInstitute                 Faculty = "สถาบันวิศวกรรมชีวการแพทย์"
	AcademicServiceCenter                          Faculty = "สำนักบริการวิชาการ"
	OfficeOfEducationQualityDevelopment            Faculty = "สำนักงานพัฒนาคุณภาพการศึกษา"
	OfficeOfTheUniversity                          Faculty = "สำนักงานมหาวิทยาลัย"
	LannaRiceResearchCenter                        Faculty = "ศูนย์วิจัยข้าวล้านนา"
	LanguageInstitute                              Faculty = "สถาบันภาษา"
	SchoolOfLifeLongEducation                      Faculty = "วิทยาลัยการศึกษาตลอดชีวิต"
	CenterOfSafetyOccupationalHealthAndEnvironment Faculty = "ศูนย์บริหารจัดการความปลอดภัยฯ SHE"
	TeachingAndLearningInnovationCenter            Faculty = "ศูนย์นวัตกรรมการสอนและการเรียนรู้ TLIC"
)

type ListQueue struct {
	ID                   uint           `json:"id"`
	Priority             int            `json:"priority"`
	Title                string         `json:"title"`
	StaffID              uint           `json:"staff_id"`
	Faculty              Faculty        `json:"faculty"`
	StaffStatusID        uint           `json:"staff_status_id"`
	StaffStatus          StaffStatus    `json:"staff_status"`
	UsersStatusID        uint           `json:"users_status_id"`
	UsersStatus          UserStatus     `json:"users_status"`
	DateWordFileSubmit   time.Time      `json:"wordfile_submit"`
	DateInfoSubmit       time.Time      `json:"info_submit"`
	DateInfoSubmit14Days time.Time      `json:"info_submit_14days"`
	DateRegister         time.Time      `json:"time_register"`
	DateLeft             *int           `json:"date_left"` // ใช้ pointer เพื่อรองรับ null
	OnWeb                time.Time      `json:"on_web"`
	AppointmentDateAW    time.Time      `json:"appointment_date_aw"`
	OrderMappings        []OrderMapping `json:"order_mappings"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
}
