package wallet

type Wallet interface {
	Name() string

	Path() string

	Create(name, passphrase string) error

	Load(name string) error

	Lock(name string) error

	Unlock(name, passphrase string) error

	Close() error

	IsLocked(name string) (bool, error)

	//CheckAccountName(name string) (bool)

}
