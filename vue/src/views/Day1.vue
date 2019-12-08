<template>
  <v-container>
    <v-row justify="space-around">
      <v-col cols="12" sm="6">
        <v-btn
          rounded
          block
          color="primary"
          class="mt-5"
          @click="parse"
        >Load Data</v-btn>
      </v-col>
    </v-row>
    <v-row justify="space-around">
      <v-col cols="12" sm="5">
        <part-one
          class="mt-5"
          :modules="modules"
        ></part-one>
      </v-col>
      <v-col cols="12" sm="5">
        <part-two
          class="mt-5"
          :modules="modules"
        ></part-two>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
// eslint-disable-next-line import/no-webpack-loader-syntax
import data from 'raw-loader!../input/day1/input.txt';
import partOne from '../components/day1/part1.vue';
import partTwo from '../components/day1/part2.vue';

export default {
  components: {
    partOne,
    partTwo,
  },
  data: () => ({
    modules: [],
  }),
  methods: {
    parse() {
      this.modules = [];
      data.split('\n').forEach((line) => {
        const mass = parseInt(line, 10);
        this.modules.push(mass);
      });
    },
  },

};
</script>
