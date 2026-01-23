package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"` // Changed to lowercase to match common JSON conventions
	Stok  int    `json:"stok"`
}

var produk = []Product{
	{ID: 1, Nama: "Pensil", Harga: 2000, Stok: 100},
	{ID: 2, Nama: "Buku Tulis", Harga: 5000, Stok: 150},
	{ID: 3, Nama: "Penghapus", Harga: 1000, Stok: 200},
}

func main() {
	// 1. Health Check Handler
	// Fixed: Changed 'q' to 'w' and 'http.request' to 'http.Request'
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API is Running",
		})
	})

	// 2. General Product Handler (GET List, POST Create)
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {
			// Fixed: Typo 'Endcode' -> 'Encode'
			json.NewEncoder(w).Encode(produk)
			return // Important: Return to stop execution
		} else if r.Method == "POST" {
			var produkBaru Product
			err := json.NewDecoder(r.Body).Decode(&produkBaru)
			if err != nil {
				http.Error(w, "Invalid input", http.StatusBadRequest)
				return
			}

			// Logic moved INSIDE the POST block
			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)

			w.WriteHeader(http.StatusCreated) // 201
			json.NewEncoder(w).Encode(produkBaru)
			return
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	// 3. ID Specific Handler (Route Dispatcher)
	// This captures any URL starting with /api/produk/ (like /api/produk/1)
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getProductByID(w, r)
		case "PUT":
			updateProduk(w, r)
		case "DELETE":
			deleteProduk(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// 4. Start the Server
	fmt.Println("Server running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// --- Helper Functions ---

func getProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for _, p := range produk {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

func updateProduk(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var updateData Product
	err = json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	for i := range produk {
		if produk[i].ID == id {
			// Update fields
			produk[i].Nama = updateData.Nama
			produk[i].Harga = updateData.Harga
			produk[i].Stok = updateData.Stok
			// Keep the ID original
			updateData.ID = id

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk[i])
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

func deleteProduk(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for i, p := range produk {
		if p.ID == id {
			// Delete from slice
			produk = append(produk[:i], produk[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Product deleted successfully",
			})
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}
