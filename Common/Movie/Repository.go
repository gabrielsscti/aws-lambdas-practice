package Movie

type Repository interface {
	GetByID(id int) ([]byte, error)
	GetAll() ([]byte, error)
	InsertMovie(movie []byte) error
}
