<template lang="pug">
  .book-details-item-name
    | {{ name }}
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component
export default class BookDetailsItemName extends Vue {
  @Prop({ type: String, required: true })
  private name!: string
}
</script>

<style lang="sass" scoped>
.book-details-item-name
  color: var(--text-gray)
  font-weight: bold
</style>
