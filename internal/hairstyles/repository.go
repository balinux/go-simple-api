package hairstyles

import (
	"gorm.io/gorm"
)

type HairStyleRepository struct {
	db *gorm.DB
}

func NewHairStyleRepository(db *gorm.DB) *HairStyleRepository {
	return &HairStyleRepository{
		db: db,
	}
}

func (r *HairStyleRepository) GetAll() ([]HairStyle, error) {
	var hairStyles []HairStyle
	if err := r.db.Find(&hairStyles).Error; err != nil {
		return nil, err
	}
	return hairStyles, nil
}

func (r *HairStyleRepository) Create(hairStyle *HairStyle) error {
	return r.db.Create(hairStyle).Error
}

func (r *HairStyleRepository) FindById(id uint) (*HairStyle, error) {
	var hairstyle HairStyle
	if err := r.db.First(&hairstyle, id).Error; err != nil {
		return nil, err
	}
	return &hairstyle, nil
}

func (r *HairStyleRepository) DeleteByID(id uint) error {
	if err := r.db.Delete(&HairStyle{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *HairStyleRepository) Update(HairStyle *HairStyle) error {
	if err := r.db.Save(HairStyle).Error; err != nil {
		return err
	}
	return nil
}
