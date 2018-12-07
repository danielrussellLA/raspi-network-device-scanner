{
    if [ ! -f ./readDevices ]; then
        echo 'compiling go service: readDevices.go...';
        go build -o readDevices readDevices.go;
        echo 'go compilation complete. created binary ./readDevices from ./readDevices.go';
    fi
    echo 'scanning for devices on current network...';
    ./scan_devices.sh >> devices.txt;
    echo 'scan finished';
    num_devices=($(./readDevices))
    echo 'devices on network: ';
    echo ${num_devices}
} || {
    echo 'Something went wrong';
    echo 'Make sure that you have arp-scan and golang installed on your raspberry pi:'
    echo '    sudo apt-get install arp-scan'
    echo '    sudo apt-get install golang'
}
