const virtualMachines = require('./virtual-machines/client');

const client = virtualMachines.Client('http://localhost:8080');

client.listVirtualMachines()
    .then((list) => {
        console.log('=== Scenario 1 ===');
        console.log('Available virtual machines:');
        console.log(list);
    })
    .catch((err) => {
        console.log(`Error while listing available virtual machines: ${err.message}`);
    });

client.connectDiskToMachine(1, 2)
    .then((resp) => {
        console.log('=== Scenario 2 ===');
        console.log('Connect disk to virtual machine response:', resp);
    })
    .catch((err) => {
        console.log(`Error while connecting disk to virtual machine: ${err.message}`);
    });

client.connectDiskToMachine(7, 1)
    .then((resp) => {
        console.log('=== Scenario 3 ===');
        console.log('Connect disk to virtual machine response:', resp);
    })
    .catch((err) => {
        console.log(`Error while connecting disk to virtual machine: ${err.message}`);
    });