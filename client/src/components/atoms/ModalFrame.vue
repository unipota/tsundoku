<template lang="pug">
  .modal-frame
    slot
    router-link(:to="{path: '../'}" append)
      icon(name="close" color="var(--text-gray)")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

import Icon from '@/components/assets/Icon'

@Component({
  components: { Icon }
})
export default class ModalFrame extends Vue {}
</script>

<style lang="sass">
.modal-frame
  position: fixed
  bottom: 0
  width: 100%
  height: calc(100% - 12px)
  background: white
  border:
    radius: 18px 18px 0 0
</style>
