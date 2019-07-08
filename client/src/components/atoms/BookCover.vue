<template lang="pug">
  v-lazy-image(:src="book.coverImageUrl").cover-image.v-lazy-image.blur
</template>

<script lang="ts">
import VLazyImage from 'v-lazy-image'
import { Vue, Component, Prop } from 'vue-property-decorator'
import { BookRecord } from '../../types/Book'

@Component({
  components: {
    VLazyImage,
  },
})
export default class BookCover extends Vue {
  @Prop({ type: Object, required: true })
  private book!: BookRecord
}
</script>

<style lang="sass">
.cover-image
  width: 101px
  border-radius: 10px
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25)
.v-lazy-image.blur
  filter: blur(5px)
  transition: filter 1s
  will-change: filter
.v-lazy-image-loaded.blur
  filter: blur(0)
</style>
