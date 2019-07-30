<template lang="pug">
  div.add-book-card(:class="`${$store.getters.viewTypeClass} ${type}`")
    book-cover(:book="book")
    div.info(:class="`${$store.getters.viewTypeClass}`")
      book-major-info(:book="book" :size="bookMajorInfoSize")
      div.price-buttons-wrapper(:class="`${$store.getters.viewTypeClass}`")
        div.price(:class="`${$store.getters.viewTypeClass}`")
          | {{ $t('currency') }}{{ book.price.toLocaleString() }}
        div.buttons
          component(:is="editButtonComponent" to="/add-books-edit")
          add-tsundoku-button(size="small" @add-tsundoku="addTsundoku")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import { BookRecord } from '../../types/Book'
import api from '../../store/general/api'
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
  private book!: BookRecord

  get editButtonComponent() {
    return (
      'edit-button' + (this.$store.state.viewType === 'mobile' ? '' : '-large')
    )
  }

  get bookMajorInfoSize() {
    return this.$store.state.viewType === 'mobile' ? 'small' : 'normal'
  }

  addTsundoku() {
    console.log(this.book)
    api
      .addNewBook(this.book)
      .then(res => {
        console.log(res.data)
        // TODO: 追加した本のカードは非表示にしたい
      })
      .catch(err => {
        window.alert(this.$t('networkError') + '\n' + err.response)
      })
  }
}
</script>

<style lang="sass" scoped>
.add-book-card
  display: flex
  padding: 10px
  &.scan
    box-shadow: 2px 4px 4px rgba(0, 0, 0, 0.3)
    border-radius: 8px
    &.is-mobile
      width: 300px
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
  min-width: 6rem;
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
</style>