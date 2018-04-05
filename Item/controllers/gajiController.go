package controllers

import (
	"day14gaji/data"
	"encoding/json"
	"net/http"
)

func GetItem(w http.ResponseWriter, r *http.Request) {
	context := Context{}

	e := DBintitial(context.DB)
	defer e.Close()

	repo := &data.ItemRepositori{e}
	gaji := data.GetAll(repo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mdl, _ := json.Marshal(item)
	w.Write(mdl)
}
