#! /usr/bin/make

compile-contract:
	cd chain && nvm use 18.7.0 && npx hardhat compile

