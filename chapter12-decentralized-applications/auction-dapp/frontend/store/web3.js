import Web3 from "web3"

export const state = () => ({
  dialer: null,
  chainID: null,
  defaultAccount: null,
})

export const mutations = {
  setDefaultAccount(state, a) {
    state.defaultAccount = a
  },
  onWeb3(state, { dialer, chainID }) {
    state.dialer = dialer
    state.chainID = chainID
  },
}

export const actions = {
  async registerWeb3({ commit }) {
    const web3js = window.web3
    if (typeof web3js !== "undefined") {
      const web3 = new Web3(web3js.currentProvider)

      let opts = {
        dialer() {
          return web3
        },
      }

      opts.chainID = await web3.eth.getChainId()

      setInterval(async () => {
        const accounts = await web3.eth.getAccounts()
        if (accounts.length > 0) {
          commit("setDefaultAccount", accounts[0])
        }
      }, 3000)

      commit("onWeb3", opts)
    }
  },
}
