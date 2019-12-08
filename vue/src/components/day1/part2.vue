<template>
  <v-card>
    <v-card-text>
      <div>Part Two</div>
      <v-divider></v-divider>
      <v-list class="title font-weight-light">
        <v-list-item>
          <v-list-item-content>Total Recursive Fuel</v-list-item-content>
          <v-list-item-content
            class="align-end font-weight-medium"
          >{{ recursiveFuel }}</v-list-item-content>
        </v-list-item>
      </v-list>
      <v-data-table
        :headers="headers"
        :items="moduleData"
        class="mt-3"
        :items-per-page="5"
      ></v-data-table>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  props: ['modules'],
  data: () => ({
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
        value: 'recursive',
      },
    ],
  }),
  computed: {
    moduleData() {
      return this.modules.map((mass, i) => ({
        index: i + 1,
        mass,
        recursive: this.calculateRecursiveFuel(mass),
      }));
    },
    recursiveFuel() {
      return this.moduleData.reduce((acc, cur) => acc + cur.recursive, 0);
    },
  },
  methods: {
    calculateFuel(mass) {
      return Math.floor(mass / 3) - 2;
    },
    calculateRecursiveFuel(mass) {
      let totalFuel = 0;
      let recursiveFuel = mass;

      while (recursiveFuel > 0) {
        recursiveFuel = this.calculateFuel(recursiveFuel);
        if (recursiveFuel > 0) totalFuel += recursiveFuel;
      }

      return totalFuel;
    },
  },
};
</script>
