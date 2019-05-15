<template>
  <v-layout row wrap>
    <v-flex sm12>
      <jumbotron/>
    </v-flex>
    <v-flex sm12>
      <v-divider></v-divider>
      <h1 class="text-sm-center">Recent Auctions</h1>
    </v-flex>
    <v-flex sm12>
      <v-layout justify-center>
        <v-btn color="teal" outline>Load More Auctions</v-btn>
      </v-layout>
    </v-flex>
  </v-layout>
</template>

<script>
import Jumbotron from "~/components/Jumbotron";

export default {
  layout: "dapp",
  components: { Jumbotron }
};
</script>
