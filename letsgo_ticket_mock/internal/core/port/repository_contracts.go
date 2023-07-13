package ports

import (
	"letsgo-flight-provider/internal/core/entities"

	"github.com/google/uuid"
)

type FlightRepositoryContract interface {
	GetFlightList(source, destination, departure string) ([]entities.Flight, error)
	GetFlightById(id uuid.UUID) (entities.Flight, error)
	UpdateFlightById(id uuid.UUID ,action string, count int) (bool, error)
	GetAircraftList() ([]string, error)
	GetcitytList() ([]string, error)
	GetListDaysWithFlight() ([]string, error)
}
