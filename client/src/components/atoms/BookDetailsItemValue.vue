<template lang="pug">
  .book-details-item-value(:style="style")
    | {{ value }}
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component
export default class BookDetailsItemValue extends Vue {
  @Prop({ type: String, required: true })
  private value!: string

  @Prop({ type: String, required: true })
  private color!: string

  get style() {
    return {
      color: this.color
    }
  }
}
</script>

<style lang="sass" scoped>
.book-details-item-value
  color: $text-black
  font-weight: bold
</style>
