<template lang="pug">
  .book-details-item
    .item-name
      | {{ name }}
    .item-value(:style="valueStyle")
      | {{ value }}
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component
export default class BookDetailsItem extends Vue {
  @Prop({ type: String, required: true })
  private name!: string

  @Prop({ type: String, required: true })
  private value!: string

  @Prop({ type: String, default: 'var(--text-black)' })
  private valueColor!: string

  get valueStyle() {
    return {
      color: this.valueColor
    }
  }
}
</script>

<style lang="sass" scoped>
.book-details-item
  display: flex
  align-items: center
  justify-content: space-between

.item-name
  color: $text-gray
  font-weight: bold

.item-value
  color: $text-black
  font-weight: bold
</style>
