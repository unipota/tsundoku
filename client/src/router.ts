import Vue from 'vue'
import Router from 'vue-router'
import Activity from './views/Activity.vue'
import AddBooks from './views/AddBooks.vue'
import BookDetails from './views/BookDetails.vue'
import Kidoku from './views/Kidoku.vue'
import Login from './views/Login.vue'
import NotFound from './views/NotFound.vue'
import Register from './views/Register.vue'
import Toukei from './views/Toukei.vue'
import Tsundoku from './views/Tsundoku.vue'
import UserDetails from './views/UserDetails.vue'

Vue.use(Router)

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'tsundoku',
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
      name: 'kidoku',
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
      name: 'toukei',
      component: Toukei
    },
    {
      path: '/user',
      name: 'userDetails',
      component: UserDetails
    },
    {
      path: '/activity',
      name: 'activity',
      component: Activity
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/register',
      name: 'register',
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
