package productrepository

type ProductRepository interface {
	Save()
	Delete()
	Search()
	UpdatePrice()
	UpdateDescription()
	UpdateTitle()
}
