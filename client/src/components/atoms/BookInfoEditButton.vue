<template lang="pug">
  div.add-tsundoku-button-wrapper(
    :class="`is-${size}`"
    :style="buttonStyle"
    @click="$emit('add-tsundoku')"
  )
    transition(name="transition-button-content" mode="out-in")
      div.add-tsundoku-button.not-added(v-if="!completed" key="not-added")
        icon.icon-plus(
          v-if="type === 'add'"
          name="plus"
          :width="iconPlusSize"
          :height="iconPlusSize"
          :color="color"
        )
        icon.icon-plus(
          v-else-if="type === 'edit'"
          name="check"
          :width="iconCheckSize"
          :height="iconCheckSize"
          :color="color"
        )
        span.text
          | {{ text }}
      div.add-tsundoku-button.added(v-else key="added")
        icon.icon-check(
          name="check"
          :width="iconCheckSize"
          :height="iconCheckSize"
          :color="color"
        )
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

import Icon from '@/components/assets/Icon.vue'

type ButtonSize = 'large' | 'small'
type ButtonType = 'add' | 'edit'

@Component({
  components: { Icon }
})
export default class BookInfoEditButton extends Vue {
  @Prop({ type: String, default: 'large' })
  private readonly size!: ButtonSize

  @Prop({ type: Boolean, default: false })
  private readonly completed!: boolean

  @Prop({ type: String, default: 'add' })
  private readonly type!: ButtonType

  get iconPlusSize() {
    switch (this.size) {
      case 'large':
        return 14
        break
      case 'small':
        return 8.2
        break
    }
  }

  get iconCheckSize() {
    switch (this.size) {
      case 'large':
        return 20
        break
      case 'small':
        return 14
        break
    }
  }

  get color() {
    return {
      add: 'var(--tsundoku-red)',
      edit: 'var(--kidoku-blue)'
    }[this.type]
  }

  get text() {
    return {
      add: this.$t('addTsundoku'),
      edit: '編集完了'
    }[this.type]
  }

  get buttonStyle() {
    return {
      color: this.color,
      borderColor: this.color
    }
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
    .icon-plus, .icon-check
      margin: auto 4px auto 0
    .text
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

.transition-button-content
  &-enter, &-leave-to
    opacity: 0
  &-leave-active
    transition: opacity .1s
  &-enter-active
    transition: opacity .3s
</style>
