package models

import (
	"time"

	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
)

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

// ListQueue table
type ListQueue struct {
	ID                   uint           `gorm:"primaryKey" json:"id"`
	Priority             int            `gorm:"not null;index" json:"priority"`
	Title                string         `gorm:"not null" json:"title"`
	StaffID              uint           `gorm:"not null;index" json:"staff_id"`
	Faculty              Faculty        `gorm:"not null" json:"faculty"`
	StaffStatusID        uint           `gorm:"not null;index" json:"staff_status_id"`
	StaffStatus          StaffStatus    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"staff_status"`
	UsersStatusID        uint           `gorm:"not null;index" json:"users_status_id"`
	UsersStatus          UserStatus     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"users_status"`
	DateWordFileSubmit   time.Time      `json:"wordfile_submit"`
	DateInfoSubmit       time.Time      `json:"info_submit"`
	DateInfoSubmit14Days time.Time      `json:"info_submit_14days"`
	DateRegister         time.Time      `json:"time_register"`
	DateLeft             *int           `json:"date_left"` // ใช้ pointer เพื่อรองรับ null
	OnWeb                time.Time      `json:"on_web"`
	AppointmentDateAW    time.Time      `json:"appointment_date_aw"`
	OrderMappings        []OrderMapping `gorm:"foreignKey:ListQueueID" json:"order_mappings"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
}

// แปลงโมเดล ListQueue เป็น entity
func (l *ListQueue) ToEntity() *entities.ListQueue {
	return &entities.ListQueue{
		ID:                   l.ID,
		Priority:             l.Priority,
		Title:                l.Title,
		StaffID:              l.StaffID,
		Faculty:              entities.Faculty(l.Faculty),
		StaffStatusID:        l.StaffStatusID,
		UsersStatusID:        l.UsersStatusID,
		DateWordFileSubmit:   l.DateWordFileSubmit,
		DateInfoSubmit:       l.DateInfoSubmit,
		DateInfoSubmit14Days: l.DateInfoSubmit14Days,
		DateRegister:         l.DateRegister,
		DateLeft:             l.DateLeft,
		OnWeb:                l.OnWeb,
		AppointmentDateAW:    l.AppointmentDateAW,
		CreatedAt:            l.CreatedAt,
		UpdatedAt:            l.UpdatedAt,
	}
}

// แปลง Entity เป็นโมเดล ListQueue
func (l *ListQueue) FromEntity(entity *entities.ListQueue) {
	l.ID = entity.ID
	l.Priority = entity.Priority
	l.Title = entity.Title
	l.StaffID = entity.StaffID
	l.Faculty = Faculty(entity.Faculty)
	l.StaffStatusID = entity.StaffStatusID
	l.UsersStatusID = entity.UsersStatusID
	l.DateWordFileSubmit = entity.DateWordFileSubmit
	l.DateInfoSubmit = entity.DateInfoSubmit
	l.DateInfoSubmit14Days = entity.DateInfoSubmit14Days
	l.DateRegister = entity.DateRegister
	l.DateLeft = entity.DateLeft
	l.OnWeb = entity.OnWeb
	l.AppointmentDateAW = entity.AppointmentDateAW
	l.CreatedAt = entity.CreatedAt
	l.UpdatedAt = entity.UpdatedAt
}
