package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Coldwws/todo/internal/models"
)

type RoomHandler struct {
	rooms []models.Room
}

func NewRoomHandler() *RoomHandler {
	return &RoomHandler{
		rooms: make([]models.Room, 0),
	}
}

func (h *RoomHandler) GetRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	json.NewEncoder(w).Encode(h.rooms)
}

func (h *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var room models.Room
	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for _, v := range h.rooms {
		if v.Number == room.Number {
			http.Error(w, "Room with this number already exists", http.StatusConflict)
			return
		}
	}

	room.ID = len(h.rooms) + 1

	h.rooms = append(h.rooms, room)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(room)
}

func (h *RoomHandler) UpdateRoom(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	var updateRoom models.UpdateRoom
	json.NewDecoder(r.Body).Decode(&updateRoom)

	found := false
	for i, room := range h.rooms {
		if room.ID == id {
			if updateRoom.Number != nil {
				h.rooms[i].Number = *updateRoom.Number
			}
			if updateRoom.Type != nil {
				h.rooms[i].Type = *updateRoom.Type
			}
			if updateRoom.Description != nil {
				h.rooms[i].Description = *updateRoom.Description
			}
			found = true
			json.NewEncoder(w).Encode(h.rooms[i])
			break
		}
	}
	if !found{
		http.Error(w,"Room not found",http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *RoomHandler) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	found := false
	for i, room := range h.rooms {
		if room.ID == id {
			h.rooms = append(h.rooms[:i], h.rooms[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		http.Error(w, "Room not found", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Room with ID %d deleted", id)
}
