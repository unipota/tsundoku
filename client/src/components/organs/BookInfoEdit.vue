<template lang="pug">
  .book-info-edit
    book-cover(
      :hasShadow="hasShadow"
      :url="value.coverImageUrl"
    )
    books-info-edit-input.edit-input(
      :label="$t('title')"
      :value="value.title"
      @input="update('title', $event)"
    )
    books-info-edit-input.edit-input(
      :label="$t('author')"
      :value="value.author[0]"
      @input="update('author', [$event])"
    )
    books-info-edit-input.edit-input(
      type="number"
      :label="$t('price')"
      :value="value.price.toString()"
      @input="update('price', parseInt($event))"
    )
    books-info-edit-input.edit-input(
      type="number"
      :label="$t('totalPages')"
      :value="value.totalPages.toString()"
      @input="update('totalPages', parseInt($event))"
    )
    books-info-edit-input.edit-input(
      :label="$t('publisher')"
      :value="value.publisher"
      @input="update('publisher', $event)"
    )
    books-info-edit-input.edit-input(
      type="textarea"
      :label="$t('memo')"
      :value="value.memo.toString()"
      @input="update('memo', $event)"
    )
    books-info-edit-input.edit-input(
      type="textarea"
      :label="$t('caption')"
      :value="value.caption.toString()"
      @input="update('caption', $event)"
    )
    books-info-edit-input.edit-input(
      :label="$t('coverImageUrl')"
      :value="value.coverImageUrl"
      @input="update('coverImageUrl', $event)"
    )
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { BookDetail } from '../../types/Book'

import Icon from '@/components/assets/Icon.vue'
import BooksInfoEditInput from '@/components/molecules/BooksInfoEditInput.vue'
import BookCover from '@/components/atoms/BookCover.vue'

@Component({
  components: {
    Icon,
    BooksInfoEditInput,
    BookCover
  }
})
export default class BookInfoEdit extends Vue {
  @Prop({ type: Object, required: true })
  private value!: BookDetail

  @Prop({ type: Boolean, default: true })
  private hasShadow!: boolean

  public update<T extends keyof BookDetail>(key: T, value: BookDetail[T]) {
    this.$emit('input', {
      ...this.value,
      [key]: value
    })
  }
}
</script>

<style lang="sass" scoped>
.book-cover
  margin: 20px auto 40px

.edit-input
  .is-desktop &
    width: 85%
    margin: 0 auto 40px
  .is-mobile &
    margin: 0 auto 30px
</style>
