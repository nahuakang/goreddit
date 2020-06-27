package web

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/gorilla/csrf"
	"github.com/nahuakang/goreddit"
)

// ThreadHandler handles threads
type ThreadHandler struct {
	store goreddit.Store
}

// List returns a webpage with the list of all Threads
func (h *ThreadHandler) List() http.HandlerFunc {
	type data struct {
		Threads []goreddit.Thread
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/threads.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		tt, err := h.store.Threads()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, data{Threads: tt})
	}
}

// Create leads to the page for creating new threads
func (h *ThreadHandler) Create() http.HandlerFunc {
	type data struct {
		CSRF template.HTML
	}
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/thread_create.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data{
			CSRF: csrf.TemplateField(r),
		})
	}
}

// Show shows all the threads
func (h *ThreadHandler) Show() http.HandlerFunc {
	type data struct {
		CSRF   template.HTML
		Thread goreddit.Thread
		Posts  []goreddit.Post
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/thread.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		// convert idStr to UUID first
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		t, err := h.store.Thread(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		pp, err := h.store.PostsByThread(t.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, data{
			CSRF:   csrf.TemplateField(r),
			Thread: t,
			Posts:  pp,
		})
	}
}

// Store saves the newly created thread to database
func (h *ThreadHandler) Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		description := r.FormValue("description")

		if err := h.store.CreateThread(&goreddit.Thread{
			ID:          uuid.New(),
			Title:       title,
			Description: description,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/threads", http.StatusFound)
	}
}

// Delete deletes a thread based on its id
func (h *ThreadHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.store.DeleteThread(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/threads", http.StatusFound)
	}
}
