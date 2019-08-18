<template lang="pug">
  .book-details-item-value(:style="style")
    | {{ value }}
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component
export default class BookDetailsItemTextarea extends Vue {
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
  width: 100%
  margin: 8px 0
  padding: 8px
  border:
    radius: 8px
  color: $text-black
  background:
    color: $bg-gray
  font-weight: normal
</style>
