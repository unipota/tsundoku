<template lang="pug">
  .dialog-frame
    .dialog-overlay(@click="closeDialog")
    .dialog-body
      slot(name="body")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

@Component({ components: {} })
export default class DialogFrame extends Vue {
  closeDialog() {
    this.$emit('close')
  }
}
</script>

<style lang="sass" scoped>
.dialog-frame
  width: 260px;
  padding: 18px
  background: var(--bg-white)
  border:
    radius: 18px
  box-shadow: 0px 0px 12px rgba(0, 0, 0, 0.25)

.dialog-overlay
  position: absolute
  z-index: -1
  left: 0
  top: 0
  width: 100%
  height: 100%
</style>
