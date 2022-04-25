package config

type StorageHandler struct {
	RaftConfiguration *RaftConfig
}

type PersistenceStore interface {
	Store(*Reminder) error
	RetrieveLatest() (*Reminder, error)
	RetrieveAll() ([]*Reminder, error)
}

// New is a factory method that is responsible for instantiating
// a new instance of the persistence storage mechanism, which is
// crucial for reading and writing data from disk
func New(reminderStoreLocation string) (*StorageHandler, error) {
	return &StorageHandler{}, nil
}

// Store is responsible for taking a reminder
// and serializing it to disk or memory
func (sh *StorageHandler) Store(r *Reminder) error {
	return nil
}

// RetrieveLatest is responsible for retrieving the most
// recent reminder from memory or from disk
func (sh *StorageHandler) RetrieveLatest() (*Reminder, error) {
	return nil, nil
}

// RetrieveAll is reponsible for retrieving a list of all the
// reminders that have been persisted in either disk or memory
func (sh *StorageHandler) RetrieveAll() ([]*Reminder, error) {
	return nil, nil
}
