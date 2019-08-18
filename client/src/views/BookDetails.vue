<template lang="pug">
  modal-frame(
    :closeColor="isEditing ? 'var(--text-gray)' : 'white'"
    :title="isEditing ? '編集' : ''"
    :path="isEditing ? '../../../' : '../../'"
    no-padding)
    .book-details(
      :class="{ editing: isEditing, 'on-transition': isOnTransitionToEdit || isOnTransitionToDetails }"
      ref="bookDetails")
      transition(name="fade")
        .header(ref="header" v-if="!isEditing")
          .header-bg(ref="headerBg" :style="headerBgStyle")
          .header-fade
          .header-filter
      transition(name="fade")
        .info(ref="info" v-if="!isEditing")
          book-major-info(:title="book.title" :authors="book.author")
      transition(name="fade")
        .actions(ref="actions" v-if="!isEditing")
          book-details-action-button.action(
            :expanded="isButtonExpanded"
            icon="pen"
            :label="$t('edit')" :iconSize="20"
            @click="onEditClick"
            v-tooltip="'本の情報を編集する'")
          book-details-action-button.action(
            :expanded="isButtonExpanded"
            icon="remove"
            :label="$t('delete')"
            @click="onDeleteClick"
            v-tooltip="'この本を削除する'")
      .cover-wrap(ref="coverWrap")
        book-cover(:url="book.coverImageUrl" has-shadow)
      transition(name="fade" mode="out-in")
        .edit-wrap(v-if="isEditing" key="edit")
          book-info-edit(v-model="book" ref="bookInfoEditInstance" :has-shadow="isEditing && !isOnTransitionToEdit")
        .detail-wrap(v-else ref="bodyWrap" @scroll="updateHeader" key="detail")
          .body(ref="body")
            .spacer
            .controller
              book-list-item-progress-controller(:book="book")
            .info-items
              book-details-item.item(
                name="ツンドク残額"
                :value="`¥ ${remainingPrice.toLocaleString()}`"
                valueColor="var(--tsundoku-red)")
              book-details-item.item(
                :name="$t('price')"
                :value="`¥ ${price.toLocaleString()}`")
              book-details-item.item(
                :name="$t('totalPages')"
                :value="`${totalPages}`")
              book-details-item.item(
                name="最後に読んだ日"
                value="0日前")
              book-details-item.item(
                textarea
                clickable
                name="メモ"
                placeholder="メモがありません"
                :value="book.memo")
              book-details-item.item(
                :name="$t('overview')"
                :value="book.caption")
              book-details-item.item(
                :name="$t('publisher')"
                :value="book.publisher")
              book-details-item.item(
                :name="$t('isbn')"
                :value="book.isbn")
</template>

<script lang="ts">
import { Vue, Component, Prop, Watch } from 'vue-property-decorator'
import { ExStore } from 'vuex'
import { BookRecord } from '../types/Book'
import BookDetailsActionButton from '@/components/atoms/BookDetailsActionButton.vue'
import BookDetailsItem from '@/components/molecules/BookDetailsItem.vue'
import ModalFrame from '@/components/atoms/ModalFrame.vue'
import BookCover from '@/components/atoms/BookCover.vue'
import BookMajorInfo from '@/components/atoms/BookMajorInfo.vue'
import BookInfoEdit from '@/components/organs/BookInfoEdit.vue'
import BookListItemProgressController from '@/components/molecules/BookListItemProgressController.vue'

const initialHeaderHeight = 200
const slimHeaderHeight = 120

const initialActionsTop = 72
const initialCoverTop = 120
const slimInfoLeftPadding = 100

const transitionDuration = 400

@Component({
  components: {
    BookDetailsActionButton,
    BookDetailsItem,
    ModalFrame,
    BookCover,
    BookInfoEdit,
    BookMajorInfo,
    BookListItemProgressController
  }
})
export default class BookDetails extends Vue {
  public $store!: ExStore
  private lastBook!: BookRecord // モーダル閉じるときのエラー対策
  private currentHeaderHeight = initialHeaderHeight
  private animationFrameRequestId?: number
  private animationEndTimeoutId?: number
  private lastShiftAmount?: number
  private isOnTransitionToEdit = false
  private isOnTransitionToDetails = false

  // router
  @Prop({ type: Boolean, default: false })
  private isEditing!: boolean

  public async mounted() {
    this.lastBook = this.book
    if (!(this.$refs.bodyWrap instanceof Element)) {
      return
    }
    window.addEventListener('resize', this.handleResize)
    this.handleResize()
    this.updateHeader()
  }

  public destroy() {
    window.removeEventListener('resize', this.handleResize)
    this.handleResize()
    if (this.animationFrameRequestId) {
      cancelAnimationFrame(this.animationFrameRequestId)
    }
    if (this.animationEndTimeoutId) {
      window.clearTimeout(this.animationEndTimeoutId)
    }
  }

