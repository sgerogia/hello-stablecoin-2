#! /usr/bin/make

build-contract:
	cd chain && nvm use 18.7.0 && npx hardhat compile

test-contract:
	cd chain && nvm use 18.7.0 && npx hardhat test

generate-abi: build-contract
	cd tpp-client && cat ../chain/artifacts/contracts/ProvableGBP.sol/ProvableGBP.json \
   		| jq '.abi' \
   		| abigen --abi - \
		  --type ProvableGBP \
		  --pkg contract \
		  --out contract/provable_gbp.go \
		  --bin ../chain/artifacts/contracts/ProvableGBP.bin

test-client:
	cd tpp-client && go test ./...

build-client: generate-abi
	cd tpp-client && go build -o tpp ./cmd && chmod +x ./tpp