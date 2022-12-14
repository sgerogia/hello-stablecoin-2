const { decryptEth } = require("../util")

task("4-get-auth-request", "Receive the latest AuthRequest message for PGBP")
    .addParam("contract", "The PGBP contract address")
    .addParam("account", "The payer's account address which initiated the flow (used for filtering)")
    .setAction(async (taskArgs) => {

    const contractAddr = taskArgs.contract
    const account = taskArgs.account
    const networkId = network.name

    console.log("Listening for AuthRequest for PGBP (contract", contractAddr, ") on network", networkId)
    const ProvableGBP = await ethers.getContractFactory("ProvableGBP")

    //Get signer information
    const signer = new ethers.Wallet(process.env.PRIVATE_KEY, ethers.provider)

    //Create connection to Contract and filter for events
    const gbpContract = new ethers.Contract(contractAddr, ProvableGBP.interface, signer)

    const eventFilter = gbpContract.filters.AuthRequest(account)
    const events = await gbpContract.queryFilter(eventFilter)
    // exit early
    if (events.length == 0) {
        console.log("No AuthRequest events yet. Try again in a few minutes...")
        return
    }

    const event = events[0].args
    console.log("AuthRequest received");
    console.log("\tRequestId", event.requestId);
    console.log("\tRequester", event.requester);
    console.log("\tEncr. data", ethers.utils.toUtf8String(event.authEncryptedData));

    const decryptedData = await decryptEth(signer.privateKey, ethers.utils.toUtf8String(event.authEncryptedData))
    const rawData = JSON.parse(decryptedData)

    const url = rawData.url

    console.log("--------------------------------------------------")
    console.log("Use the following URL to grant the consent")
    console.log("URL:", url)
  })

module.exports = {}