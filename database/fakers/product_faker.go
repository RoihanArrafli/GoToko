package fakers

import (
	"log"
	"math"
	"math/rand"
	"time"

	"gorm.io/gorm"

	"github.com/RoihanArrafli/GoToko/app/models"
	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/shopspring/decimal"
)

func ProductFaker(db *gorm.DB) *models.Product {
	user := UserFaker(db)
	err := db.Create(&user).Error
	if err != nil {
		log.Fatal(err)
	}

	name := faker.Name()
	return &models.Product{
		ID:     uuid.New().String(),
		UserID: user.ID,
		// ProductImages:    nil,
		// Categories:       nil,
		Sku:              slug.Make(name),
		Name:             name,
		Slug:             slug.Make(name),
		Price:            decimal.NewFromFloat(fakerPrice()),
		Stock:            rand.Intn(100),
		Weight:           decimal.NewFromFloat(rand.Float64()),
		ShortDescription: faker.Paragraph(),
		Description:      faker.Paragraph(),
		Status:           0,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		DeletedAt:        gorm.DeletedAt{},
	}
}

func fakerPrice() float64 {
	return precision(rand.Float64()*math.Pow10(rand.Intn(8)), rand.Intn(2)+1)
}

func precision(val float64, pre int) float64 {
	div := math.Pow10(pre)
	return float64(int64(val+div)) / div
}