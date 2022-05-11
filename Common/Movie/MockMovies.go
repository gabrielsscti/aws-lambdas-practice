package Movie

import (
	"encoding/json"
	"errors"
)

type MockMovies struct{}

var movies = []Movie{
	{
		ID:   1,
		Name: "Avengers",
	},
	{
		ID:   2,
		Name: "Ant-Man",
	},
	{
		ID:   3,
		Name: "Thor",
	},
	{
		ID:   4,
		Name: "Hulk",
	},
	{
		ID:   5,
		Name: "Doctor Strange",
	},
}

func CreateNewMockMovies() MockMovies {
	return MockMovies{}
}

func (m MockMovies) GetByID(id int) ([]byte, error) {
	id--
	if (id < 0) || (id > len(movies)-1) {
		return nil, errors.New("Not found")
	}

	return json.Marshal(movies[id])
}

func (m MockMovies) GetAll() ([]byte, error) {
	return json.Marshal(movies)
}

func (m MockMovies) InsertMovie(movie []byte) error {
	var movieObj Movie
	err := json.Unmarshal(movie, &movieObj)
	if err != nil {
		return err
	}
	movies = append(movies, movieObj)
	return nil
}
