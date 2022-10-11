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

// Recover the public key of the sender given a transaction.
const recoverPublicKey = async (trx) => {

    const expandedSig = {
      r: trx.r,
      s: trx.s,
      v: trx.v
    }
    const signature = ethers.utils.joinSignature(expandedSig)
    const txData = {
      gasPrice: trx.gasPrice,
      gasLimit: trx.gasLimit,
      value: trx.value,
      nonce: trx.nonce,
      data: trx.data,
      chainId: trx.chainId,
      to: trx.to // you might need to include this if it's a regular tx and not simply a contract deployment
    }
    const rsTx = await ethers.utils.resolveProperties(txData)
    const raw = ethers.utils.serializeTransaction(rsTx) // returns RLP encoded tx
    const msgHash = ethers.utils.keccak256(raw) // as specified by ECDSA
    const msgBytes = ethers.utils.arrayify(msgHash) // create binary hash
    const recoveredPubKey = ethers.utils.recoverPublicKey(msgBytes, signature)

    return recoveredPubKey
}

module.exports = {
    encryptEth,
    decryptEth,
    recoverPublicKey,
}
