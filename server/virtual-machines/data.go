package virtualmachines

import (
	"database/sql"
)

type VirtualMachine struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	CpuCount       int64  `json:"cpuCount"`
	TotalDiskSpace int64  `json:"totalDiskSpace"`
}

type MachineDisk struct {
	DiskId    int64 `json:"disk_id"`
	MachineId int64 `json:"machine_id"`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) VirtualMachinesList() ([]*VirtualMachine, error) {
	var sqlQuery = `SELECT virtual_machines.id as id, name, "cpuCount", SUM(volume) as totalDiskSpace
										FROM virtual_machines
										INNER JOIN virtual_machine_discs 
											ON virtual_machines.id = virtual_machine_discs.vm_id
										INNER JOIN discs
											ON virtual_machine_discs.disk_id = discs.id
										GROUP BY virtual_machines.id
										ORDER BY virtual_machines.id ASC`
	rows, err := s.Db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*VirtualMachine
	for rows.Next() {
		var vm VirtualMachine
		if err := rows.Scan(&vm.Id, &vm.Name, &vm.CpuCount, &vm.TotalDiskSpace); err != nil {
			return nil, err
		}
		res = append(res, &vm)
	}
	if res == nil {
		res = make([]*VirtualMachine, 0)
	}
	return res, nil
}

func (s *Store) AddDiskToMachine(diskId int64, machineId int64) error {
	_, err := s.Db.Exec("INSERT INTO virtual_machine_discs (vm_id, disk_id) VALUES ($1, $2)", machineId, diskId)
	return err
}
