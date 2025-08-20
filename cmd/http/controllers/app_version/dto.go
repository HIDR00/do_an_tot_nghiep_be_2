package app_version

import "time"

type ListAppVersionResponse struct {
	ID                 int        `json:"id"`
	Os                 string     `json:"os"`
	CurrentVersionCode int        `json:"current_version_code"`
	CurrentVersionName string     `json:"current_version_name"`
	MinimumVersionCode int        `json:"minimum_version_code"`
	MinimumVersionName string     `json:"minimum_version_name"`
	DescriptionForceEn string     `json:"description_force_en"`
	DescriptionForceVn string     `json:"description_force_vn"`
	DescriptionAlertEn string     `json:"description_alert_en"`
	DescriptionAlertVn string     `json:"description_alert_vn"`
	IsMaintenance      bool       `json:"is_maintenance"`
	UpdatedAt          *time.Time `json:"updated_at"`
	CreatedAt          *time.Time `json:"created_at"`
}
