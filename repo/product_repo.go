package repo

func GetProductsByIds(ids []int) ([]Product, error) {
	var products []Product
	err := getDb().Where("id IN ?", ids).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
