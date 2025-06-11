package outbound

// Repository defines a generic interface for repository operations.
// It is parameterized by K for the key type and T for the entity type.
type Repository[K any, T any] interface {

	// Save adds a new T struct to the repository.
	// It returns the saved T struct.
	// If there is an error during the process, it returns an error.
	Save(T) (T, error)

	// FindById retrieves a T struct by its ID from the repository.
	// It returns a pointer to the T struct if found, or nil if not found.
	FindById(id K) *T

	// FindAll retrieves all T structs from the repository.
	// It returns a slice of T structs.
	FindAll() []T

	// Update modifies an existing T struct in the repository, it expected to update the full struct.
	// It returns the updated T struct if successful, or nil if the update fails.
	Update(K, T) *T

	// UpdatePartial modifies specific fields of an existing T struct in the repository.
	// It takes a map of field names to values to update.
	// It returns the updated T struct if successful, or nil if the update fails.
	UpdatePartial(K int, updates map[string]interface{}) *T

	// Delete removes a T struct from the repository by its ID.
	// It returns an error if the deletion fails.
	Delete(T) error
}
