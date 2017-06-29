package resource


import (
	"net/http"
	api "semverce/http"
)


type Project struct {
	Id		string		`json:"id"`
	Name		string		`json:"name"`
	InitialVersion	float32		`json:"initialVersion"`
	Repositories	[] Repository	`json:"repositories"`
}


/**
 * Retrieves a project list
 */
func GetProjects(res http.ResponseWriter, req *http.Request) {
	projects := [] Project {}
	api.JsonResponse(res, projects)
}


/**
 * Creates a new project
 */
func PostProjects(res http.ResponseWriter, req *http.Request) {
}


/**
 * Retrieves a single project
 */
func GetProject(res http.ResponseWriter, req *http.Request) {
}


/**
 * Updates a single project
 */
func PutProject(res http.ResponseWriter, req *http.Request) {
}


/**
 * Removes a single project
 */
func DeleteProject(res http.ResponseWriter, req *http.Request) {
}