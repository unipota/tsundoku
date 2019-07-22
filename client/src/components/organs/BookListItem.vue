<template lang="pug">
  .book-list-item
    router-link.book-list-item__cover(:to="`${$route.matched[0].path}/book/${book.id}`")
      book-cover(:book="book")
    .book-list-item__info(:class="`is-${$store.state.viewType}`")
      .book-list-item__detail
        book-major-info(:book="book")
      .book-list-item__price(v-if="kidoku")
        span.book-list-item__total-price
          | {{ book.price }}
      .book-list-item__progress(v-else)
        book-list-item-progress(:book="book")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { BookRecord } from '../../types/Book'
import BookCover from '@/components/atoms/BookCover.vue'
import BookMajorInfo from '@/components/atoms/BookMajorInfo.vue'
import BookListItemProgress from '@/components/molecules/BookListItemProgress.vue'

@Component({
  components: {
    BookCover,
    BookMajorInfo,
    BookListItemProgress
  }
})
export default class BookListItem extends Vue {
  @Prop({ type: Object, required: true })
  private book!: BookRecord

  @Prop({ type: Boolean, default: false })
  private kidoku!: boolean
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

  width: calc(100% - 80px)
  height: 100%

  padding: 16px 16px 16px 64px
  margin: 8px
  border-radius: 8px
  background-color: $bg-suppressed-gray

  &.is-desktop
    display: flex
    justify-content: space-between

  &.is-mobile
    display: flex
    flex-direction: column

.book-list-item__detail
  width: 100%
  margin-right: 4px
  flex:
    basis: 100%
    grow: 1
    shrink: 1

.book-list-item__progress
  .is-desktop &
    width: 280px
    flex:
      basis: 280px
      grow: 0
      shrink: 1
    align-self: flex-end
  .is-mobile &
    width: 100%

.book-list-item__price
  width: 140px
  text-align: right
  flex:
    basis: 140px
    grow: 0
    shrink: 1
  align-self: flex-end

.book-list-item__total-price
  color: $kidoku-blue
  font:
    weight: bold
    size: 1.1rem
  &::before
    content: '¥'
    margin: 0 0.25rem
</style>
