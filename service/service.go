package service

import (
	"errors"
	"strconv"

	"github.com/durotimicodes/xanda_task_R3_D3/cmd/database"
	"github.com/durotimicodes/xanda_task_R3_D3/models"
	"gorm.io/gorm"
)

type SpaceshipRepository interface {
	GetAll() ([]models.Spaceship, error)
	FilterAllByName(name string) ([]models.Spaceship, error)
	FilterAllByClass(class string) ([]models.Spaceship, error)
	FilterAllByStatus(status string) ([]models.Spaceship, error)
	GetSingleSpaceship(id int) (*models.Spaceship, error)
	CreateSpaceship(name, class, status string, crew int, value float32, armaments []models.Armament) (map[string]bool, error)
	DeleteSpaceship(id int) (map[string]bool, error)
	UpdateSpaceship(id int, name, class, status string, crew int, value float32, armaments []models.Armament) (map[string]bool, error)
}

// -------- Repository Begin --------
type MySQLDb struct {
	DB *gorm.DB
}

func NewMySqlDB(db *gorm.DB) *MySQLDb {
	return &MySQLDb{
		DB: db,
	}
}

func (db *MySQLDb) GetAll() ([]models.Spaceship, error) {

	spaceships := []models.Spaceship{}
	database.DB.Find(&spaceships)

	return spaceships, nil
}

func (db *MySQLDb) FilterAllByName(name string) ([]models.Spaceship, error) {
	var spaceships []models.Spaceship

	resp := database.DB.Where("name = ?", name).Find(&spaceships)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return spaceships, nil
}

func (db *MySQLDb) FilterAllByClass(class string) ([]models.Spaceship, error) {
	var spaceships []models.Spaceship

	resp := database.DB.Where("class = ?", class).Find(&spaceships)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return spaceships, nil
}

func (db *MySQLDb) FilterAllByStatus(status string) ([]models.Spaceship, error) {

	var spaceships []models.Spaceship

	resp := database.DB.Where("status = ?", status).Find(&spaceships)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return spaceships, nil
}

func (db *MySQLDb) GetSingleSpaceship(Id int) (*models.Spaceship, error) {

	spaceship := &models.Spaceship{}

	if err := database.DB.Where("id = ?", Id).First(spaceship).Error; err != nil {
		return nil, err
	}

	return spaceship, nil

}

func (db *MySQLDb) CreateSpaceship(name, class, status string, crew int, value float32, armaments []models.Armament) (map[string]bool, error) {

	spaceship := models.Spaceship{
		Name:   name,
		Class:  class,
		Status: status,
		Crew:   crew,
		Value:  value,
	}

	isvalid := spaceship.IsValidSpaceship()
	if !isvalid {
		return nil, errors.New("spaceship not valid")
	}

	resp := database.DB.Create(&spaceship)
	if resp.Error != nil {
		return nil, resp.Error
	}

	for i := range armaments {
		armaments[i].SpaceshipID = spaceship.ID
	}

	resp = database.DB.Create(&armaments)

	return map[string]bool{"success": true}, resp.Error

}

func (db *MySQLDb) UpdateSpaceship(id int, name, class, status string, crew int, value float32, armaments []models.Armament) (map[string]bool, error) {

	spaceship := models.Spaceship{
		Name:   name,
		Class:  class,
		Status: status,
		Crew:   crew,
		Value:  value,
	}

	isvalid := spaceship.IsValidSpaceship()
	if !isvalid {
		return nil, errors.New("spaceship not valid")
	}

	resp := db.DB.Model(&spaceship).Where("id = ?", id).
		Update("name", name).
		Update("class", class).
		Update("armaments", armaments).
		Update("crew", crew).
		Update("value", value).
		Update("status", status)

	if resp.Error != nil {
		return nil, resp.Error
	}

	for i := range armaments {
		armaments[i].SpaceshipID = spaceship.ID
	}

	resp = db.DB.Model(&armaments).Where("id = ?", strconv.Itoa(id))

	return map[string]bool{"success": true}, resp.Error

}

func (db *MySQLDb) DeleteSpaceship(Id int) (map[string]bool, error) {
	spaceship := models.Spaceship{}

	err := database.DB.Where("spaceship_id = ?", Id).Delete(&spaceship.Armaments).Error
	if err != nil {
		return nil, err
	}

	err = database.DB.Where("id = ?", Id).Delete(&spaceship).Error
	if err != nil {
		return nil, err
	}

	return map[string]bool{"success": true}, nil
}

// -------- Repository End --------



// -------- Services Begin --------
func GetAllSpaceShips() ([]models.Spaceship, error) {
	r := &MySQLDb{}
	return r.GetAll()
}

func GetAllSpaceShipsByName(name string) ([]models.Spaceship, error) {
	r := &MySQLDb{}
	return r.FilterAllByName(name)
}

func GetAllSpaceShipsByClass(class string) ([]models.Spaceship, error) {
	r := &MySQLDb{}
	return r.FilterAllByClass(class)
}

func GetAllSpaceShipsByStatus(status string) ([]models.Spaceship, error) {
	r := &MySQLDb{}
	return r.FilterAllByStatus(status)
}

func GetSpaceship(id int) (*models.Spaceship, error) {
	r := &MySQLDb{}
	return r.GetSingleSpaceship(id)
}

func UpdateSpaceship(id int, name, class, status string, crew int, value float32, armaments []models.Armament) (map[string]bool, error) {
	r := &MySQLDb{}
	return r.UpdateSpaceship(id, name, class, status, crew, value, armaments)
}


func DeleteSpaceshipByID(id int) (map[string]bool, error) {
	r := &MySQLDb{}
	return r.DeleteSpaceship(id)
}

// -------- Services End --------