<template lang="pug">
  .scan-card
    .cover
      book-cover(:url="book.coverImageUrl")
    div
      .isbn
        | {{ book.isbn }}
      .title
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

<style lang="sass" scoped>
.scan-card
  display: flex
  align-items: center
  height: 160px
  background: white
  border-radius: 4px
  padding: 8px
.cover
  flex-shrink: 0
</style>
