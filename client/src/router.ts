import Vue from 'vue'
import Router from 'vue-router'
import AddBooksSearch from './views/AddBooksScan.vue'
import AddBooksScan from './views/AddBooksSearch.vue'
import BookDetails from './views/BookDetails.vue'
import Kidoku from './views/Kidoku.vue'
import Login from './views/Login.vue'
import NotFound from './views/NotFound.vue'
import Register from './views/Register.vue'
import Toukei from './views/Toukei.vue'
import Tsundoku from './views/Tsundoku.vue'
import UserDetails from './views/UserDetails.vue'
import { VueConstructor } from 'vue'
import { RouteConfig } from 'vue-router'

const kebabaize = (camel: string): string =>
  camel.replace(/([a-z0-9]|(?=[A-Z]))([A-Z])/g, '$1-$2').toLowerCase()

export type ViewNames =
  | 'tsundoku'
  | 'kidoku'
  | 'toukei'
  | 'addBooksSearch'
  | 'addBooksScan'
  | 'bookDetails'
  | 'login'
  | 'user'
  | 'register'
  | 'notfound'

type ViewNamesToComponentMap = {
  [P in ViewNames]: VueConstructor
}

interface ViewDeclaration {
  view: ViewNames
  path?: string
  children?: ViewDeclaration[]
}

const toRoutesObject = (
  routes: ViewDeclaration[],
  componentMap: ViewNamesToComponentMap,
  basePath = '/',
  baseName = ''
): RouteConfig[] => {
  return routes.map(
    (route): RouteConfig => {
      const path =
        route.path === '/'
          ? '/'
          : (basePath.endsWith('/') ? basePath : basePath + '/') +
            (route.path !== undefined ? route.path : kebabaize(route.view))
      const name =
        route.view +
        (baseName === '' ? '' : baseName[0].toUpperCase() + baseName.slice(1))
      return {
        path,
        name,
        component: componentMap[route.view],
        children: route.children
          ? toRoutesObject(route.children, componentMap, path, name)
          : undefined
      }
    }
  )
}

const viewNamesToComponentMap: ViewNamesToComponentMap = {
  tsundoku: Tsundoku,
  kidoku: Kidoku,
  toukei: Toukei,
  addBooksSearch: AddBooksSearch,
  addBooksScan: AddBooksScan,
  bookDetails: BookDetails,
  login: Login,
  user: UserDetails,
  register: Register,
  notfound: NotFound
}

const modalSubTree: ViewDeclaration[] = [
  { view: 'addBooksSearch' },
  { view: 'addBooksScan' },
  { view: 'bookDetails', path: 'book/:id' }
]

const routes: ViewDeclaration[] = [
  { view: 'tsundoku', path: '/', children: modalSubTree },
  { view: 'kidoku', children: modalSubTree },
  { view: 'toukei', children: modalSubTree },
  { view: 'user' },
  { view: 'register' },
  { view: 'notfound', path: '*' }
]

Vue.use(Router)

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: toRoutesObject(routes, viewNamesToComponentMap)
})

export default router
