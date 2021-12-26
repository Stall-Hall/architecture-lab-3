const http = require('../common/http.js');

const Client = (baseUrl) => {
    const client = http.Client(baseUrl);

    return {
        listVirtualMachines: () => client.get('/virtual machines'),
        connectDiskToMachine: (disk_id, machine_id) => client.post('/virtual machines', { disk_id, machine_id })
    }
};

module.exports = { Client };