package outbound

// Repository defines a generic interface for repository operations.
// It is parameterized by K for the key type and T for the entity type.
type Repository[K any, T any] interface {
	// Save adds a new T struct to the repository.
	// It returns the saved T struct.
	// If there is an error during the process, it returns an error.
	Save(T) (T, error)
	// FindById retrieves a T struct by its ID from the in-memory repository.
	// It returns a pointer to the T struct if found, or nil if not found.
	FindById(id K) *T
	FindAll() []T
	Update(K, T) *T
	UpdatePartial(K int, updates map[string]interface{}) *T
	Delete(T) error
}
