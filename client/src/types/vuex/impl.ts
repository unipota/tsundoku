import 'vuex'
import * as Root from '../../store/general/type'

declare module 'vuex' {
  type RootState = Root.S & {
    general: Root.S
  }
  type RootGetters = Root.RG
  type RootMutations = Root.RM
  type RootActions = Root.RA
}
