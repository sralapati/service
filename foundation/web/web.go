package web

import (
	"context"
	"net/http"
	"os"
)

// HandlerFunc represents a function that handles a http request within our own
// little mini framework.
type Handler func(context.Context, http.ResponseWriter, *http.Request) error

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct.
type App struct {
	*http.ServeMux
	shutdown chan os.Signal
}

// NewApp creates an App value that handle a set of routes for the application.
func NewApp(shutdown chan os.Signal) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		shutdown: shutdown,
	}
}

func (a *App) HandleFunc(pattern string, handler Handler) {
	h := func(w http.ResponseWriter, r *http.Request) {
		if err := handler(r.Context(), w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	a.ServeMux.HandleFunc(pattern, h)
}
