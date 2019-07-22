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
      .book-list-item__progress
        book-list-item-progress(:book="book")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { BookRecord } from '../../types/Book'
import BookCover from '@/components/atoms/BookCover.vue'
import BookListItemProgress from '@/components/molecules/BookListItemProgress.vue'

@Component({
  components: {
    BookCover,
    BookListItemProgress
  }
})
export default class BookListItem extends Vue {
  @Prop({ type: Object, required: true })
  private book!: BookRecord
}
</script>

<style lang="sass">
.book-list-item
  position: relative
  align-items: center
  width: 100%
  height: 160px

.book-list-item__cover
  position: absolute
  top: 0
  left: 0
  display: flex
  align-items: center
  height: 100%
  margin-right: 16px

.book-list-item__info
  position: absolute
  top: 0
  left: 64px
  z-index: -1

  display: flex
  justify-content: space-between

  width: calc(100% - 64px)
  height: 100%

  padding: 16px 16px 16px 64px
  margin: 8px
  border-radius: 8px
  background-color: $bg-suppressed-gray

.book-list-item__detail
  width: 100%
  margin-right: 4px
  flex:
    basis: 100%
    grow: 1
    shrink: 1

.book-list-item__title
  font:
    size: 1.2rem
    weight: bold
  width: 100%
  margin-bottom: 8px
  overflow: hidden

  display: -webkit-box
  -webkit-line-clamp: 2
  -webkit-box-orient: vertical

  max-height: 3.6rem

.book-list-item__author
  height: 1.5rem
  width: 100%
  overflow: hidden
  white-space: nowrap
  text-overflow: ellipsis

.book-list-item__progress
  width: 280px
  flex:
    basis: 280px
    grow: 0
    shrink: 1
  align-self: flex-end
</style>
