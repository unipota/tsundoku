import 'vuex'
import * as Root from '../../store/general/type'
import * as Modal from '../../store/modal/type'

declare module 'vuex' {
  type RootState = Root.S & {
    modal: Modal.S
  }
  type RootGetters = Root.RG & Modal.RG
  type RootMutations = Root.RM & Modal.RM
  type RootActions = Root.RA & Modal.RA
}
