<template lang="pug">
.text-input(@click="focusInput")
  .icon-wrap
    slot
  input(
    type="text"
    :value="value"
    :type="type"
    :placeholder="placeholder"
    @input="$emit('input', $event.target.value)"
    @keyup.enter="$emit('keyup-enter')"
    ref="input"
    v-focus="focus"
  )
  transition(name="transition-clear")
    .clear(v-if="withClearButton && value" @click="handleClear")
      icon(
        name="close"
        color="var(--text-gray)"
        :height="12"
        :width="12"
      )
</template>

<script lang="ts">
import { Vue, Prop, Component } from 'vue-property-decorator'
import Icon from '@/components/assets/Icon.vue'
import { DirectiveOptions } from 'vue'

const focusDirective: DirectiveOptions = {
  inserted(el, binding) {
    if (binding.value) {
      window.setTimeout(() => {
        el.focus()
      }, 500)
    }
  }
}

@Component({
  directives: {
    focus: focusDirective
  },
  components: {
    Icon
  }
})
export default class TextInput extends Vue {
  @Prop({ type: String, default: '' }) value!: string
  @Prop({ type: String, default: '' }) placeholder!: string
  @Prop({ type: Boolean, default: false }) withClearButton!: boolean
  @Prop({ type: String, default: 'text' }) type!: string
  @Prop({ type: Boolean, default: false }) focus!: boolean

  handleClear() {
    this.$emit('input', '')
    this.focusInput()
  }

  focusInput() {
    const el = this.$refs.input as HTMLInputElement
    el.focus()
  }
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
  cursor: text

.text-input input
  width: 100%
  height: 100%
  flex: 100% 1 1
  font:
    weight: bold
    size: 1rem
  color: var(--text-black)
  &::placeholder
    color: var(--text-gray)
    font:
      weight: bold

.icon-wrap
  display: flex
  align-items: center
  justify-content: center
  transform: translateX(-6px)

.clear
  display: flex
  flex-shrink: 0
  justify-content: center
  align-items: center
  width: 26px
  height: 26px
  cursor: pointer
  opacity: 0.75
  border:
    radius: 100%
  transition: opacity .5s, background .5s

  &:hover
    opacity: 1
    background: rgba(100,100,100,0.1)

.transition-clear
  &-enter, &-leave-to
    opacity: 0

  &-enter-active, &-leave-active
    transition: opacity .3s
</style>
