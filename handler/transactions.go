package handler
import (
    "context"
    "fmt"
    "net/http"
    "github.com/go-chi/chi"
    "github.com/go-chi/render"
)
var accIDKey = "accountID"
func transactions(router chi.Router) {
		router.Route("/{accountId}", func(router chi.Router){
			router.Use(TransactionContext)
			router.Get("/", getTransactionsByAccId)
		})
}

func TransactionContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			accountId := chi.URLParam(r, "accountId")
			if accountId == "" {
					render.Render(w, r, ErrorRenderer(fmt.Errorf("accountId is required")))
					return
			}

			ctx := context.WithValue(r.Context(), accIDKey, accountId)
			next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getTransactionsByAccId(w http.ResponseWriter, r *http.Request) {
	accountId := r.Context().Value(accIDKey).(string)
	items, err := dbInstance.GetTransactionsByAccountId(accountId)
	if err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
	}
	if err := render.Render(w, r, items); err != nil {
			render.Render(w, r, ErrorRenderer(err))
	}
	
}