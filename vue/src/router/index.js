import Vue from 'vue';
import VueRouter from 'vue-router';
import Day1 from '../views/Day1.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/day1',
    name: 'day1',
    component: Day1,
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
