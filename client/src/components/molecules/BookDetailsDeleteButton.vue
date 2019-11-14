<template lang="pug">
  .book-details-delete-button
    book-details-action-button.action(
      :expanded="expanded"
      icon="remove"
      :label="$t('delete')"
      @click="handleClick")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import BookDetailsActionButton from '@/components/atoms/BookDetailsActionButton.vue'

@Component({
  components: {}
})
export default class BookDetailsDeleteButton extends Vue {
  @Prop({ type: Boolean, default: false })
  private expanded!: boolean

  handleClick(e: MouseEvent) {
    this.$emit('click', e)
  }
}
</script>

<style lang="sass" scoped></style>
