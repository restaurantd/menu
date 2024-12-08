package data

import "menu/internal/data/models"

func (d DB) GetProduct(title string) (models.Product, error) {
	var product models.Product

	if err := d.Where(models.Product{Title: title}).First(&product).Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (d DB) GetElement(listname string) ([]models.Element, error) {
	var elements []models.Element

	if err := d.Where(models.Element{Listname: listname}).Find(&elements).Error; err != nil {
		return elements, err
	}

	return elements, nil
}
