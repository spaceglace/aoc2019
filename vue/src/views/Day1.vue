<template>
  <v-container>
    <v-row justify="space-around">
      <v-col cols="12" sm="5">
        <v-btn class="mt-5" @click="parse">Load Data</v-btn>
        <div class="title font-weight-light mt-5">
          Total module count: <span class="font-weight-bold">{{ moduleCount }}</span>
        </div>
        <div class="title font-weight-light mt-3">
          Total fuel: <span class="font-weight-bold">{{ fuelSum }}</span>
        </div>
      </v-col>
      <v-col cols="12" sm="5">
        <v-data-table
          :headers="headers"
          :items="modules"
        ></v-data-table>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
// eslint-disable-next-line import/no-webpack-loader-syntax
import data from 'raw-loader!../input/day1/input.txt';

export default {
  data: () => ({
    modules: [],
    headers: [
      {
        text: 'Module #',
        align: 'left',
        value: 'index',
      },
      {
        text: 'Mass',
        align: 'left',
        value: 'mass',
      },
      {
        text: 'Needed Fuel',
        align: 'left',
        value: 'fuel',
      },
    ],
  }),
  computed: {
    moduleCount() {
      return this.modules.length;
    },
    fuelSum() {
      return this.modules.reduce((acc, cur) => acc + cur.fuel, 0);
    },
  },
  methods: {
    calculateFuel(mass) {
      return Math.floor(mass / 3) - 2;
    },
    parse() {
      data.split('\n').forEach((line, index) => {
        const mass = parseInt(line, 10);
        this.modules.push({
          index: index + 1,
          mass,
          fuel: this.calculateFuel(mass),
        });
      });
    },
  },

};
</script>
