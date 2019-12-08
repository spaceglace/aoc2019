<template>
  <v-app>
    <v-navigation-drawer permanent app>
      <v-list nav>
        <v-list-item
          v-for="day in days"
          :key="day.name"
          link
          :to="day.link"
        >
          <v-list-item-icon>
            <v-icon
              :color="day.color"
            >mdi-star</v-icon>
          </v-list-item-icon>
          <v-list-item-content>{{ day.name }}</v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-app-bar app>
      <img
        :src="require('./assets/glaceon.png')"
        height="80"
      />
      <v-toolbar-title
        class="display-1 font-weight-thin text-upper"
      >AOC2019</v-toolbar-title>
    </v-app-bar>
    <v-content>
      <router-view></router-view>
    </v-content>
  </v-app>
</template>

<script>
const colors = {
  gold: 'amber accent-2',
  silver: 'blue-grey lighten-3',
};

export default {
  data: () => ({
    days: [
      {
        name: 'Day 1',
        link: '/day1',
        color: colors.silver,
      },
    ],
  }),
};
</script>
