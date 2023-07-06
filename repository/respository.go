package repository

import (
	"errors"

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
	DeleteSpaceship(id int) (map[string]bool, error)
	CreateSpaceship(spaceship *models.Spaceship) (map[string]bool, error)
	UpdateSpaceship(id int, spaceship *models.Spaceship) (map[string]bool, error)
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

//Get all spaceships
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


//Filter spaceships by Name
func (db *MySQLDb) FilterAllByClass(class string) ([]models.Spaceship, error) {
	var spaceships []models.Spaceship

	resp := database.DB.Where("class = ?", class).Find(&spaceships)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return spaceships, nil
}

//Filter spaceships by Status
func (db *MySQLDb) FilterAllByStatus(status string) ([]models.Spaceship, error) {

	var spaceships []models.Spaceship

	resp := database.DB.Where("status = ?", status).Find(&spaceships)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return spaceships, nil
}

//Get Single Spaceship
func (db *MySQLDb) GetSingleSpaceship(Id int) (*models.Spaceship, error) {

	spaceship := &models.Spaceship{}

	if err := database.DB.Where("id = ?", Id).First(spaceship).Error; err != nil {
		return nil, err
	}

	return spaceship, nil

}

//Create Spaceship
func (db *MySQLDb) CreateSpaceship(spaceship *models.Spaceship) (map[string]bool, error) {
	isvalid := spaceship.IsValidSpaceship()
	if !isvalid {
		return nil, errors.New("spaceship not valid")
	}

	resp := database.DB.Create(&spaceship)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return map[string]bool{"success": true}, resp.Error
}

//Update Spaceship
func (db *MySQLDb) UpdateSpaceship(id int, spaceship *models.Spaceship) (map[string]bool, error) {

	isvalid := spaceship.IsValidSpaceship()
	if !isvalid {
		return nil, errors.New("spaceship not valid")
	}

	resp := db.DB.Model(&models.Spaceship{}).Where("id = ?", id).Updates(spaceship)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return map[string]bool{"success": true}, resp.Error
}


//Delete Spaceship
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

