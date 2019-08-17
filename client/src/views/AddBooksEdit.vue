<template lang="pug">
  modal-frame.add-books-edit(
    :title="$t('addBooksSearchTitle')"
    :class="$store.getters.viewTypeClass"
  )
    book-info-edit(v-model="book")
    .add-button-container
      add-tsundoku-button(
        @add-tsundoku="onAddTsundoku"
      )
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import AddTsundokuButton from '@/components/atoms/AddTsundokuButton.vue'
import BookInfoEdit from '@/components/organs/BookInfoEdit.vue'
import { BookRecord } from '../types/Book'

@Component({
  components: {
    AddTsundokuButton,
    BookInfoEdit
  }
})
export default class AddBooksSearch extends Vue {
  public $store!: ExStore

  private book: BookRecord = {
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
    memo: '',
    createdAt: '',
    updatedAt: ''
  }

  mounted() {
    if (typeof this.$route.query.title === 'string') {
      this.$set(this.book, 'title', this.$route.query.title)
    }
    if (typeof this.$route.query.author === 'string') {
      this.$set(this.book, 'author', [this.$route.query.author])
    }
    if (typeof this.$route.query.price === 'string') {
      this.$set(this.book, 'price', parseInt(this.$route.query.price))
    }
    if (typeof this.$route.query.totalPages === 'string') {
      this.$set(this.book, 'totalPages', parseInt(this.$route.query.totalPages))
    }
    if (typeof this.$route.query.coverImageUrl === 'string') {
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
