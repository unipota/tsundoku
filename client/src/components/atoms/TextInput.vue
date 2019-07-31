<template lang="pug">
.text-input
  input(
    type="text"
    :value="value"
    :placeholder="placeholder"
    @input="$emit('input', $event.target.value)"
    @keyup.enter="$emit('keyup-enter')"
  )
</template>

<script lang="ts">
import { Vue, Prop, Component } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class TextInput extends Vue {
  @Prop({ type: String, default: '' }) value!: string
  @Prop({ type: String, default: '' }) placeholder!: string
}
</script>

<style lang="sass">
.text-input
  width: 100%
  height: 40px
  border-radius: 8px
  background: var(--bg-gray)
  padding: 0.5rem 1rem

.text-input input
  width: 100%
  height: 100%
  font-size: 1rem
  color: var(--text-black)
  &::placeholder
    color: var(--text-gray)
</style>
