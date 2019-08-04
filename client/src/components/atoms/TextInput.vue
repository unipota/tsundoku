<template lang="pug">
.text-input
  input(
    type="text"
    :value="value"
    :placeholder="placeholder"
    @input="$emit('input', $event.target.value)"
    @keyup.enter="$emit('keyup-enter')"
  )
  .close(@click="$emit('input', '')")
    transition(name="transition-close")
      icon(
        v-if="withClearButton && value"
        name="close"
        color="var(--text-gray)"
        :height="16"
      )
</template>

<script lang="ts">
import { Vue, Prop, Component } from 'vue-property-decorator'
import Icon from '@/components/assets/Icon.vue'

@Component({
  components: {
    Icon
  }
})
export default class TextInput extends Vue {
  @Prop({ type: String, default: '' }) value!: string
  @Prop({ type: String, default: '' }) placeholder!: string
  @Prop({ type: Boolean, default: false }) withClearButton!: boolean
}
</script>

<style lang="sass" scoped>
.text-input
  display: flex
  align-items: center
  justify-content: space-between
  width: 100%
  height: 40px
  border-radius: 8px
  background: var(--bg-gray)
  padding: 0.5rem 1rem

.text-input input
  width: 100%
  height: 100%
  flex: 100% 1 1
  font-size: 1rem
  color: var(--text-black)
  &::placeholder
    color: var(--text-gray)

.close
  display: flex
  align-items: center
  flex: 16px 0 0
  cursor: pointer
  opacity: 0.75
  &:hover
    opacity: 1

.transition-close
  &-enter, &-leave-to
    opacity: 0

  &-enter-active, &-leave-active
    transition: opacity .3s
</style>
