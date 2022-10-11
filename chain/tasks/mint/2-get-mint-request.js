const { decryptEth } = require("../util")

task("2-get-mint-request", "Receive the latest MintRequest message for PGBP")
    .addParam("contract", "The PGBP contract address")
    .addParam("account", "The account address which sent the message (used for filtering)")
    .setAction(async (taskArgs) => {

    const contractAddr = taskArgs.contract
    const account = ethers.utils.getAddress(taskArgs.account)
    const networkId = network.name

    console.log("Listening for MintRequest for PGBP (contract", contractAddr, ") on network", networkId)
    const ProvableGBP = await ethers.getContractFactory("ProvableGBP")

    //Get signer information
    const signer = new ethers.Wallet(process.env.PRIVATE_KEY, ethers.provider)

    //Create connection to Contract and filter for events
    const gbpContract = new ethers.Contract(contractAddr, ProvableGBP.interface, signer)

    const eventFilter = gbpContract.filters.MintRequest(account)
    const events = await gbpContract.queryFilter(eventFilter)
    // exit early
    if (events.length == 0) {
        console.log("No MintRequest events yet. Try again in a few minutes...")
        return
    }

    const event = events[0].args
    console.log("MintRequest received");
    console.log("\tRequestId", event.requestId);
    console.log("\tRequester", event.requester);
    console.log("\tAmount", event.amount);
    console.log("\tExpiration", event.expiration);
    console.log("\tEncr. data", ethers.utils.toUtf8String(event.encryptedData));

    const decryptedData = await decryptEth(signer.privateKey, ethers.utils.toUtf8String(event.encryptedData))
    // doing something really silly to need double JSON parsing; should have been obvious but isn't...
    const rawData = JSON.parse(JSON.parse(decryptedData))

    const requester = event.requester
    const amount = ethers.utils.formatEther(event.amount)
    const institutionId = rawData.institutionId
    const name = rawData.name
    const sortCode = rawData.sortCode
    const accountNumber = rawData.accountNumber

    console.log("--------------------------------------------------")
    console.log("Use the following to create a consent\n")
    console.log(`
    curl -k -X POST \
      https://ob.sandbox.natwest.com/open-banking/v3.1/pisp/domestic-payment-consents \
      -H 'Authorization: Bearer ACCESS_TOKEN' \\
      -H 'Content-Type: application/json' \\
      -H 'x-idempotency-key: IDEMPOTENCY_KEY' \\
      -H 'x-jws-signature: DUMMY_SIG' \\
      -d '{
      "Data": {
        "Initiation": {
          "InstructionIdentification": "instr-identification",
          "EndToEndIdentification": "e2e-identification",
          "CreditorAccount": {
            "SchemeName": "SortCodeAccountNumber",
            "Identification": "${sortCode}${accountNumber}",
            "Name": "${name}",
          },
          "InstructedAmount": {
            "Amount": "${amount}",
            "Currency": "GBP"
          },
          "CreditorAccount": {
            "SchemeName": "SortCodeAccountNumber",
            "Identification": "${sortCode}${accountNumber}",
            "Name": "Provable Ltd",
          },
          "RemittanceInformation": {
            "Unstructured": "Provable GBP",
            "Reference": "Provable GBP"
          }
        }
      },
      "Risk": {
        "PaymentContextCode": "Services",
        "MerchantCategoryCode": null,
        "MerchantCustomerIdentification": null,
        "DeliveryAddress": null
      }
    }'`)
  })

module.exports = {}