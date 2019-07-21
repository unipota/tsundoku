import Vue from 'vue'
import Router from 'vue-router'
import AddBooks from './views/AddBooks.vue'
import BookDetails from './views/BookDetails.vue'
import Kidoku from './views/Kidoku.vue'
import Login from './views/Login.vue'
import NotFound from './views/NotFound.vue'
import Register from './views/Register.vue'
import Toukei from './views/Toukei.vue'
import Tsundoku from './views/Tsundoku.vue'
import UserDetails from './views/UserDetails.vue'

import { RouteNames } from './types/RouteNames'

Vue.use(Router)

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: RouteNames.tsundoku,
      component: Tsundoku,
      children: [
        {
          path: 'add-books',
          name: 'addBooksTsundoku',
          component: AddBooks
        },
        {
          path: 'book/:id',
          name: 'BookDetailsTsundoku',
          component: BookDetails
        }
      ]
    },
    {
      path: '/kidoku',
      name: RouteNames.kidoku,
      component: Kidoku,
      children: [
        {
          path: 'add-books',
          name: 'addBooksKidoku',
          component: AddBooks
        },
        {
          path: 'book/:id',
          name: 'bookDetailsKidoku',
          component: BookDetails
        }
      ]
    },
    {
      path: '/toukei',
      name: RouteNames.toukei,
      component: Toukei
    },
    {
      path: '/user',
      name: RouteNames.user,
      component: UserDetails
    },
    {
      path: '/login',
      name: RouteNames.login,
      component: Login
    },
    {
      path: '/register',
      name: RouteNames.register,
      component: Register
    },
    {
      path: '*',
      name: 'notfound',
      component: NotFound
    }
  ]
})

export default router
