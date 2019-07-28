<template lang="pug">
  .modal-frame-overlay(:class="`is-${$store.state.viewType}`")
    .modal-frame-wrapper(:class="`is-${$store.state.viewType}`")
      .modal-frame-close
        router-link.close-link(:to="{ path }" append)
          icon(name="close" :color="closeColor" :width="16")
      .modal-frame-body
        slot
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

import Icon from '@/components/assets/Icon.vue'
import { ExStore } from 'vuex'

@Component({
  components: { Icon }
})
export default class ModalFrame extends Vue {
  public $store!: ExStore

  @Prop({ type: String, default: '../' })
  private path!: string

  @Prop({ type: String, default: 'var(--text-gray)' })
  private closeColor!: string
}
</script>

<style lang="sass" scoped>
.modal-frame-overlay
  position: fixed
  bottom: 0
  height: 100%
  z-index: 10
  &.is-mobile
    width: 100%
  &.is-desktop
    width: 526px

.modal-frame-wrapper
  position: absolute
  z-index: 1000
  bottom: 0
  width: 100%
  background: white
  box-shadow: 0px -4px 12px rgba(0, 0, 0, 0.25)
  padding:
    top: 32px
  &.is-mobile
    height: calc(100% - 24px)
    border:
      radius: 24px 24px 0 0
  &.is-desktop
    height: 100%
    border:
      radius: 0 24px 24px 0

.modal-frame-body
  padding:
    left: 24px
    right: 24px
  height: calc(100% - 30px - 12px)
  overflow:
    x: hidden
    y: auto

.modal-frame-close
  position: relative
  width: 100%
  display: flex
  justify-content: flex-end
  z-index: 10
  padding:
    right: 24px
</style>
