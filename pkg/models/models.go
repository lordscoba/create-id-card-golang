package models

import(
	"github.com/jinzhu/gorm"
	"github.com/lordscoba/create-id-card-golang/pkg/config"
)
var db *gorm.DB


type Card struct{
	gorm.Model
	Name string `gorm:""json:"name"`
	YearOfBirth int64 `json:"Year"`
	State string `json:"state"`
	Country string `json:"country"`
	Height string `json:"height"`
	Weight string `json:"weight"`
	About string `json:"About"`
}

func init(){
	// config.Connect(){}
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Card{})
}

func (b *Card) CreateCard() *Card{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllCards() []Card{
	var Cards []Card
	db.Find(&Cards)
	return Cards
}

func GetCardById(Id int64) (*Card, *gorm.DB){
	var getCard Card
	db:=db.Where("ID=?", Id).Find(&getCard)
	return &getCard, db
}

func DeleteCard(ID int64) Card{
	var card Card
	db.Where("ID=?", ID).Delete(card)
	// db.Delete(&card, ID)
	// db.Unscoped().Where("ID = ID").Find(&card)
// SELECT * FROM users WHERE age = 20;

	return card
}
