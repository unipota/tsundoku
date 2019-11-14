<template lang="pug">
  .book-details-item-value(:style="style" v-html="sanitizedValue")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import sanitizeHTML from 'sanitize-html'

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

  get sanitizedValue(): string {
    return this.insertBrTag(sanitizeHTML(this.value))
  }

  insertBrTag(value: string): string {
    return value.replace(/\r?\n/g, '<br>')
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
  color: var(--text-gray)
  background:
    color: var(--bg-gray)
  font-weight: normal
</style>