  public handleResize() {
    if (
      !(this.$refs.body instanceof HTMLElement) ||
      !(this.$refs.bodyWrap instanceof HTMLElement)
    ) {
      console.log('po')
      return
    }
    // 本体がラッパー以上の場合はラッパー + ヘッダの高さ変化分までは伸ばす
    const {
      height: bodyWrapHeight
    } = this.$refs.bodyWrap.getBoundingClientRect()
    const { height: bodyHeight } = this.$refs.body.getBoundingClientRect()
    if (bodyHeight >= bodyWrapHeight) {
      this.$refs.body.style.minHeight = `${bodyWrapHeight +
        (initialHeaderHeight - slimHeaderHeight) * 2}px`
    } else {
      this.$refs.body.style.minHeight = ''
    }
  }

  public updateHeader() {
    if (
      !this.$refs.body ||
      !this.$refs.bodyWrap ||
      !this.$refs.header ||
      !this.$refs.headerBg ||
      !this.$refs.info ||
      !this.$refs.actions ||
      !this.$refs.coverWrap
    ) {
      return
    }

    const bodyElement = this.$refs.body as HTMLElement
    const bodyWrapElement = this.$refs.bodyWrap as HTMLElement
    const headerElement = this.$refs.header as HTMLElement
    const headerBgElement = this.$refs.headerBg as HTMLElement
    const infoElement = this.$refs.info as HTMLElement
    const actionsElement = this.$refs.actions as HTMLElement
    const coverWrapElement = this.$refs.coverWrap as HTMLElement

    const { top: bodyWrapTop } = bodyWrapElement.getBoundingClientRect()
    const { top: bodyTop } = bodyElement.getBoundingClientRect()

    const scrollAmount = bodyWrapTop - bodyTop
    let newHeight = initialHeaderHeight - scrollAmount * 0.5

    if (newHeight < slimHeaderHeight) {
      newHeight = slimHeaderHeight
    } else if (newHeight > initialHeaderHeight + 80) {
      newHeight = initialHeaderHeight - (scrollAmount - 80) * 0.25
    } else if (newHeight > initialHeaderHeight) {
      newHeight = initialHeaderHeight - scrollAmount * 0.5
    }

    const shiftAmount = newHeight - this.currentHeaderHeight

    // カクつき対策
    if (
      this.lastShiftAmount &&
      Math.abs(shiftAmount / this.lastShiftAmount) > 3
    ) {
      this.lastShiftAmount = newHeight - this.currentHeaderHeight
      return
    }
    if (shiftAmount === 0) {
      this.lastShiftAmount = undefined
    }

    this.lastShiftAmount = newHeight - this.currentHeaderHeight

    const progress =
      (initialHeaderHeight - newHeight) /
      (initialHeaderHeight - slimHeaderHeight)

    this.animationFrameRequestId = requestAnimationFrame(() => {
      // リアクティブガン無視コーナー
      headerElement.style.transform = `scaleY(${newHeight /
        initialHeaderHeight})`
      headerBgElement.style.transform = `scaleY(${initialHeaderHeight /
        newHeight})`
      infoElement.style.transform = `translateX(${Math.max(
        0,
        progress * slimInfoLeftPadding
      )}px)`
      actionsElement.style.transform = `translateY(${(1 - progress) *
        initialActionsTop}px)`
      coverWrapElement.style.transform = `translateY(${(1 - progress) *
        initialCoverTop}px) scale(${1 - progress / 2})`

      this.currentHeaderHeight = newHeight
    })
  }

  public async onDeleteClick() {
    await this.$store.dispatch('deleteBook', { id: this.book.id })
    this.$router.push('../../')
  }

  @Watch('isEditing')
  public async onIsEditingChanged(val: boolean) {
    if (val) {
      // 編集状態へのトランジション(手動)
      const modalElement = this.$refs.bookDetails as HTMLElement
      const coverWrapElement = this.$refs.coverWrap as HTMLElement

      await this.$nextTick()

      // 移動先計算
      const modalWidth = modalElement.clientWidth
      const coverLeft = coverWrapElement.offsetLeft
      const coverWidth = coverWrapElement.clientWidth

      const translateX = modalWidth / 2 - coverWidth / 2 - coverLeft

      // 110pxはほぼマジックナンバー
      coverWrapElement.style.transform = `translateX(${translateX}px) translateY(110px) scale(1)`

      // 終了後にcoverWrapを消す
      this.isOnTransitionToEdit = true
      window.setTimeout(() => {
        coverWrapElement.style.display = 'none'
        this.isOnTransitionToEdit = false
        coverWrapElement.style.opacity = '0'
      }, transitionDuration * 1.25)
    } else {
      const modalElement = this.$refs.bookDetails as HTMLElement
      const editAreaElement = this.$refs.bookInfoEditInstance.$el as HTMLElement
      const coverWrapElement = this.$refs.coverWrap as HTMLElement

      coverWrapElement.style.transform = ''

      const modalWidth = modalElement.getBoundingClientRect().width
      const editAreaTop = editAreaElement.getBoundingClientRect().top

      coverWrapElement.style.visibility = 'auto'
      coverWrapElement.style.display = 'block'
      coverWrapElement.style.opacity = '1'
      const {
        left: coverLeft,
        width: coverWidth
      } = coverWrapElement.getBoundingClientRect()

      const translateX = modalWidth / 2 - coverWidth / 2 - coverLeft
      const translateY = editAreaTop - 24

      coverWrapElement.style.transform = `translateX(${translateX}px) translateY(${translateY}px) scale(1)`

      await this.$nextTick()

      // coverWrapの位置をあわせてからtransitionを開始する
      this.isOnTransitionToDetails = true
      coverWrapElement.style.transform = `translateX(${0}px) translateY(${initialCoverTop}px) scale(1)`

      window.setTimeout(() => {
        this.isOnTransitionToDetails = false
      }, transitionDuration * 1.25)

      await this.$nextTick()
      this.updateHeader()
      this.handleResize()
    }
  }

