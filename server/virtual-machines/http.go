package virtualmachines

import (
	"architecture-lab-3/server/tools"
	"encoding/json"
	"log"
	"net/http"
)

type HttpHandlerFunc http.HandlerFunc

func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleVirtualMachines(store, rw)
		} else if r.Method == "POST" {
			handleDiskAdd(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleDiskAdd(r *http.Request, rw http.ResponseWriter, store *Store) {
	var disk MachineDisk
	if err := json.NewDecoder(r.Body).Decode(&disk); err != nil {
		log.Printf("Error decoding input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.AddDiskToMachine(disk.DiskId, disk.MachineId)
	if err == nil {
		tools.WriteJsonOk(rw, &disk)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}

func handleVirtualMachines(store *Store, rw http.ResponseWriter) {
	res, err := store.VirtualMachinesList()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}
