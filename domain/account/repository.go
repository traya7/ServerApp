package account

type Repository interface {
	GetAccountByUsername(string) (*Account, error)
	GetAccountById(string) (*Account, error)
}
