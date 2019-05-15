<template>
  <v-stepper v-model="idx">
    <v-stepper-header>
      <v-stepper-step :complete="metaMaskOK" step="1">Metamask installed</v-stepper-step>

      <v-divider></v-divider>

      <v-stepper-step :complete="chainOK" step="2">Connected to network: {{ chainID }}</v-stepper-step>

      <v-divider></v-divider>

      <v-stepper-step :complete="accountOK" step="3">Account {{ defaultAccount }}</v-stepper-step>
    </v-stepper-header>
  </v-stepper>
</template>

<script>
//import {crea}
import { createNamespacedHelpers } from "vuex";
//import { mapState } from "vuex";
const { mapState } = createNamespacedHelpers("web3");

export default {
  data() {
    return {
      idx: 2
    };
  },
  computed: {
    ...mapState(["chainID", "defaultAccount"]),
    metaMaskOK() {
      return typeof web3 !== "undefined";
    },
    chainOK() {
      return !!(this.metaMaskOK && this.chainID);
    },
    accountOK() {
      return !!(this.chainOK && this.defaultAccount);
    }
  },
  methods: {},
  async mounted() {
    await this.$store.dispatch("web3/registerWeb3");
  }
};
</script>

