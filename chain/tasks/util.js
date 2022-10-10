const { ethers } = require("ethers");
const EthCrypto = require('eth-crypto');
const { stripHexPrefix } = require('ethereumjs-util');

const encryptEth = async (publicKey, data) => {
    const pKey = stripHexPrefix(publicKey).slice(2);

    const encr = await EthCrypto.encryptWithPublicKey(pKey, JSON.stringify(data));
    return JSON.stringify(encr);
}

const decryptEth = async (privateKey, encrData) => {
    const encr = JSON.parse(encrData);
    const decr = await EthCrypto.decryptWithPrivateKey(privateKey, encr);
    return decr;
}

module.exports = {
    encryptEth,
    decryptEth,
}
