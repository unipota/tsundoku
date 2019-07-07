import Vue from 'vue'
import Router from 'vue-router'
import Activity from './views/Activity.vue'
import BookDetails from './views/BookDetails.vue'
import Home from './views/Home.vue'
import Kidoku from './views/Kidoku.vue'
import Login from './views/Login.vue'
import NotFound from './views/NotFound.vue'
import Register from './views/Register.vue'
import Tsundoku from './views/Tsundoku.vue'
import UserDetails from './views/UserDetails.vue'

Vue.use(Router)

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
    },
    {
      path: '/tsundoku',
      name: 'tsundoku',
      component: Tsundoku,
    },
    {
      path: '/kidoku',
      name: 'kidoku',
      component: Kidoku,
    },
    {
      path: '/book/:id',
      name: 'bookDetails',
      component: BookDetails,
    },
    {
      path: '/user',
      name: 'userDetails',
      component: UserDetails,
    },
    {
      path: '/activity',
      name: 'activity',
      component: Activity,
    },
    {
      path: '/login',
      name: 'login',
      component: Login,
    },
    {
      path: '/register',
      name: 'register',
      component: Register,
    },
    {
      path: '*',
      name: 'notfound',
      component: NotFound,
    },
  ],
})

export default router
