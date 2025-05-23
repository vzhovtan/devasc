package main

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

// Map of usernames to passwords.
var login = map[string]string{
	"admin": "password123",
}

// Rconf struct to hold Restconf default data.
type Rconf struct {
	Restconf           string `json:"ietf-restconf:restconf,omitempty"`
	Data               string `json:"data,omitempty"`
	Operations         string `json:"operations,omitempty"`
	YangLibraryVersion string `json:"yang-library-version"`
}

// Global var rconf to store restconf default data.
var rconf = Rconf{Restconf: " ", Data: "{}", Operations: "{}", YangLibraryVersion: "2016-06-21"}

// Inter struct to hold Restconf interface data
type Inter struct {
	IetfInterfacesInterfaces struct {
		Interface []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Type        string `json:"type"`
			Enabled     bool   `json:"enabled"`
			IetfIpIpv4  struct {
			} `json:"ietf-ip:ipv4"`
			IetfIpIpv6 struct {
			} `json:"ietf-ip:ipv6"`
		} `json:"interface"`
	} `json:"ietf-interfaces:interfaces"`
}

// Global var inter to store restconf interfaces data.
var inter = []Inter{
	{
		IetfInterfacesInterfaces: struct {
			Interface []struct {
				Name        string `json:"name"`
				Description string `json:"description"`
				Type        string `json:"type"`
				Enabled     bool   `json:"enabled"`
				IetfIpIpv4  struct {
				} `json:"ietf-ip:ipv4"`
				IetfIpIpv6 struct {
				} `json:"ietf-ip:ipv6"`
			} `json:"interface"`
		}{
			Interface: []struct {
				Name        string `json:"name"`
				Description string `json:"description"`
				Type        string `json:"type"`
				Enabled     bool   `json:"enabled"`
				IetfIpIpv4  struct {
				} `json:"ietf-ip:ipv4"`
				IetfIpIpv6 struct {
				} `json:"ietf-ip:ipv6"`
			}{
				{
					Name:        "GigabitEthernet1",
					Description: "Main Interface",
					Type:        "iana-if-type:ethernetCsmacd",
					Enabled:     true,
				},
				{
					Name:        "Loopback1",
					Description: "Loopback Interface",
					Type:        "iana-if-type:softwareLoopback",
					Enabled:     true,
				},
			},
		},
	},
}

// Get default Restconf response
func GetDefault(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rconf)
}

// Get Restconf All Interfaces data
func GetInterfaces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inter)
}

// Get a single interface by name.
func GetInterface(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["interface"]
	name := strings.Split(id, "=")[1]
	if name != "GigabitEthernet1" && name != "Loopback1" {
		http.Error(w, "Invalid Interface Name", http.StatusBadRequest)
		return
	}

	for _, intf := range inter {
		for _, specificInt := range intf.IetfInterfacesInterfaces.Interface {
			if specificInt.Name == "GigabitEthernet1" {
				json.NewEncoder(w).Encode(specificInt)
				return
			}
		}
	}

	http.Error(w, "Interface not found", http.StatusNotFound)
}

// Mocking for system authentication
func basicAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || parts[0] != "Basic" {
			http.Error(w, "Malformed Authorization header", http.StatusBadRequest)
			return
		}

		decoded, err := base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			http.Error(w, "Invalid base64 encoding", http.StatusBadRequest)
			return
		}

		creds := strings.SplitN(string(decoded), ":", 2)
		if len(creds) != 2 {
			http.Error(w, "Invalid credentials format", http.StatusBadRequest)
			return
		}

		username, password := creds[0], creds[1]

		if storedPassword, ok := login[username]; ok && storedPassword == password {
			// Authentication successful, call the actual handler
			handler(w, r)
			return
		}

		// Authentication failed
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

func main() {
	// Create a new router using gorilla/mux.
	r := mux.NewRouter()

	// Define API endpoints.  Use constants for the paths.
	const restconfPath = "/restconf"
	const restconfPathIntfs = "/restconf/data/ietf-interfaces:interfaces"
	const restconfPathInterface = "/restconf/data/ietf-interfaces:interfaces/{interface}"
	// Register the handlers.  Use method chaining for cleaner syntax.
	r.HandleFunc(restconfPath, basicAuth(GetDefault)).Methods(http.MethodGet)
	r.HandleFunc(restconfPathIntfs, basicAuth(GetInterfaces)).Methods(http.MethodGet)
	r.HandleFunc(restconfPathInterface, basicAuth(GetInterface)).Methods(http.MethodGet)

	// Start the server.
	const port = ":8080"
	log.Printf("Server listening on port %s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
