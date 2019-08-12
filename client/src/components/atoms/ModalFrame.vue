<template lang="pug">
  .modal-frame-overlay(
    :class="`${$store.getters.viewTypeClass} ${$store.getters.modalTransitionClass}`"
    v-click-outside="handleClickOutside"
    )
    .modal-frame-wrapper(:class="modalClass")
      .modal-frame-top
        .title(v-if="title")
          | {{ title }}
        router-link.close-link(:to="{ path }" append)
          icon.close-icon(name="close" :color="closeColor" :width="16")
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

  @Prop({ type: Boolean, default: false })
  private noPadding!: boolean

  @Prop({ type: String, required: false, default: '' })
  private title!: string

  get modalClass() {
    return {
      [this.$store.getters.viewTypeClass]: true,
      'no-padding': this.noPadding
    }
  }

  handleClickOutside() {
    let path = this.path
    let backCount = 0
    while (path.startsWith('../')) {
      path = path.substring(3)
      backCount += 1
    }
    const splitted = this.$route.path.split('/')
    const pathLength = splitted.length
    const base = splitted.slice(0, pathLength - Math.min(backCount, pathLength))
    base.push(path)

    this.$nextTick(() => {
      this.$router.push(base.join('/'))
    })
  }
}
</script>

<style lang="sass" scoped>
.modal-frame-overlay
  position: absolute
  bottom: 0
  height: 100%
  z-index: 10
  pointer-events: auto
  &.is-mobile
    width: 100%
  &.is-desktop
    width: 600px

.modal-frame-wrapper
  position: absolute
  z-index: 1000
  bottom: 0
  width: 100%
  background: white
  box-shadow: 0px -4px 12px rgba(0, 0, 0, 0.25)
  padding:
    top: 32px
  overflow: hidden
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
  height: calc(100% - 30px - 24px) // .modal-frame-close の高さとborder-radiusの分を引いた
  overflow:
    x: hidden
    y: auto

  .is-desktop &
    padding-top: 48px
  .is-mobile &
    padding-top: 24px

  .no-padding &
    padding: 0

.modal-frame-top
  position: relative
  width: 100%
  display: flex
  justify-content: flex-end
  z-index: 10
  height: 34px
  .title
    margin: auto
    font-size: 22px
    font-weight: bold
  .close-link
    cursor: pointer
    position: absolute
    right: 18px
    top: 50%
    transform: translateY(-50%)
    width: 48px
    height: 48px
    padding: 12px
    display: flex
    border:
      radius: 100%
    transition: background .5s
    &:hover
      background: rgba(100,100,100,0.2)
    .close-icon
      margin: auto
</style>
