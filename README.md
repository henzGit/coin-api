# coin-api
Experimental API to make bids to buy and sell coins
    
## Specs:
+ Written in Go 1.11
    
## Install dependencies
    dep ensure

## How to run:
    go run src/*
    
## How to test:
     curl -X POST -d @data.json -H "Content-Type: application/json" http://localhost:3000/bid-buy
