<template lang="pug">
  .book-list-item
    .book-list-item__cover
      book-cover(:book="book")
    .book-list-item__info
      .book-list-item__detail
        .book-list-item__title
          | {{ book.title }}
        .book-list-item__author
          | {{ book.author }}
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import BookCover from '@/components/atoms/BookCover.vue'
import { BookRecord } from '../../types/Book'

@Component({
  components: {
    BookCover
  }
})
export default class BookListItem extends Vue {
  @Prop({ type: Object, required: true })
  private book!: BookRecord
}
</script>

<style lang="sass">
.book-list-item
  display: flex
  align-items: center
  width: 600px
.book-list-item__info
  position: relative
  z-index: -1
  width: 100%
  height: 176px
  left: -64px
  padding: 16px 16px 16px 64px
  margin: 8px
  border-radius: 8px
  background-color: $bg-suppressed-gray
.book-list-item__cover
  margin-right: 16px
.book-list-item__title
  font:
    size: 1.2rem
    weight: bold
  margin-bottom: 8px
  height: 2lh
  overflow: hidden
</style>
