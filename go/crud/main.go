// main.go
// Sederhana CRUD aplikasi web menggunakan Go dan net/http
// Mengelola data Item secara in-memory (map)
// Menyajikan halaman HTML untuk antarmuka pengguna

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Item merepresentasikan data yang akan kita kelola
type Item struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ItemStore menyimpan koleksi Item dengan mutex untuk akses concurrent
type ItemStore struct {
	items map[string]Item
	mu    sync.RWMutex
}

// NewItemStore membuat instance ItemStore baru
func NewItemStore() *ItemStore {
	return &ItemStore{
		items: make(map[string]Item),
	}
}

// GetAll mengembalikan semua Item
func (s *ItemStore) GetAll() []Item {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]Item, 0, len(s.items))
	for _, item := range s.items {
		list = append(list, item)
	}
	return list
}

// Get mengembalikan Item berdasarkan ID
func (s *ItemStore) Get(id string) (Item, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	item, exists := s.items[id]
	return item, exists
}

// Create menambahkan Item baru
func (s *ItemStore) Create(item Item) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[item.ID] = item
}

// Update memperbarui Item yang sudah ada
func (s *ItemStore) Update(id string, item Item) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.items[id]; !exists {
		return false
	}
	s.items[id] = item
	return true
}

// Delete menghapus Item berdasarkan ID
func (s *ItemStore) Delete(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.items[id]; !exists {
		return false
	}
	delete(s.items, id)
	return true
}

// Handler untuk menangani request HTTP
type Handler struct {
	store *ItemStore
}

// NewHandler membuat Handler baru
func NewHandler(store *ItemStore) *Handler {
	return &Handler{store: store}
}

// respondJSON mengirimkan respons JSON dengan status code
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Error encoding JSON: %v", err)
	}
}

// errorResponse mengirimkan pesan error JSON
func errorResponse(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}

// getItemID meng-extract ID dari URL path
func getItemID(r *http.Request) string {
	// format path: /items/{id}
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")
	if len(parts) < 2 {
		return ""
	}
	return parts[1] // parts[0] is "items", parts[1] is {id}
}

// Handler methods

// ListItems menangani GET /items
func (h *Handler) ListItems(w http.ResponseWriter, r *http.Request) {
	items := h.store.GetAll()
	respondJSON(w, http.StatusOK, items)
}

// GetItem menangani GET /items/{id}
func (h *Handler) GetItem(w http.ResponseWriter, r *http.Request) {
	id := getItemID(r)
	if id == "" {
		errorResponse(w, http.StatusBadRequest, "Missing item ID")
		return
	}
	item, exists := h.store.Get(id)
	if !exists {
		errorResponse(w, http.StatusNotFound, "Item not found")
		return
	}
	respondJSON(w, http.StatusOK, item)
}

// CreateItem menangani POST /items
func (h *Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		errorResponse(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	// Generate simple ID jika tidak disediakan
	if item.ID == "" {
		item.ID = strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	h.store.Create(item)
	respondJSON(w, http.StatusCreated, item)
}

// UpdateItem menangani PUT /items/{id}
func (h *Handler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	id := getItemID(r)
	if id == "" {
		errorResponse(w, http.StatusBadRequest, "Missing item ID")
		return
	}
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		errorResponse(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	item.ID = id // pastikan ID sesuai URL
	if !h.store.Update(id, item) {
		errorResponse(w, http.StatusNotFound, "Item not found")
		return
	}
	respondJSON(w, http.StatusOK, item)
}

// DeleteItem menangani DELETE /items/{id}
func (h *Handler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	id := getItemID(r)
	if id == "" {
		errorResponse(w, http.StatusBadRequest, "Missing item ID")
		return
	}
	if !h.store.Delete(id) {
		errorResponse(w, http.StatusNotFound, "Item not found")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// serveHTML menangani permintaan untuk halaman HTML utama
func serveHTML(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// Jika bukan root path, coba sertakan sebagai file statis
		http.ServeFile(w, r, "static"+r.URL.Path)
		return
	}
	http.ServeFile(w, r, "static/index.html")
}

// main fungsi entry point
func main() {
	store := NewItemStore()
	handler := NewHandler(store)

	// Route definitions untuk API
	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.ListItems(w, r)
		case http.MethodPost:
			handler.CreateItem(w, r)
		default:
			errorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})

	http.HandleFunc("/items/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetItem(w, r)
		case http.MethodPut:
			handler.UpdateItem(w, r)
		case http.MethodDelete:
			handler.DeleteItem(w, r)
		default:
			errorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})

	// Route untuk halaman HTML dan file statis
	http.HandleFunc("/", serveHTML)

	fmt.Println("Server berjalan di http://localhost:8080")
	fmt.Println("Buka browser dan akses http://localhost:8080 untuk melihat antarmuka")
	log.Fatal(http.ListenAndServe(":8080", nil))
}