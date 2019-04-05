<template>
  <v-layout row>
    <v-flex sm12>
      <jumbotron/>
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
