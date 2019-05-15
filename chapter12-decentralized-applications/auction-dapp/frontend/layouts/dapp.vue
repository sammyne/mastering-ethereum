<template>
  <v-app>
    <v-toolbar>
      <v-toolbar-title>Auction Dapp</v-toolbar-title>
      <v-spacer></v-spacer>
      <!--
      <v-btn outline>
        <v-icon left>mdi-plus</v-icon>New Auction
      </v-btn>
      -->
      <new-auction-panel/>
    </v-toolbar>
    <v-content>
      <v-container>
        <nuxt/>
      </v-container>
    </v-content>

    <!--<meta-mask/>-->
  </v-app>
</template>

<script>
//import MetaMask from "~/components/MetaMask";
import NewAuctionPanel from "~/components/NewAuctionPanel";

export default {
  //components: { MetaMask }
  components: { NewAuctionPanel }
};
</script>

