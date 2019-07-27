<template lang="pug">
  .modal-frame-overlay
    .modal-frame-body
      slot
      router-link.close-button(:to="{path: '../'}" append)
        icon(name="close" color="var(--text-gray)")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

import Icon from '@/components/assets/Icon.vue'

@Component({
  components: { Icon }
})
export default class ModalFrame extends Vue {}
</script>

<style lang="sass">
.modal-frame-overlay
  position: fixed
  bottom: 0
  width: 100%
  height: 100%

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

.close-button
  position: absolute
  right: 16px
  top: 16px
  padding: 12px
</style>
