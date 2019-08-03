<template lang="pug">
  modal-frame.add-books-edit(
    :title="$t('addBooksSearchTitle')"
    :class="$store.getters.viewTypeClass"
  )
    book-cover(
      :hasShadow="true"
      :url="coverImageUrl"
    )
    add-books-edit-input.edit-input(
      :label="$t('title')"
      v-model="title"
    )
    add-books-edit-input.edit-input(
      :label="$t('author')"
      v-model="author"
    )
    add-books-edit-input.edit-input(
      :label="$t('price')"
      v-model="price"
      type="number"
    )
    add-books-edit-input.edit-input(
      :label="$t('totalPages')"
      v-model="totalPages"
      type="number"
    )
    .add-button-container
      add-tsundoku-button(
        @add-tsundoku="onAddTsundoku"
      )
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import Icon from '@/components/assets/Icon.vue'
import AddTsundokuButton from '@/components/atoms/AddTsundokuButton.vue'
import AddBooksEditInput from '@/components/molecules/AddBooksEditInput.vue'
import ModalFrame from '@/components/atoms/ModalFrame.vue'
import BookCover from '@/components/atoms/BookCover.vue'

@Component({
  components: {
    Icon,
    AddTsundokuButton,
    AddBooksEditInput,
    ModalFrame,
    BookCover
  }
})
export default class AddBooksSearch extends Vue {
  public $store!: ExStore
  private title = ''
  private author = ''
  private price = ''
  private totalPages = ''
  private coverImageUrl = ''

  mounted() {
    if (typeof this.$route.query.title === 'string') {
      this.title = this.$route.query.title
    }
    if (typeof this.$route.query.author === 'string') {
      this.author = this.$route.query.author
    }
    if (typeof this.$route.query.price === 'string') {
      this.price = this.$route.query.price
    }
    if (typeof this.$route.query.totalPages === 'string') {
      this.totalPages = this.$route.query.totalPages
    }
    if (typeof this.$route.query.coverImageUrl === 'string') {
      this.coverImageUrl = this.$route.query.coverImageUrl
    }
  }

  async onAddTsundoku() {
    await this.$store.dispatch('addNewBook', {
      book: {
        id: '',
        isbn: '',
        title: this.title,
        author: [this.author],
        totalPages: parseInt(this.totalPages),
        price: parseInt(this.price),
        caption: null,
        publisher: '',
        coverImageUrl: this.coverImageUrl,
        readPages: 0,
        memo: ''
      }
    })
    await this.$store.dispatch('getMyBooks')
    this.$router.push('/')
  }
}
</script>

<style lang="sass" scoped>
.add-books-search
  -webkit-overflow-scrolling: touch

.book-cover
  margin: 20px auto 40px auto

.edit-input
  width: 85%

  .is-desktop &
    margin: 0 auto 40px auto
  .is-mobile &
    margin: 0 auto 30px auto

.add-button-container
  display: flex
  justify-content: flex-end
  width: 85%
  margin: 2rem auto
</style>
