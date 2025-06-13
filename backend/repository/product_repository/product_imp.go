package productrepository

type ProductRepositoryImpl struct{}

// Delete implements ProductRepository.
func (p ProductRepositoryImpl) Delete() {
	panic("unimplemented")
}

// Save implements ProductRepository.
func (p ProductRepositoryImpl) Save() {
	panic("unimplemented")
}

// Search implements ProductRepository.
func (p ProductRepositoryImpl) Search() {
	panic("unimplemented")
}

// UpdateDescription implements ProductRepository.
func (p ProductRepositoryImpl) UpdateDescription() {
	panic("unimplemented")
}

// UpdatePrice implements ProductRepository.
func (p ProductRepositoryImpl) UpdatePrice() {
	panic("unimplemented")
}

// UpdateTitle implements ProductRepository.
func (p ProductRepositoryImpl) UpdateTitle() {
	panic("unimplemented")
}

func NewProductRepository() ProductRepository {
	return ProductRepositoryImpl{}
}
