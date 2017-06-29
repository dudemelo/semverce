package main

import (
	"github.com/pressly/chi"
	"net/http"
	"semverce/resource"
)

func main() {

	router := chi.NewRouter()

	router.Route("/projects", func (router chi.Router) {

		router.Post("/", resource.PostProjects)
		router.Get("/", resource.GetProjects)

		router.Route("/:projectId", func (router chi.Router) {
			router.Get("/", resource.GetProject)
			router.Put("/", resource.PutProject)
			router.Delete("/", resource.DeleteProject)
		})
	})

	router.Route("/repositories", func (router chi.Router) {

		router.Get("/", resource.GetRepositories)
		router.Post("/", resource.PostRepositories)

		router.Route("/:repositoryId", func (router chi.Router) {

			router.Get("/", resource.GetRepository)
			router.Put("/", resource.PutRepository)
			router.Delete("/", resource.DeleteRepository)

			router.Post("/major", resource.PostMajorRepository)
			router.Post("/minor", resource.PostMinorRepository)
			router.Post("/fix", resource.PostFixRepository)
		})
	})

	http.ListenAndServe(":4200", router)
}