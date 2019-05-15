<template>
  <v-stepper-content step="2">
    <v-card>
      <v-card-text>
        <v-form>
          <v-select v-model="deed" :items="deeds" label="Asset"></v-select>
          <v-btn
            :disabled="deed==null || transfering"
            @click="transfer()"
            outline
            color="teal"
          >Transfer Ownership</v-btn>
        </v-form>
      </v-card-text>
      <!--
      <v-layout xs12 row wrap>
        <v-flex style="height:100%; padding-bottom:20px;" xs12 sm12 md12>
          <v-flex style="height:100%; padding-bottom:20px;" xs12 sm12 md12>
            <v-select v-model="selectedDeed" :items="deeds" label="Asset"></v-select>
          </v-flex>
          <div v-show="transferDeedSuccess">
            <v-icon color="teal">check_circle</v-icon>
            <h3 style="color:green !important;">Ownership was successfully transfered</h3>
          </div>

          <div v-show="transferingDeed">
            <h3>Please wait while transfering ownership</h3>
            <v-progress-linear v-bind:indeterminate="true"></v-progress-linear>
          </div>
          <v-btn
            :disabled="selectedDeed==null || transferDeedSuccess || transferingDeed "
            @click="transferTo()"
            outline
            color="teal"
          >Transfer Ownership</v-btn>
          <v-btn v-show="transferDeedSuccess" outline color="primary" @click="stepperIndex++">Next</v-btn>
        </v-flex>
      </v-layout>
      -->
    </v-card>
    <v-btn color="primary">Continue</v-btn>
    <v-btn flat>Cancel</v-btn>
  </v-stepper-content>
</template>

<script>
export default {
  data() {
    return {
      deed: null,
      deeds: [],
      transfering: false,
      transferred: false
    };
  },
  methods: {
    transfer() {}
  }
};
</script>

