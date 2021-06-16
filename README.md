# Crypto Server

# Run from From Command line
How to build
    1. From base folder, run "./build.sh"

How to Run
    1. Build first
    2. From base folder, run "nohup ./run/bin/crypto-server-app --conf run/cfg/dev/config.yaml &"

How to Test
    1. Use curl from command line
        $ curl -L http://host:port/api/v1/currency/ETHBTC
        {"id":"","fullName":"ETHBTC","Ask":0.063795,"Bid":0.06379,"Last":0.063815,"Open":0.06294,"Low":0.062782,"High":0.065119,"feeCurrency":""}
        $ curl -L http://host:port/api/v1/currency/BTCUSD
        {"id":"","fullName":"BTCUSD","Ask":40040.47,"Bid":40034.4,"Last":40061.41,"Open":40823.57,"Low":39354.59,"High":41038.09,"feeCurrency":""}
    2. From browser
        http://host:port/api/v1/currency/ETHBTC
        http://host:port/api/v1/currency/BTCUSD

# Run as Docker Container

    1. Build using build.sh
    2. Push to registry and run from registry
    3. Or Run as local container
