package main

import (
	_ "embed"
	"log"
	"net/http"
	"os"

	"github.com/nathancastelein/go-introduction-api/examples/database/bob"
	"github.com/nathancastelein/go-introduction-api/exercise/api/nethttp"
	"github.com/nathancastelein/go-introduction-api/shelter"
)

//go:embed openapi.yaml
var openapiSpec []byte

const swaggerUI = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Swagger UI</title>
    <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css">
</head>
<body>
    <div id="swagger-ui"></div>
    <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
    <script>
        SwaggerUIBundle({
            url: "/openapi.yaml",
            dom_id: '#swagger-ui',
        });
    </script>
</body>
</html>`

type PetRepositoryImpl struct{}

// SelectAllPets implements [shelter.PetRepository].
func (p *PetRepositoryImpl) SelectAllPets() ([]shelter.Pet, error) {
	return bob.SelectAllPets()
}

// SelectOnePet implements [shelter.PetRepository].
func (p *PetRepositoryImpl) SelectOnePet(id int) (*shelter.Pet, error) {
	return bob.SelectOnePet(id)
}

// UpdatePetName implements [shelter.PetRepository].
func (p *PetRepositoryImpl) UpdatePetName(id int, newName string) error {
	return bob.UpdatePetName(id, newName)
}

func main() {
	mux := nethttp.NewServer(&PetRepositoryImpl{})

	mux.HandleFunc("GET /openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/yaml")
		w.Write(openapiSpec)
	})

	mux.HandleFunc("GET /swagger", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(swaggerUI))
	})

	if err := mux.Run(os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
}
