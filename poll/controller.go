package poll

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

type Storer interface {
	Create(p Poll)
	Delete(id int)
	GetOne(id int) Poll
	GetMany() []Poll
	Dump() interface{}
}

type Controller struct {
	storer Storer
}

func (pc *Controller) GetMany(w http.ResponseWriter, r *http.Request) {
	polls := pc.storer.GetMany()
	names := []string{}
	for _, poll := range polls {
		names = append(names, strconv.Itoa(poll.Id))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strings.Join(names, ", ")))
	fmt.Println(pc.storer.Dump())
}

func (pc *Controller) GetOne(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idParam)
	poll := pc.storer.GetOne(id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(poll.Title))
	fmt.Println(pc.storer.Dump())
}

func (pc *Controller) Post(w http.ResponseWriter, r *http.Request) {
	pc.storer.Create(Poll{
		Title: "Just another poll",
	})
	w.WriteHeader(http.StatusCreated)
	fmt.Println(pc.storer.Dump())
}

func (pc *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idParam)
	pc.storer.Delete(id)
	w.WriteHeader(http.StatusOK)
	fmt.Println(pc.storer.Dump())
}

func NewController(storer Storer) *Controller {
	return &Controller{storer}
}
