package settingsRepository

import "sls/internal/entity/adminEntity"

type SettingsRepo interface {
	CreateSettings(settings *adminEntity.AdminSettings) (*adminEntity.AdminSettings, error)
	EditSettings(id string, settings *adminEntity.AdminSettings) (*adminEntity.AdminSettings, error)
	FetchSettings() (*adminEntity.AdminSettings, error)
}
