package repository

import (
	"errors"
	"fmt"

	"github.com/durotimicodes/xanda_task_R3_D3/cmd/database"
	"github.com/durotimicodes/xanda_task_R3_D3/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SpaceshipRepository interface {
	GetAll() (map[string]any, error)
	FilterAllByName(name string) ([]models.Spaceship, error)
	FilterAllByClass(class string) ([]models.Spaceship, error)
	FilterAllByStatus(status string) ([]models.Spaceship, error)
	GetSingleSpaceship(id uint) (*models.Spaceship, error)
	DeleteSpaceship(id uint) (map[string]bool, error)
	CreateSpaceship(spaceship *models.Spaceship) (map[string]bool, error)
	UpdateSpaceship(id uint, spaceship *models.Spaceship) (map[string]bool, error)
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

func (mydb *MySQLDb) Init(host, user, password, dbName, port string) error {
	fmt.Println("connecting to Database.....")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode-disable TimeZone=Afirca/Lagos", host, user, password, dbName, port)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if db == nil {
		return fmt.Errorf("database was not initialized")
	} else {
		fmt.Println("Connected to database")
	}
	return nil

}

// Get all spaceships
func (db *MySQLDb) GetAll() (map[string]interface{}, error) {

	spaceships := []models.Spaceship{}

	database.DB.Find(&spaceships)

	payload := map[string]any{"data": spaceships}

	return payload, nil
}

// Filter spaceships by Name
func (db *MySQLDb) FilterAllByName(name string) ([]models.Spaceship, error) {
	var spaceships []models.Spaceship

	resp := database.DB.Where("name = ?", name).Find(&spaceships)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return spaceships, nil
}

// Filter spaceships by Class
func (db *MySQLDb) FilterAllByClass(class string) ([]models.Spaceship, error) {
	var spaceships []models.Spaceship

	resp := database.DB.Where("class = ?", class).Find(&spaceships)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return spaceships, nil
}

// Filter spaceships by Status
func (db *MySQLDb) FilterAllByStatus(status string) ([]models.Spaceship, error) {

	var spaceships []models.Spaceship

	resp := database.DB.Where("status = ?", status).Find(&spaceships)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return spaceships, nil
}

// Get Single Spaceship
func (db *MySQLDb) GetSingleSpaceship(Id uint) (*models.Spaceship, error) {

	spaceship := &models.Spaceship{}

	// checker := database.DB.Where("id = ?", Id).Limit(1).First(&spaceship)
	// if checker.RowsAffected == 0 {
	// 	return nil, errors.New("no row found with this id")
	// } else {
	database.DB.Where("id = ?", Id).First(&spaceship)

	return spaceship, nil

}

// Create Spaceship
func (db *MySQLDb) CreateSpaceship(spaceship *models.Spaceship) (map[string]bool, error) {
	isvalid := spaceship.IsValidSpaceship()
	if !isvalid {
		return nil, errors.New("spaceship not valid")
	}

	resp := database.DB.Create(&spaceship)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return map[string]bool{"success": true}, nil
}

// Update Spaceship
func (db *MySQLDb) UpdateSpaceship(id uint, spaceship *models.Spaceship) (map[string]bool, error) {

	isvalid := spaceship.IsValidSpaceship()
	if !isvalid {
		return nil, errors.New("spaceship not valid")
	}

	// checker := database.DB.Where("id = ?", id).Limit(1).Find(&spaceship)
	// if checker.RowsAffected == 0 {
	//return map[string]bool{"no row found on this id": false}, errors.New("no row found on this id")
	//} else {
	resp := db.DB.Model(&models.Spaceship{}).Where("id = ?", id).Updates(spaceship)

	if resp.Error != nil {
		return nil, resp.Error
	}
	//}

	return map[string]bool{"success": true}, nil
}

// Delete Spaceship
func (db *MySQLDb) DeleteSpaceship(Id uint) (map[string]bool, error) {
	spaceship := models.Spaceship{}

	// checker := database.DB.Where("id = ?", Id).Limit(1).Find(&spaceship)
	// if checker.RowsAffected == 0 {
	// 	return map[string]bool{"no row found on this id": false}, errors.New("no row found on this Id")
	// } else {
	err := database.DB.Where("spaceship_id = ?", Id).Delete(&spaceship.Armaments).Error
	if err != nil {
		return nil, err
	}

	err = database.DB.Where("id = ?", Id).Delete(&spaceship).Error
	if err != nil {
		return nil, err
	}

	// }

	return map[string]bool{"success": true}, nil
}
