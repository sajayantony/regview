package registry

import (
	"context"

	"github.com/sajayantony/rv/models"
	"github.com/sajayantony/rv/store"
)

// Fetch repositories from the registry
func GetRepositories() <-chan []models.Repository {
	ch := make(chan []models.Repository)
	r1 := []models.Repository{
		models.Repository{
			ID:   0,
			Name: "repo1",
		},
	}
	repos := [][]models.Repository{r1}
	go func() {
		for _, val := range repos {
			ch <- val
		}
		close(ch) // Remember to close or the loop never ends!
	}()
	return ch
}

//WriteRepositories fetches all repositories and writes to datastore
func WriteRepositories(ctx context.Context) {
	for repos := range GetRepositories() {
		for _, r := range repos {
			store.CreateRepository(ctx, &r)
		}
	}
}
