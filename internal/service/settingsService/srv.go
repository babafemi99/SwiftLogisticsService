package settingsService

import (
	"github.com/google/uuid"
	"log"
	"sls/internal/Repository/settingsRepository"
	"sls/internal/entity/adminEntity"
	"sls/internal/service/validationService"
)

type SettingsService interface {
	CreateSettings(access *adminEntity.AdminSettings) (*adminEntity.AdminSettings, error)
	EditSettings(id string, settings *adminEntity.AdminSettings) (*adminEntity.AdminSettings, error)
	GetSettings() (*adminEntity.AdminSettings, error)
}
type settingsService struct {
	settingsRepo settingsRepository.SettingsRepo
	vldSrv       validationService.ValidationService
}

func (s *settingsService) GetSettings() (*adminEntity.AdminSettings, error) {
	settings, err := s.settingsRepo.FetchSettings()
	if err != nil {
		return nil, err
	}
	return settings, nil
}

func NewSettingsService(settingsRepo settingsRepository.SettingsRepo, vldSrv validationService.ValidationService) SettingsService {
	return &settingsService{settingsRepo: settingsRepo, vldSrv: vldSrv}
}

func (s *settingsService) CreateSettings(access *adminEntity.AdminSettings) (*adminEntity.AdminSettings, error) {
	access.Id = uuid.New().String()
	err := s.vldSrv.Validate(access)
	if err != nil {
		return nil, err
	}

	err = s.vldSrv.ValidateDPR(&access.DeliveryPricingRate)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	settings, err := s.settingsRepo.CreateSettings(access)
	if err != nil {
		return nil, err
	}
	return settings, nil
}

func (s *settingsService) EditSettings(id string, settings *adminEntity.AdminSettings) (*adminEntity.AdminSettings, error) {
	err := s.vldSrv.Validate(settings)
	if err != nil {
		log.Println(1)
		return nil, err
	}

	err = s.vldSrv.ValidateDPR(&settings.DeliveryPricingRate)
	if err != nil {
		log.Println(2)
		return nil, err
	}

	settings, err = s.settingsRepo.EditSettings(id, settings)
	if err != nil {
		return nil, err
	}
	return settings, nil
}
