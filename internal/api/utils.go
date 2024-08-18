package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/JhonyMesquita/Go-Heights/internal/store/pgstore"
)

func (h apiHandler) readRoom(w http.ResponseWriter, r *http.Request) (pgstore.Room, error) {
	var room pgstore.Room
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		slog.Error("Error decoding room", err)
		http.Error(w, "Error decoding room", http.StatusBadRequest)
		return room, err
	}
	return room, nil
}
