package service

import (
	"log"

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
	CreateSpaceship(name, class, status string, crew int, value float32) map[string]bool
	CreateArmament(title, qty string)
	DeleteSpaceship(id int) (map[string]bool, error)
	UpdateSpaceship(id int, ship *models.Spaceship) (map[string]bool, error)
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
	var spaceships, filteredShips []models.Spaceship

	database.DB.Find(&spaceships)

	for _, spaceship := range spaceships {
		if spaceship.Name == name {
			filteredShips = append(filteredShips, spaceship)

		}
	}

	err := database.DB.Preload("Name").Find(&filteredShips).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return filteredShips, nil
}

func (db *MySQLDb) FilterAllByClass(class string) ([]models.Spaceship, error) {
	var spaceships []models.Spaceship

	database.DB.Find(&spaceships)

	for _, spaceship := range spaceships {
		if spaceship.Class == class {
			spaceships = append(spaceships, spaceship)
		}
	}

	err := database.DB.Preload("Class").Find(&spaceships).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return spaceships, nil
}

func (db *MySQLDb) FilterAllByStatus(status string) ([]models.Spaceship, error) {

	var spaceships, filteredShips []models.Spaceship

	database.DB.Find(&spaceships)

	for _, spaceship := range spaceships {
		if spaceship.Status == status {
			filteredShips = append(filteredShips, spaceship)

		}
	}

	err := database.DB.Preload("Status").Find(&filteredShips).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return filteredShips, nil
}

func (db *MySQLDb) GetSingleSpaceship(Id int) (*models.Spaceship, error) {

	spaceship := &models.Spaceship{}

	if err := database.DB.Where("id = ?", Id).First(spaceship).Error; err != nil {
		return nil, err
	}

	return spaceship, nil

}

func (db *MySQLDb) CreateArmament(title, qty string) {
	
	armament := models.Armament{
		Title: title,
		Qty:   qty,
	}
	

	
}

func (db *MySQLDb) CreateSpaceship(name, class, status, title, qty string, crew int, value float32) map[string]bool {

	armaments := []models.Armament{}
	armament := models.Armament{
		Title: title,
		Qty: qty,
	}

	armaments = append(armaments, armament)
	
	
	spaceship := models.Spaceship{
		Name: name,
		Class: class,
		Status: status,
		Crew: crew,
		Value: value,
		Armaments: armaments,
	}
	

	spaceship.IsValidSpaceship()
	database.DB.Create(spaceship)
	database.DB.Create(armaments)

	return map[string]bool{"success": true}

}

func (db *MySQLDb) UpdateSpaceship(id int, ship *models.Spaceship) (map[string]bool, error) {
	spaceship := models.Spaceship{}

	err := db.DB.Model(&spaceship).Where("id = ?", id).
		Update("name", ship.Name).
		Update("class", ship.Class).
		Update("armaments", ship.Armaments).
		Update("crew", ship.Crew).
		Update("Image", ship.Image).
		Update("value", ship.Value).
		Update("status", ship.Status).Error

	if err != nil {
		log.Println("error in updating mysql db")
		return nil, err
	}

	return map[string]bool{"success": true}, nil

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

func UpdateSpaceship(id int, ship *models.Spaceship) (map[string]bool, error) {
	r := &MySQLDb{}
	return r.UpdateSpaceship(id, ship)
}

func CreateNewSpaceship(name, class, status string, crew int, value float32) map[string]bool {
	r := &MySQLDb{}
	return r.CreateSpaceship(name, class, status, crew, value)
}

func DeleteSpaceshipByID(id int) (map[string]bool, error) {
	r := &MySQLDb{}
	return r.DeleteSpaceship(id)
}
