<template lang="pug">
  .book-details
    .header
      .close
        router-link.icon-close(:to="parentPath")
          icon-plus(color='white' :height="16")
      .info
        book-major-info(:book="book")
      .cover-wrap
        book-cover(:book="book")
    .body
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'
import BookCover from '@/components/atoms/BookCover.vue'
import BookMajorInfo from '@/components/atoms/BookMajorInfo.vue'
import IconPlus from '@/components/assets/IconPlus.vue'
import Vibrant from 'node-vibrant'
import { BookRecord } from '../types/Book'
import { RouteRecord } from 'vue-router'

@Component({
  components: {
    BookCover,
    BookMajorInfo,
    IconPlus
  }
})
export default class BookDetails extends Vue {
  public $store!: ExStore
  private vibrant!: Vibrant

  public async mounted() {}

  get book(): BookRecord {
    const bookId = this.$route.params['id']
    return this.$store.getters.getBookById(bookId)
  }

  get parentPath(): RouteRecord {
    const matched = this.$route.matched
    return matched[0]
  }
}
</script>

<style lang="sass" scoped>
// FIXME
.book-details
  position: fixed
  top: 0
  right: 0
  height: 100vh
  width: 320px
  background: white

.header
  position: relative

  width: 100%
  height: 240px
  padding: 24px

  color: white
  background: linear-gradient(180deg, #2488D0 0%, #56CCF2 100%)

.close
  display: flex
  justify-content: flex-end

  width: 100%

.icon-close
  transform: rotate(45deg)

.info
  margin: 24px 0

.cover-wrap
  position: absolute
  bottom: -96px
</style>
