<template lang="pug">
  div.add-tsundoku-button-wrapper(:class="'is-' + size" @click="$emit('add-tsundoku')")
    div.add-tsundoku-button
      icon.icon-plus(name="plus" :width="iconSize" :height="iconSize" :color="iconColor")
      span.text
        | {{ $t('addTsundoku') }}
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

import Icon from '@/components/assets/Icon.vue'

type ButtonSize = 'large' | 'small'

@Component({
  components: { Icon }
})
export default class AddTsundokuButton extends Vue {
  @Prop({ type: String, default: 'large' })
  private readonly size!: ButtonSize

  get iconSize() {
    switch (this.size) {
      case 'large':
        return 14
        break
      case 'small':
        return 8.2
        break
    }
  }

  get iconColor() {
    return 'var(--tsundoku-red)'
  }
}
</script>

<style lang="sass" scoped>
.add-tsundoku-button-wrapper
  cursor: pointer
  border: solid var(--tsundoku-red)
  border-radius: 33px
  display: flex
  padding:
    top: 2px
  .add-tsundoku-button
    margin: 0 auto
    display: flex
    .icon-plus
      margin: auto 4px auto 0
    .text
      color: var(--tsundoku-red)
      font-weight: bold
      margin: auto 0
  &.is-small
    width: 99px
    height: 33px
    .text
      font-size: 13px
  &.is-large
    width: 166px
    height: 57px
    border-width: 4px
    .text
      font-size: 22px
</style>
