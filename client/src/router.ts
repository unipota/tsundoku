import Vue from 'vue'
import Router from 'vue-router'
import AddBooksScan from './views/AddBooksScan.vue'
import AddBooksEdit from './views/AddBooksEdit.vue'
import AddBooksSearch from './views/AddBooksSearch.vue'
import BookDetails from './views/BookDetails.vue'
import Kidoku from './views/Kidoku.vue'
import NotFound from './views/NotFound.vue'
import Toukei from './views/Toukei.vue'
import Tsundoku from './views/Tsundoku.vue'
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
  | 'addBooksEdit'
  | 'bookDetails'
  | 'bookDetailsEdit'
  | 'notfound'

type ViewNamesToComponentMap = {
  [P in ViewNames]: VueConstructor
}

interface ViewDeclaration {
  view: ViewNames
  path?: string
  children?: ViewDeclaration[]
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  props?: Record<string, any>
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  meta?: Record<string, any>
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
          : undefined,
        props: route.props ? route.props : {},
        meta: route.meta ? route.meta : {}
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
  addBooksEdit: AddBooksEdit,
  bookDetails: BookDetails,
  bookDetailsEdit: BookDetails,
  notfound: NotFound
}

const modalSubTree: ViewDeclaration[] = [
  { view: 'addBooksSearch', meta: { isModal: true } },
  { view: 'addBooksScan', meta: { isModal: true } },
  { view: 'addBooksEdit', meta: { isModal: true } },
  {
    view: 'bookDetailsEdit',
    path: 'book/:id/edit',
    props: {
      isEditing: true
    },
    meta: { isModal: true }
  },
  { view: 'bookDetails', path: 'book/:id', meta: { isModal: true } }
]

const routes: ViewDeclaration[] = [
  { view: 'tsundoku', path: '/', children: modalSubTree },
  { view: 'kidoku', children: modalSubTree },
  { view: 'toukei', children: modalSubTree },
  { view: 'notfound', path: '*' }
]

Vue.use(Router)

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: toRoutesObject(routes, viewNamesToComponentMap)
})

router.afterEach((to, from) => {
  // @ts-ignore
  if (typeof gtag !== 'undefined') {
    // @ts-ignore
    gtag('config', 'UA-152571593-1', { page_path: to.path })
  }
})

export default router
