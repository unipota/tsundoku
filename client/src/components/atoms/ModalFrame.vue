<template lang="pug">
  .modal-frame-overlay(
      :class="`${$store.getters.viewTypeClass} ${$store.getters.modalTransitionClass}`"
      v-click-outside="handleClickOutside"
    )
    .modal-frame-wrapper(
      :class="modalClass"
      :style="{ transform: `translateY(${modalDeltaY}px)` }"
      @touchstart.capture="handleTouchStart"
      @touchmove.capture="handleTouchMove"
      @touchend="handleTouchEnd"
      @touchcancel="handleTouchEnd"
    )
      .modal-frame-top
        .title(v-if="title")
          | {{ title }}
        router-link.close-link(:to="{ path }" append)
          icon.close-icon(name="close" :color="closeColor" :width="16")
      .modal-frame-body(
        ref="modalFrameBody"
        :class="modalBodyClass"
        @scroll="handleModalBodyScroll"
      )
        slot
</template>

<script lang="ts">
import { Vue, Component, Prop, Watch } from 'vue-property-decorator'

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

  @Prop({ type: Boolean, required: false, default: false })
  private isContentNonScrollable!: boolean

  @Prop({ type: Boolean, required: false, default: false })
  private overrideModalInteractivity!: boolean

  @Prop({ type: Boolean, required: false, default: false })
  private overridedIsModalInteractive!: boolean

  /** モーダルが下に降りてる分 */
  private modalDeltaY = 0

  /** モーダルを下におろせるかどうか */
  private isModalInteractive = false
  private hasIsHalfModalInteractiveChanged = false

  // タッチ関連
  private flickSpeed = 0
  private previousTouchClientY = 0
  private scrollTimeoutId = 0

  private modalAnimationState: 'none' | 'css' | 'animationFrame' = 'none'

  readonly flickSpeedThreshold = 20
  readonly modalHideThreshold = 200
  readonly animationDurationMs = 300

  mounted() {
    if (this.overrideModalInteractivity) {
      this.isModalInteractive = true
      return
    }
    this.isModalInteractive = this.isContentNonScrollable
  }

  @Watch('isContentNonScrollable')
  onIsContentNonSctollableChanged(val: boolean) {
    if (this.hasIsHalfModalInteractiveChanged) {
      // すでにスクロールが成功していたら無視
      return
    }
    this.isModalInteractive = val
  }

  get modalClass() {
    return {
      [this.$store.getters.viewTypeClass]: true,
      'no-padding': this.noPadding,
      'css-animating': this.modalAnimationState === 'css'
    }
  }
  get modalBodyClass() {
    return {
      'non-smooth': this.isModalInteractive
    }
  }

  handleModalBodyScroll() {
    clearTimeout(this.scrollTimeoutId)
    this.isModalInteractive = false
    this.scrollTimeoutId = window.setTimeout(() => {
      const body = this.$refs.modalFrameBody as HTMLDivElement
      this.isModalInteractive = body.scrollTop <= 0
      if (!this.hasIsHalfModalInteractiveChanged) {
        this.hasIsHalfModalInteractiveChanged = true
      }
    }, 100)
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

  get isModalAcceptingTouch() {
    const interactive = this.overrideModalInteractivity
      ? this.overridedIsModalInteractive
      : this.isModalInteractive
    return (
      this.$store.state.viewType === 'mobile' &&
      interactive &&
      this.modalAnimationState === 'none'
    )
  }
  handleTouchStart(event: TouchEvent) {
    if (!this.isModalAcceptingTouch) {
      return
    }
    this.flickSpeed = 0
    this.previousTouchClientY = event.touches[0].clientY
  }
  handleTouchMove(event: TouchEvent) {
    if (!this.isModalAcceptingTouch) {
      return
    }
    const y = event.touches[0].clientY
    const flickSpeed = y - this.previousTouchClientY
    if (!this.isContentNonScrollable && flickSpeed > 0) {
      event.preventDefault()
      event.stopPropagation()
    }
    this.flickSpeed = flickSpeed
    this.previousTouchClientY = y
    requestAnimationFrame(() => {
      if (this.modalDeltaY >= -flickSpeed) {
        this.modalDeltaY += flickSpeed
      }
    })
  }
  handleTouchEnd() {
    if (!this.isModalAcceptingTouch) {
      return
    }
    if (this.flickSpeed > this.flickSpeedThreshold) {
      this.handleClickOutside()
      this.modalAnimationState = 'animationFrame'
      return
    }
    if (this.modalDeltaY > this.modalHideThreshold) {
      this.handleClickOutside()
      return
    }
    this.modalAnimationState = 'css'
    this.flickSpeed = 0
    this.previousTouchClientY = 0
    this.$nextTick(() => {
      this.modalDeltaY = 0
    })
    setTimeout(() => {
      this.modalAnimationState = 'none'
    }, this.animationDurationMs)
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
  &.css-animating
    transition: transform 0.3s ease

.modal-frame-body
  padding:
    left: 24px
    right: 24px
  height: calc(100% - 30px - 24px)  // .modal-frame-close の高さとborder-radiusの分を引いた
  overflow:
    x: hidden
    y: auto

  .is-desktop &
    padding-top: 48px
  .is-mobile &
    padding-top: 24px

  .no-padding &
    height: calc(100% + 32px)  // .modal-frame-wrapperのpadding分を戻す
    top: 0
    left: 0
    padding: 0
    margin:
      top: -32px - 34px  // .modal-frame-wrapperと.modal-frame-wrapのpadding分マイナスマージンで上に

  &.non-smooth
    -webkit-overflow-scrolling: auto

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
