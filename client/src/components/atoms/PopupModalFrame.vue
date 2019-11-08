<template lang="pug">
  .popup-modal-frame-wrapper(v-if="showModal")
    .popup-modal-header
      .return-button(@click="handleClickReturn")
        icon(name="left-arrow" color="var(--text-gray)")
      .title(v-if="title")
        | {{ title }}
    .popup-modal-body
      slot
</template>

<script lang="ts">
import { ExStore } from 'vuex'
import { Vue, Component, Prop } from 'vue-property-decorator'

import Icon from '@/components/assets/Icon.vue'

@Component({
  components: { Icon }
})
export default class PopupModalFrame extends Vue {
  $store!: ExStore

  @Prop({ type: String, required: true })
  private name!: string

  @Prop({ type: String, required: false, default: '' })
  private title!: string

  get showModal() {
    return this.$store.getters['modal/currentModalName'] === this.name
  }

  handleClickReturn() {
    this.$store.commit('modal/pop')
  }
}
</script>

<style lang="sass" scoped>
.popup-modal-frame-wrapper
  display: flex
  flex-flow: column
  width: 360px
  height: 70vh
  background: var(--bg-white)
  border:
    radius: 16px
  pointer-events: auto
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25)

.popup-modal-header
  position: relative
  display: inline-flex
  width: 100%
  padding:
    top: 24px
    left: 32px
    right: 32px
    bottom: 8px
  justify-content: center
  align-items: center

  .title
    color: var(--text-black)
    font:
      weight: bold
      size: 18px

  .return-button
    display: flex
    justify-content: center
    position: absolute
    left: 32px
    cursor: pointer

.popup-modal-body
  padding:
    top: 16px
    left: 16px
    right: 16px
  overflow:
    y: scroll
  flex: 1
</style>
