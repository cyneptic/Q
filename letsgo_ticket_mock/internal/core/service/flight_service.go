package service

import (
	repositories "letsgo-flight-provider/infrastructure/repository"
	"letsgo-flight-provider/internal/core/entities"
	ports "letsgo-flight-provider/internal/core/port"
	"time"

	"github.com/google/uuid"
)

type FlightService struct {
	db ports.FlightRepositoryContract
}

func NewFlightService() *FlightService {
	db := repositories.NewGormDatabase()

	return &FlightService{
		db: db,
	}
}

func (svc *FlightService) GetFlightList(source, destination, departure string) ([]entities.Flight, error) {
	return svc.db.GetFlightList(source, destination, departure)
}

func (svc *FlightService) GetFlightById(id uuid.UUID) (entities.Flight, error) {
	return svc.db.GetFlightById(id)
}

func (svc *FlightService) UpdateFlightById(id uuid.UUID, action string, count int) (bool, error) {
	return svc.db.UpdateFlightById(id, action, count)
}

func (svc *FlightService) GetAircraftList() ([]string, error) {
	return svc.db.GetAircraftList()
}

func (svc *FlightService) GetcitytList() ([]string, error) {
	cities, err := svc.db.GetcitytList()
	if err != nil {
		return nil, err
	}
	uniqueCities := make([]string, 0, len(cities))
	encountered := map[string]bool{}

	for _, city := range cities {
		if !encountered[city] {
			encountered[city] = true
			uniqueCities = append(uniqueCities, city)
		}
	}
	return uniqueCities, nil
}

func (svc *FlightService) GetListDaysWithFlight() ([]string, error) {
	timeDates, err := svc.db.GetListDaysWithFlight()
	if err != nil {
		return nil, err
	}
	var dates []string
	for _, d := range timeDates {
		tm, err := time.Parse(time.RFC3339, d)
		if err != nil {
			return nil, err
		}
		date := tm.Format("2006-01-02")
		dates = append(dates, date)
	}

	return dates, nil
}
