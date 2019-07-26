<template lang="pug">
  div.book-cover
    v-lazy-image(v-if="book && book.coverImageUrl" :src="book.coverImageUrl").cover-image
</template>

<script lang="ts">
import VLazyImage from 'v-lazy-image'
import { Vue, Component, Prop } from 'vue-property-decorator'
import { BookRecord } from '../../types/Book'

@Component({
  components: {
    VLazyImage
  }
})
export default class BookCover extends Vue {
  @Prop({ type: Object, required: true })
  private book!: BookRecord
}
</script>

<style lang="sass">
.book-cover
  width: 100px
  height: 140px
  border-radius: 10px
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25)
  background-color: var(--text-white)
  overflow: hidden

.cover-image
  width: 100%
  height: 100%
  transition: filter .3s .1s, opacity .1s
  will-change: filter, opacity
  opacity: 0

  &.v-lazy-image
    filter: blur(1px)

  &.v-lazy-image-loaded
    opacity: 1
    filter: blur(0)
</style>