  public onEditClick() {
    this.$router.push(`${this.$route.path}/edit`)
  }

  get isButtonExpanded() {
    return (
      this.currentHeaderHeight >= (initialHeaderHeight + slimHeaderHeight) / 2
    )
  }

  get headerTransitionProgress() {
    return (
      (initialHeaderHeight - this.currentHeaderHeight) /
      (initialHeaderHeight - slimHeaderHeight)
    )
  }

  get book(): BookRecord {
    const bookId = this.$route.params['id']
    const book = this.$store.getters.getBookById(bookId)
    return book || this.lastBook
  }

  get headerBgStyle() {
    return {
      backgroundImage:
        this.book && this.book.coverImageUrl !== ''
          ? `url(${this.book.coverImageUrl})`
          : 'url("https://tn.smilevideo.jp/smile?i=26783258.L")'
    }
  }

  get price() {
    return this.book && this.book.price
  }

  get totalPages() {
    return this.book && this.book.totalPages
  }

  get remainingPrice(): number {
    return Math.round(
      (1 - this.book.readPages / this.book.totalPages) * this.book.price
    )
  }
}
</script>

<style lang="sass" scoped>
$transition-duration: 0.2s
$cover-transition-duration: 0.3s

.book-details
  position: relative
  width: 100%
  height: 100%

  overflow-y: hidden
  overflow-

.header
  position: absolute
  top: 0
  left: 0
  z-index: 3

  height: 200px
  width: 100%
  padding: 24px

  overflow: hidden
  transform: scaleY(1)
  transform-origin: top left

  will-change: transform

.detail-wrap
  position: absolute
  top: 0
  height: 100%
  width: 100%
  overflow-y: scroll
  -webkit-overflow-scrolling: touch

.body
  padding-top: 200px

.header-bg
  position: absolute
  top: 0
  left: 0
  width: 100%
  height: 100%

  background:
    size: cover
    repeat: no-repeat
    position: center center

.header-filter
  position: absolute
  top: 0
  left: 0
  width: 100%
  height: 100%
  backdrop-filter: blur(15px)

.header-fade
  position: absolute
  top: 0
  left: 0
  width: 100%
  height: 100%

  background: linear-gradient(to bottom, rgba(0, 0, 0, 1) 0%, rgba(0, 0, 0, 0.7) 30%, rgba(0, 0, 0, 0) 70%)
  opacity: 0.5

.close
  display: flex
  justify-content: flex-end

  width: 100%

.info
  position: absolute
  width: calc(100% - 192px)
  top: 24px
  left: 32px
  z-index: 5

  color: white
  transform: translateZ(0)
  will-change: transform

.actions
  position: absolute
  top: 72px
  right: 0
  z-index: 5
  padding:
    top: 0
    right: 16px
    bottom: 16px
    left: 0
  display: flex
  transform: translateY(72px);

.action
  margin: 0 4px

.cover-wrap
  position: absolute
  top: 0
  left: 24px
  z-index: 4
  transform: translateY(120px) scale(1)
  will-change: transform

.spacer
  height: 60px

.controller
  padding:
    right: 24px
    left: 24px

.info-items
  padding:
    right: 32px
    left: 32px
  margin:
    top: 16px

.item
  margin:
    bottom: 16px

.book-details.editing
  .info, .actions, .header
    visibility: hidden
  &:not(.on-transition) .cover-wrap
    visibility: hidden

.book-details.on-transition .cover-wrap
  transition: all $cover-transition-duration $easeInOutQuad

.edit-wrap
  position: absolute
  top: 0
  height: calc(100% - 90px)
  width: 100%
  padding: 0 24px  // modal-frameに揃える
  margin: 90px 0 0 0
  overflow-y: scroll
  -webkit-overflow-scrolling: touch


.fade-enter-active, .fade-leave-active
  transition: opacity $transition-duration $easeInOutQuad
.fade-enter, .fade-leave-to
  opacity: 0
</style>
