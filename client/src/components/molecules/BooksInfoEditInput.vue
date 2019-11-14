<template lang="pug">
  .books-info-edit-input
    .label
      | {{ label }}
    text-area-input(
      v-if="type === 'textarea'"
      :value="value"
      @input="handleInput"
    )
    text-input(
      v-else
      with-clear-button
      :value="value"
      :type="type"
      @input="handleInput"
    )
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import TextInput from '@/components/atoms/TextInput.vue'
import TextAreaInput from '@/components/atoms/TextAreaInput.vue'

@Component({
  components: {
    TextInput,
    TextAreaInput
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
  margin:
    bottom: 0.25rem
</style>
