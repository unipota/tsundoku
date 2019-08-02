<template lang="pug">
  .body
    .progress-wrap
      book-list-item-progress(:book="book")
    .record-read-pages-button-wrap
      record-read-pages-button
    .check-button-wrap
      check-button
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { BookRecord } from '../../types/Book'

import BookListItemProgress from '@/components/molecules/BookListItemProgress.vue'
import RecordReadPagesButton from '@/components/atoms/RecordReadPagesButton.vue'
import CheckButton from '@/components/atoms/CheckButton.vue'

@Component({
  components: { BookListItemProgress, RecordReadPagesButton, CheckButton }
})
export default class BookListItemProgressContainer extends Vue {
  @Prop({ type: Object, required: true })
  private book!: BookRecord
}
</script>

<style lang="sass" scoped>
.body
  display: flex
  align-items: center
  flex-wrap: nowrap
  margin:
    top: 4px

.progress-wrap
  margin-right: 8px
  flex-grow: 1
  max-width: 250px
  margin:
    left: auto

.record-read-pages-button-wrap
  margin-right: 8px
  flex-shrink: -1

.check-button-wrap
  flex-shrink: -1
</style>
