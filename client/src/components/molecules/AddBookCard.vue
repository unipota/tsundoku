<template lang="pug">
  transition(name="transition-card")
    div.add-book-card(
      :class="`${$store.getters.viewTypeClass} ${type}`"
      v-if="showCard"
    )

      book-cover(:url="book.coverImageUrl")

      div.info(:class="`${$store.getters.viewTypeClass}`")
        book-major-info(:title="book.title" :authors="book.author" :size="bookMajorInfoSize")
        div.price-buttons-wrapper(:class="`${$store.getters.viewTypeClass}`")
          div.price(:class="`${$store.getters.viewTypeClass}`")
            | {{ $t('currency') + book.price.toLocaleString() }}
          div.buttons
            component(
              :is="editButtonComponent"
              :to="firstRouteName + '/add-books-edit'"
              :book="book"
            )
            add-tsundoku-button(
              size="small"
              :bookAdded="bookAdded"
              @add-tsundoku="addTsundoku"
            )
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import { BookRecord } from '../../types/Book'
import AddTsundokuButton from '../atoms/AddTsundokuButton.vue'
import BookCover from '../atoms/BookCover.vue'
import BookMajorInfo from '../atoms/BookMajorInfo.vue'
import EditButton from '../atoms/EditButton.vue'
import EditButtonLarge from '../atoms/EditButtonLarge.vue'
import PriceDisplay from '../atoms/PriceDisplay.vue'

type cardType = 'search' | 'scan'

@Component({
  components: {
    AddTsundokuButton,
    BookCover,
    BookMajorInfo,
    EditButton,
    EditButtonLarge,
    PriceDisplay
  }
})
export default class AddBookCard extends Vue {
  public $store!: ExStore

  @Prop({ type: String, required: true })
  private readonly type!: cardType

  @Prop({ type: Object, required: true })
  private readonly book!: BookRecord

  private bookAdded = false
  private showCard = true

  async mounted() {
    this.showCard = this.type == 'search'
    if (this.type == 'scan') {
      this.$emit('appear-start')

      await this.$nextTick()
      this.showCard = true

      await new Promise(r => window.setTimeout(r, 1000))
      this.$emit('appear')
    }
  }

  get editButtonComponent() {
    return (
      'edit-button' + (this.$store.state.viewType === 'mobile' ? '' : '-large')
    )
  }

  get bookMajorInfoSize() {
    return this.$store.state.viewType === 'mobile' ? 'small' : 'normal'
  }

  get firstRouteName() {
    return this.$route.matched[0].path
  }

  async addTsundoku() {
    console.log(this.book)
    try {
      const res = await this.$store.dispatch('addNewBook', {
        book: this.book
      })
      console.log(res.data)

      this.bookAdded = true // → AddTsundokuButtonがチェックに変わる

      await new Promise(r => window.setTimeout(r, 800))
      this.showCard = false // → カードが消える
      this.$emit('to-remove')

      await new Promise(r => window.setTimeout(r, 1000))
      this.$emit('remove')

      this.$store.dispatch('getMyBooks')
    } catch (err) {
      window.alert(this.$t('networkError') + '\n' + err.response)
    }
  }
}
</script>

<style lang="sass" scoped>
.add-book-card
  display: flex
  padding: 10px
  &.scan
    filter: drop-shadow2px 4px 4px rgba(0, 0, 0, 0.3)
    border-radius: 8px
    background: white
  &.is-mobile
    width: 100%
    max-width: 300px
  &.scan.is-desktop
    width: 500px
  &.search.is-desktop
    width: 100%

.book-cover
  min-width: 96px
  min-height: 149px

.info
  width: 100%
  padding:
    top: 5px
    right: 0
    bottom: 5px
    left: 20px
  display: grid

.price-buttons-wrapper
  margin: auto 0 0 0
  &.is-desktop
    display: flex

  .price
    margin:
      right: 21px
    font-weight: bold
    font-size: 20px
    color: var(--tsundoku-red)
    min-width: 6rem
    &.is-mobile
      text-align: right
    &.is-desktop
      text-align: left
  .buttons
    display: flex
    width: min-content
    margin: 0 14px 0 auto
    .add-tsundoku-button-wrapper
      margin:
        left: 6px

.transition-card
  &-enter.scan
    transform: translateY(100%) // scanCardは下から生える
    opacity: 0

  &-enter-to.scan
    opacity: 1

  &-enter-active
    transition: transform 1s $easeInOutQuint, opacity .5s

  &-leave-to
    opacity: 0
    &.search
      transform: translateX(100%) // searchCardは右に消える
    &.scan
      transform: translateY(100%) // scanCardは下に消える

  &-leave-active
    transition: transform 1s $easeInOutQuint, opacity .5s
</style>
