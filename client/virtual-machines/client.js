const http = require('../common/http.js');

const Client = (baseUrl) => {
    const client = http.Client(baseUrl);

    return {
        listVirtualMachines: () => client.get('/virtualmachines'),
        connectDiskToMachine: (disk_id, machine_id) => client.post('/virtualmachines', { disk_id, machine_id })
    }
};

module.exports = { Client };