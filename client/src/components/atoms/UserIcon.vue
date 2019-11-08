<template lang="pug">
  .user-icon(:style="iconStyle")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component
export default class UserIcon extends Vue {
  @Prop({ type: String, default: '' }) readonly src!: string

  get iconStyle() {
    return {
      backgroundImage: `url('${this.src}')`
    }
  }
}
</script>

<style lang="sass" scoped>
.user-icon
  width: 100%
  height: 100%
  border:
    radius: 100%
  background: var(--text-gray)
  overflow: hidden
  background-size: cover
</style>
