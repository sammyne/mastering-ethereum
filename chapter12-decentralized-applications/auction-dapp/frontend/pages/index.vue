<template>
  <v-layout row>
    <v-flex sm12>
      <v-card dark flat>
        <v-card-text>
          <v-layout row>
            <v-flex sm5>
              <h1>Auction Dapp</h1>
              <h2>Decentralized auction platform powered by Ethereum</h2>
            </v-flex>
            <v-flex sm6>
              <v-img src="~/assets/images/auction.svg"></v-img>
            </v-flex>
          </v-layout>
        </v-card-text>
      </v-card>
    </v-flex>
  </v-layout>
</template>

<script>
export default {
  layout: "dapp"
};
</script>
