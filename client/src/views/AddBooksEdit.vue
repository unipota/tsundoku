<template lang="pug">
  modal-frame.add-books-edit(
    :title="$t('addBooksSearchTitle')"
    :class="$store.getters.viewTypeClass"
  )
    book-info-edit(v-model="book")
    .add-button-container
      book-info-edit-button(
        @add-tsundoku="onAddTsundoku"
      )
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'
import { Route } from 'vue-router'

import ModalFrame from '@/components/atoms/ModalFrame.vue'
import BookInfoEditButton from '@/components/atoms/BookInfoEditButton.vue'
import BookInfoEdit from '@/components/organs/BookInfoEdit.vue'
import { BookSimpleRecord } from '../types/Book'

@Component({
  components: {
    ModalFrame,
    BookInfoEditButton,
    BookInfoEdit
  }
})
export default class AddBooksSearch extends Vue {
  public $store!: ExStore
  public $route!: Route

  private book: BookSimpleRecord = {
    id: '',
    isbn: '',
    title: '',
    author: [],
    totalPages: 0,
    price: 0,
    caption: null,
    publisher: '',
    coverImageUrl: '',
    readPages: 0,
    memo: ''
  }

  mounted() {
    if (this.$route.query.title) {
      this.$set(this.book, 'title', this.$route.query.title)
    }
    if (this.$route.query.author) {
      this.$set(this.book, 'author', [this.$route.query.author])
    }
    if (this.$route.query.price) {
      this.$set(this.book, 'price', parseInt(this.$route.query.price as string))
    }
    if (this.$route.query.publisher) {
      this.$set(this.book, 'publisher', this.$route.query.publisher)
    }
    if (this.$route.query.isbn) {
      this.$set(this.book, 'isbn', this.$route.query.isbn)
    }
    if (this.$route.query.coverImageUrl) {
      this.$set(this.book, 'coverImageUrl', this.$route.query.coverImageUrl)
    }
  }

  async onAddTsundoku() {
    await this.$store.dispatch('addNewBook', {
      book: this.book
    })
    await this.$store.dispatch('getMyBooks')
    this.$router.push('/')
  }
}
</script>

<style lang="sass" scoped>
.add-books-search
  -webkit-overflow-scrolling: touch

.add-button-container
  display: flex
  justify-content: flex-end
  margin: 2rem auto

  .is-desktop &
    width: 85%
</style>
