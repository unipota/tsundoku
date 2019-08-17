<template lang="pug">
  .book-details-item(:class="{ textarea, clickable }")
    book-details-item-name(:name="name")
    book-details-item-textarea.value-textarea(
      v-if="textarea"
      :value="valueWithPlaceholder"
      :color="color"
      @click="clickable ? $emit('click') : 0"
      )
    book-details-item-value(v-else :value="valueWithPlaceholder" :color="color")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import BookDetailsItemName from '@/components/atoms/BookDetailsItemName.vue'
import BookDetailsItemValue from '@/components/atoms/BookDetailsItemValue.vue'
import BookDetailsItemTextarea from '@/components/atoms/BookDetailsItemTextarea.vue'

@Component({
  components: {
    BookDetailsItemName,
    BookDetailsItemValue,
    BookDetailsItemTextarea
  }
})
export default class BookDetailsItem extends Vue {
  @Prop({ type: String, required: true })
  private name!: string

  @Prop({ type: String, required: true })
  private value!: string

  @Prop({ type: String, default: '' })
  private placeholder!: string

  @Prop({ type: String, default: 'var(--text-black)' })
  private valueColor!: string

  @Prop({ type: String, default: 'var(--text-gray)' })
  private placeholderColor!: string

  @Prop({ type: Boolean, default: false })
  private textarea!: boolean

  @Prop({ type: Boolean, default: false })
  private clickable!: boolean

  get color() {
    return this.value !== '' ? this.valueColor : this.placeholderColor
  }

  get valueWithPlaceholder() {
    return this.value !== '' ? this.value : this.placeholder
  }
}
</script>

<style lang="sass" scoped>
.book-details-item
  display: flex
  align-items: center
  justify-content: space-between

.textarea
  flex-flow: column
  align-items: flex-start

.clickable .value-textarea
  cursor: pointer
</style>
