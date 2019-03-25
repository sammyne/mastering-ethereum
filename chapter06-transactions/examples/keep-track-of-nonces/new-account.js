const Web3 = require('web3')

// this endpoint is applied in INFURA
const INFURA = "https://ropsten.infura.io/v3/f3df74d615a74774821985274dedcc9e"

const web3 = new Web3(INFURA)
//console.log(web3.eth.personal)
//console.log(web3.personal)

//web3.eth.personal.newAccount('@hello-infura').then(console.log)
web3.eth.personal.getAccounts().then(console.log)