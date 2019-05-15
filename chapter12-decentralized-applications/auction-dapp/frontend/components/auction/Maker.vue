<template>
  <v-stepper-content step="3">
    <v-card>
      <v-card-text>
        <v-form>
          <v-select :items="idVec" label="DeedID" v-model="id"></v-select>
          <v-text-field label="Auction title" placeholder="e.g. My NFT" required v-model="title"></v-text-field>
          <v-text-field label="Image(png/jpg)" required type="file" v-model="image"></v-text-field>
          <v-text-field
            label="Starting Price"
            placeholder="10 Ethers"
            required
            v-model="startingPrice"
          ></v-text-field>
          <v-menu
            ref="datePicker"
            v-model="datePicker"
            :close-on-content-click="false"
            :nudge-right="40"
            :return-value.sync="expiration"
            lazy
            transition="scale-transition"
            offset-y
            full-width
            min-width="290px"
          >
            <template v-slot:activator="{ on }">
              <v-text-field
                v-model="expiration"
                label="Picker in menu"
                prepend-icon="event"
                readonly
                v-on="on"
              ></v-text-field>
            </template>
            <v-date-picker
              color="green"
              v-model="expiration"
              no-title
              scrollable
              :allowed-dates="allowedExpiration"
            >
              <v-spacer></v-spacer>
              <v-btn flat color="warning" @click="datePicker = false">Cancel</v-btn>
              <v-btn flat color="primary" @click="$refs.datePicker.save(expiration)">OK</v-btn>
            </v-date-picker>
          </v-menu>
          <v-btn color="primary">Create Auction</v-btn>
        </v-form>
      </v-card-text>
    </v-card>
  </v-stepper-content>
</template>

<script>
export default {
  data() {
    return {
      datePicker: false,
      expiration: null,
      id: null,
      idVec: [],
      image: null,
      startingPrice: null,
      title: null
    };
  },
  computed: {
    expirationInUnix() {
      //return !!expiration ?
      //return Date.getTime()/1000+
    }
  },
  methods: {
    allowedExpiration(v) {
      return new Date(v).getTime() >= Date.now();
    }
  }
};
</script>

