<template lang="pug">
  .book-details-item
    .item-name
      | {{ name }}
    .item-value
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
