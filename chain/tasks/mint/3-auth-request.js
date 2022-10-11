const { encryptEth, recoverPublicKey } = require("../util")

task("3-auth-request", "Triggers an authRequest for PGBP, in response to a mintRequest")
  .addParam("contract", "The PGBP contract address")
  .addParam("requestId", "The original mint-request requestId")
  .addParam("url", "The user authentication url to encrypt")
  .setAction(async (taskArgs) => {

    const contractAddr = taskArgs.contract
    const networkId = network.name
    const requestId = taskArgs.requestId
    const url = taskArgs.url

    console.log("Initiating authRequest for RequestId", requestId, "PGBP (contract", contractAddr, ") on network", networkId)
    const ProvableGBP = await ethers.getContractFactory("ProvableGBP")

    //Get signer information
    const [ signer ] = await ethers.getSigners()

    //Create connection to Contract and call the getter function
    const gbpContract = new ethers.Contract(contractAddr, ProvableGBP.interface, signer)

    // Find MintRequest message
    const eventFilter = gbpContract.filters.MintRequest(null, requestId)
    const events = await gbpContract.queryFilter(eventFilter)
    // exit early
    if (events.length == 0) {
        console.log("No MintRequest events yet. Try again in a few minutes...")
        return
    }
    const trx = await events[0].getTransaction()
    const event = events[0].args
    // ... and recover sender's public key
    const requesterPublicKey = await recoverPublicKey(trx)

    const payload = {
        url: url
    }
    const encryptedData = await encryptEth(requesterPublicKey, JSON.stringify(payload))

    await gbpContract.authRequest(requestId, ethers.utils.toUtf8Bytes(encryptedData))
    console.log("Done")
  })

module.exports = {}