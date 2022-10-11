const { getNamedAccounts, deployments, network } = require("hardhat")

module.exports = async ({ getNamedAccounts, deployments }) => {
  const { deploy, log, get } = deployments
  const chainId = network.config.chainId
  //set log level to ignore non errors
  ethers.utils.Logger.setLogLevel(ethers.utils.Logger.levels.ERROR)

  const deployer = new ethers.Wallet(process.env.PRIVATE_KEY, ethers.provider);

  const args = [ deployer.publicKey ]
  const gbp = await deploy("ProvableGBP", {
    from: deployer.address,
    args: args,
    log: true,
    waitConfirmations: 1,
  })


  log("Deployed!")
  log("----------------------------------------------------")
  log("Run the ProvableGBP contract with the following commands:")
  const networkName = network.name == "hardhat" ? "localhost" : network.name
  log(`npx hardhat mint-request \\
    --contract ${gbp.address} \\
    --network ${networkName} \\
    --amount <amount to mint> \\
    --institution-id <Yapily API institution id> \\
    --sort-code <payer sort code> \\
    --account-number <payer account number> \\
    --name <account holder full name>`)
  log(`npx hardhat get-mint-request \\
    --contract ${gbp.address} \\
    --network ${networkName} \\
    --account <payer ETH address>`)
  log(`npx hardhat read-data --contract ${gbp.address} --network ${networkName}`)
  log("----------------------------------------------------")
}
module.exports.tags = ["all", "gbp"]
