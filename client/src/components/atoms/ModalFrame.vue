<template lang="pug">
  .modal-frame-overlay
    .modal-frame-body
      .modal-frame-close
        router-link.close-link(:to="{ path }" append)
          icon(name="close" :color="closeColor" :width="16")
      slot
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

import Icon from '@/components/assets/Icon.vue'

@Component({
  components: { Icon }
})
export default class ModalFrame extends Vue {
  @Prop({ type: String, default: '../' })
  private path!: string

  @Prop({ type: String, default: 'var(--text-gray)' })
  private closeColor!: string
}
</script>

<style lang="sass">
.modal-frame-overlay
  position: fixed
  bottom: 0
  width: 100%
  height: 100%
  z-index: 10

.modal-frame-body
  position: absolute
  z-index: 1000
  bottom: 0
  width: 100%
  height: calc(100% - 24px)
  overflow:
    x: hidden
    y: scroll
  background: white
  border:
    radius: 24px 24px 0 0
  padding:
    top: 32px
    left: 24px
    right: 24px
  box-shadow: 0px -4px 12px rgba(0, 0, 0, 0.25)

.modal-frame-close
  position: relative
  width: 100%
  display: flex
  justify-content: flex-end
  z-index: 10
</style>
