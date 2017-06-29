package resource


import (
	"net/http"
	api "semverce/http"
)


type Repository struct {
	Id	string	`json:"id"`
	Name 	string 	`json:"name"`
	Url 	string 	`json:"url"`
	Key	string 	`json:"key"`
}


/**
 * Retrieves the repository list
 */
func GetRepositories(res http.ResponseWriter, req *http.Request) {
	repositories := [] Repository {}
	api.JsonResponse(res, repositories)
}


/**
 * Creates a new repository
 */
func PostRepositories(res http.ResponseWriter, req *http.Request) {
}


/**
 * Retrieves a single repository
 */
func GetRepository(res http.ResponseWriter, req *http.Request) {
}


/**
 * Updates a single repository
 */
func PutRepository(res http.ResponseWriter, req *http.Request) {
}


/**
 * Removes a single repository
 */
func DeleteRepository(res http.ResponseWriter, req *http.Request) {
}


/**
 * Creates a new major version of a single repository
 */
func PostMajorRepository(res http.ResponseWriter, req *http.Request) {
}


/**
 * Creates a new minor version of a single repository
 */
func PostMinorRepository(res http.ResponseWriter, req *http.Request) {
}


/**
 * Creates a new fix version of a single repository
 */
func PostFixRepository(res http.ResponseWriter, req *http.Request) {
}