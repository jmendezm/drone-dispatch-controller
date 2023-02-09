package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmendezm/drone-dispatch-controller/internal/entity"
	log "github.com/sirupsen/logrus"
	"strings"
)

var ERR_INTERNAL = errors.New("internal error")

type DroneSqliteRepository struct {
	Conn *sql.DB
}

func (repo *DroneSqliteRepository) GetDroneBySerialNumber(ctx context.Context, serialNumber string) (*entity.Drone, error) {
	var (
		model         string
		state         string
		serial_number string
		battery       float32
		weight        float32
	)
	q := `select serial_number, model, weight_limit, battery_capacity, state from drone where serial_number = ?`
	r := repo.Conn.QueryRowContext(ctx, q, serialNumber)
	if err := r.Scan(&serial_number, &model, &weight, &battery, &state); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Error("sqlite_repo | GetDroneBySerialNumber() | scanning results | ", err)
		return nil, ERR_INTERNAL
	}
	m, _ := entity.DroneModelFromString(model)
	s, _ := entity.DroneStateFromString(state)
	drone := &entity.Drone{
		SerialNumber:    serial_number,
		Model:           m,
		WeightLimit:     weight,
		BatteryCapacity: battery,
		State:           s,
		Load:            nil,
	}
	return drone, nil
}

func (repo *DroneSqliteRepository) RegisterDrone(ctx context.Context, drone *entity.Drone) error {
	q := `insert into drone(serial_number, model, weight_limit, battery_capacity, state) values(?, ?, ?, ?, ?)`
	model := entity.DroneModelToString(drone.Model)
	state := entity.DroneStateToString(drone.State)
	_, err := repo.Conn.ExecContext(ctx, q, drone.SerialNumber, model, drone.WeightLimit, drone.BatteryCapacity, state)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint") {
			return errors.New("drone already exists")
		}
		log.Error("sqlite_repo | RegisterDrone() | registering drone | ", err)
		return ERR_INTERNAL
	}
	return nil
}

func (repo *DroneSqliteRepository) LoadDrone(ctx context.Context, drone *entity.Drone) error {
	q := `insert into drone_load(drone_serial_number, name, weight, code, image) values(?, ?, ?, ?, ?)`

	tx, err := repo.Conn.BeginTx(ctx, nil)
	if err != nil {
		log.Error("sqlite_repo | LoadDrone() | starting transaction | ", err)
		return ERR_INTERNAL
	}

	defer func() {
		tx.Rollback()
	}()

	_, err = tx.ExecContext(ctx, `update drone set state = ? where serial_number = ?`, drone.State, drone.SerialNumber)
	if err != nil {
		log.Error("sqlite_repo | LoadDrone() | updating drone state | ", err)
		return ERR_INTERNAL
	}

	for _, m := range drone.Load {
		_, err := tx.ExecContext(ctx, q, drone.SerialNumber, m.Name, m.Weight, m.Code, m.Image)
		if err != nil {
			log.Error("sqlite_repo | LoadDrone() | inserting medication item | ", err)
			return ERR_INTERNAL
		}
	}

	if err = tx.Commit(); err != nil {
		log.Error("sqlite_repo | LoadDrone() | committing transaction | ", err)
		return ERR_INTERNAL
	}
	return nil
}

func (repo *DroneSqliteRepository) GetAvailableDrones(ctx context.Context) ([]*entity.Drone, error) {
	q := `select serial_number, model, weight_limit, battery_capacity, state from drone where state = 'IDLE' and battery_capacity > 25`
	rows, err := repo.Conn.QueryContext(ctx, q)
	if err != nil {
		if err == sql.ErrNoRows {
			return make([]*entity.Drone, 0), nil
		}
		log.Error("sqlite_repo | GetAvailableDrones() | executing query | ", err)
		return nil, ERR_INTERNAL
	}
	var (
		model         string
		state         string
		serial_number string
		battery       float32
		weight        float32
	)
	drones := make([]*entity.Drone, 0)
	for rows.Next() {
		if err := rows.Scan(&serial_number, &model, &weight, &battery, &state); err != nil {
			log.Error("sqlite_repo | GetAvailableDrones() | scanning rows | ", err)
			return nil, ERR_INTERNAL
		}
		m, _ := entity.DroneModelFromString(model)
		s, _ := entity.DroneStateFromString(state)
		drones = append(drones, &entity.Drone{
			SerialNumber:    serial_number,
			Model:           m,
			WeightLimit:     weight,
			BatteryCapacity: battery,
			State:           s,
			Load:            nil,
		})
	}
	return drones, nil
}

func (repo *DroneSqliteRepository) GetDroneLoad(ctx context.Context, droneSerial string) ([]*entity.Medication, error) {
	q := `select name, weight, code, image from drone_load where drone_serial_number = ?`
	rows, err := repo.Conn.QueryContext(ctx, q, droneSerial)
	if err != nil {
		if err == sql.ErrNoRows {
			return make([]*entity.Medication, 0), nil
		}
		log.Error("sqlite_repo | GetDroneLoad() | executing query | ", err)
		return nil, ERR_INTERNAL
	}
	var (
		name   string
		weight float32
		code   string
		image  string
	)
	items := make([]*entity.Medication, 0)
	for rows.Next() {
		if err := rows.Scan(&name, &weight, &code, &image); err != nil {
			log.Error("sqlite_repo | GetDroneLoad() | scanning rows | ", err)
			return nil, ERR_INTERNAL
		}
		items = append(items, &entity.Medication{
			Name:   name,
			Weight: weight,
			Code:   code,
			Image:  image,
		})
	}
	return items, nil
}
