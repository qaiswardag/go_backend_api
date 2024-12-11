package passwordresetcontroller

import (
	"fmt"
	"net/http"
)

/*
|--------------------------------------------------------------------------
| Controller Method Naming Convention
|--------------------------------------------------------------------------
|
| Controller methods: index, create, store, show, edit, update, destroy.
| Please aim for consistency by using these method names in all controllers.
|
*/
func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Password Reset")
}
