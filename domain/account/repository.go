package account

type Repository interface {
	Store(acc Account) error
	FindOneByUsername(string) (*Account, error)
	FindOneByID(string) (*Account, error)
}
