<template lang="pug">
  .add-books-input
    .label
      | {{ label }}
    text-input(
      with-clear-button
      :value="value"
      :type="type"
      @input="handleInput"
    )
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import TextInput from '@/components/atoms/TextInput.vue'

@Component({
  components: {
    TextInput
  }
})
export default class AddBooksSearch extends Vue {
  @Prop(String) label!: string
  @Prop(String) value!: string
  @Prop({ type: String, default: 'text' }) type!: string

  handleInput(value: string) {
    this.$emit('input', value)
  }
}
</script>

<style lang="sass" scoped>
.label
  font-weight: bold
  color: $text-black
  margin-bottom: 0.25rem
</style>
