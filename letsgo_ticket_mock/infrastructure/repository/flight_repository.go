package repositories

import (
	"errors"
	"fmt"
	"letsgo-flight-provider/internal/core/entities"
	"time"

	"github.com/google/uuid"
)

func (r *PGRepository) GetFlightList(source, destination, departureDateStr string) ([]entities.Flight, error) {
	var flights []entities.Flight

	dayStart, err := time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00Z", departureDateStr))
	if err != nil {
		return nil, err
	}
	dayEnd := dayStart.Add(24 * time.Hour)

	g := r.DB.Where("source = ? AND destination = ? AND departure_date >= ? AND departure_date < ?", source, destination, dayStart, dayEnd).Find(&flights)
	if g.Error != nil {
		return nil, g.Error
	}

	return flights, nil
}

func (r *PGRepository) GetFlightById(id uuid.UUID) (entities.Flight, error) {
	var flight entities.Flight

	g := r.DB.Where("id = ?", id).First(&flight)
	if g.Error != nil {
		return entities.Flight{}, g.Error
	}

	return flight, nil

}

func (r *PGRepository) UpdateFlightById(id uuid.UUID, action string, count int) (bool, error) {
	var flight entities.Flight

	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Model(entities.Flight{}).Where("id = ?", id).First(&flight).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}

	switch action {
	case "cancel":
		flight.RemainingSeat += count

	case "reserve":
		if flight.RemainingSeat-count < 0 {
			tx.Rollback()
			return false, errors.New("not enough remaining seats")
		}

		flight.RemainingSeat -= count

	default:
		tx.Rollback()
		return false, errors.New("invalid action")
	}

	flight.ModifiedAt = time.Now()

	if err := tx.Save(&flight).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()

	return true, nil
}

func (r *PGRepository) GetAircraftList() ([]string, error) {
	var aircrafts []string
	var flight entities.Flight

	if err := r.DB.Model(&flight).Distinct().Pluck("aircraft_name", &aircrafts).Error; err != nil {
		return nil, err
	}

	return aircrafts, nil
}

func (r *PGRepository) GetcitytList() ([]string, error) {
	var cities []string
	var destinationCities []string
	var flight entities.Flight

	if err := r.DB.Model(&flight).Distinct().Pluck("source", &cities).Error; err != nil {
		return nil, err
	}
	if err := r.DB.Model(&flight).Distinct().Pluck("destination", &destinationCities).Error; err != nil {
		return nil, err
	}
	cities = append(cities, destinationCities...)
	return cities, nil
}

func (r *PGRepository) GetListDaysWithFlight() ([]string, error) {
	var timeDates []string

	today := time.Now().Format("2006-01-02")
	g := r.DB.Model(&entities.Flight{}).Select("DISTINCT DATE(departure_date)").Where("departure_date >= ?", today).Pluck("DISTINCT DATE(departure_date)", &timeDates)
	if g.Error != nil {
		return nil, g.Error
	}
	return timeDates, nil
}
