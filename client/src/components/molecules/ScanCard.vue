<template lang="pug">
  .scan-card
    book-cover(:book="book")
    span
      | {{ book.isbn }}
    span
      | {{ book.title }}
</template>

<script lang="ts">
import BookCover from '@/components/atoms/BookCover.vue'
import { Vue, Component, Prop } from 'vue-property-decorator'
import { BookRecord } from '../../types/Book'

@Component({
  components: {
    BookCover
  }
})
export default class ScanCard extends Vue {
  @Prop({ type: Object, required: true })
  private book!: BookRecord
}
</script>

<style lang="sass">
</style>
