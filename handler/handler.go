package handler
import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/render"
    "github.com/juliocesar1235/bank-api/db"
)
var dbInstance db.Database
func NewHandler(db db.Database) http.Handler {
    router := chi.NewRouter()
    dbInstance = db
    router.MethodNotAllowed(methodNotAllowedHandler)
    router.NotFound(notFoundHandler)
    router.Route("/transactions", transactions)
    return router
}
func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")
    w.WriteHeader(405)
    render.Render(w, r, ErrMethodNotAllowed)
}
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")
    w.WriteHeader(400)
    render.Render(w, r, ErrNotFound)
}