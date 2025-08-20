package entities

import "time"

type AppVersionData struct {
	ID                 int        `json:"id" db:"id"`
	Os                 string     `json:"os" db:"os"`
	CurrentVersionCode int        `json:"current_version_code" db:"current_version_code"`
	CurrentVersionName string     `json:"current_version_name" db:"current_version_name"`
	MinimumVersionCode int        `json:"minimum_version_code" db:"minimum_version_code"`
	MinimumVersionName string     `json:"minimum_version_name" db:"minimum_version_name"`
	DescriptionForceEn string     `json:"description_force_en" db:"description_force_en"`
	DescriptionForceVn string     `json:"description_force_vn" db:"description_force_vn"`
	DescriptionAlertEn string     `json:"description_alert_en" db:"description_alert_en"`
	DescriptionAlertVn string     `json:"description_alert_vn" db:"description_alert_vn"`
	IsMaintenance      bool       `json:"is_maintenance" db:"is_maintenance"`
	UpdatedAt          *time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt          *time.Time `json:"created_at" db:"created_at"`
}
