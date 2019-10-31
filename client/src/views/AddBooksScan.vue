<template lang="pug">
  modal-frame.add-books-scan(
      close-color="white"
      no-padding
      is-content-non-scrollable
    )
    video#video(autoplay muted playsinline)
    #overlay
      #crop-area
        .barcode-reader-container
          add-book-scan-barcode-reader(:color="scannerColor")
    .info
      .scan-message.scan-message--error(v-if="state === 'nodevice'")
        | {{ $t('scan_error_no_device') }}
      .scan-message.scan-message--error(v-else-if="state === 'incorrect'")
        | {{ $t('scan_error_not_book') }}
      .scan-message(v-else-if="searchingCount >= 1")
        | {{ $t('scan_message_searching') }}
      .scan-message(v-else-if="state === 'scanned'")
        | {{ $t('scan_message_searched') }}
      .scan-message(v-else-if="state === 'known'")
        | {{ $t('scan_message_known') }}
      .scan-message(v-else-if="scannedBooks.length === 0")
        | {{ $t('scan_message_scanning') }}
      carousel.cards(
        :paginationEnabled="false"
        :perPage="1"
        :class="$store.getters.viewTypeClass"
        @page-change="handleCarouselPageChange"
      )
        slide.card-wrap(
          v-for="( book, i ) in scannedBooks"
          :style="getCardWrapStyle(i)"
          :class="`${$store.getters.viewTypeClass} ${getCardWrapClass(i)}`"
          :key="book.isbn"
        )
          add-book-card.card(
            :book="book"
            type="scan"
            @appear-start="handleAppearStart"
            @appear="handleAppear"
            @to-remove="handleCardToRemove(i)"
            @remove="handleCardRemove(i)"
          )
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { ExStore } from 'vuex'
import {
  BrowserBarcodeReader,
  BarcodeFormat,
  DecodeHintType,
  Result
} from '@zxing/library'
import { BookRecord } from '../types/Book'
import { Carousel, Slide } from 'vue-carousel'

import AddBookScanBarcodeReader from '@/components/atoms/AddBookScanBarcodeReader.vue'
import AddBookCard from '@/components/molecules/AddBookCard.vue'
import ModalFrame from '@/components/atoms/ModalFrame.vue'

const codeReader = new BrowserBarcodeReader(
  500,
  new Map([[DecodeHintType.POSSIBLE_FORMATS, [BarcodeFormat.EAN_13]]])
)

const stateResetMs = 2000

interface VideoInputDevice {
  deviceId: string
  groupId: string
  kind: string
  label: string
}

type ScanState =
  | 'nodevice'
  | 'scanning'
  | 'incorrect'
  | 'scanned'
  | 'known'
  | 'noresult'

const stateColorMap: Record<ScanState, string> = {
  nodevice: 'rgba(255, 255, 255, 0.5)',
  scanning: 'white',
  incorrect: 'var(--tsundoku-red)',
  scanned: 'var(--succeed-blue)',
  known: 'rgba(255, 255, 255, 0.5)',
  noresult: 'var(--tsundoku-red)'
}

@Component({
  components: {
    Carousel,
    Slide,
    ModalFrame,
    AddBookCard,
    AddBookScanBarcodeReader
  }
})
export default class AddBooksScan extends Vue {
  $store!: ExStore

  videoInputDevices: VideoInputDevice[] = []
  captureIntervalID = 0
  stateResetTimeoutId = 0

  scannedBooksMap: Record<string, BookRecord> = {}
  scannedBooks: BookRecord[] = []

  state: ScanState = 'scanning'

  searchingCount = 0

  toAddIndex = 0
  isCardToAppear = false
  isCardAppearing = false

  toRemoveIndex = -1

  cardShiftWidth = 0

  requestAnimationFrameID = 0
  stream: MediaStream | null = null

  async mounted() {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    ;(window as any).addByIsbn = (isbn: string) =>
      this.barcodeScanned({
        getText() {
          return isbn
        }
      } as any) // eslint-disable-line @typescript-eslint/no-explicit-any
    if (!codeReader.isMediaDevicesSuported) {
      this.setScanState('nodevice')
      return
    }

    try {
      const videoInputDevices = await codeReader.listVideoInputDevices()
      this.videoInputDevices = videoInputDevices
    } catch (err) {
      console.error(err)
      this.setScanState('nodevice')
    }

    const video = this.$el.querySelector('video')
    if (!video) {
      return
    }

    this.stream = await navigator.mediaDevices.getUserMedia({
      video: {
        facingMode: {
          ideal: 'environment'
        }
      },
      audio: false
    })
    video.srcObject = this.stream

    await this.$nextTick()

    const canvas = codeReader.createCaptureCanvas()
    const cropArea = this.$el.querySelector('#crop-area')
    if (!cropArea) {
      return
    }
    canvas.width = cropArea.clientWidth
    canvas.height = cropArea.clientHeight
    canvas.classList.add('barcode-reader-canvas')
    cropArea.appendChild(canvas)
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const anyCropArea = cropArea as any

    const ctx = canvas.getContext('2d')

    const getDrawArea = () => {
      if (
        video.videoWidth / video.videoHeight <
        video.clientWidth / video.clientHeight
      ) {
        const sx =
          video.videoWidth * (anyCropArea.offsetLeft / video.clientWidth)
        const sw =
          video.videoWidth * (anyCropArea.clientWidth / video.clientWidth)

        const realHeight =
          (video.clientWidth * video.videoHeight) / video.videoWidth
        const sy =
          video.videoHeight *
          ((realHeight - anyCropArea.clientHeight) / 2 / realHeight)
        const sh = video.videoHeight * (anyCropArea.offsetHeight / realHeight)

        return [sx, sy, sw, sh]
      } else {
        const sy =
          video.videoHeight * (anyCropArea.offsetTop / video.clientHeight)
        const sh =
          video.videoHeight * (anyCropArea.clientHeight / video.clientHeight)

        const realWidth =
          video.clientHeight * (video.videoWidth / video.videoHeight)
        const sx =
          video.videoWidth *
          ((realWidth - anyCropArea.clientWidth) / 2 / realWidth)
        const sw = video.videoWidth * (anyCropArea.offsetWidth / realWidth)

        return [sx, sy, sw, sh]
      }
    }

    const draw = () => {
      const [sx, sy, sw, sh] = getDrawArea()
      if (!ctx) {
        return
      }
      ctx.drawImage(video, sx, sy, sw, sh, 0, 0, canvas.width, canvas.height)
      this.requestAnimationFrameID = requestAnimationFrame(draw)
    }

    this.captureIntervalID = window.setInterval(async () => {
      try {
        const scanResult = await codeReader.decodeFromImageUrl(
          canvas.toDataURL()
        )
        this.barcodeScanned(scanResult)
      } catch (_) {
        //
      }
    }, 500)
    draw()
  }

  beforeDestroy() {
    cancelAnimationFrame(this.requestAnimationFrameID)
    clearInterval(this.captureIntervalID)
    codeReader.reset()
    if (this.stream) {
      this.stream.getVideoTracks().forEach((track: MediaStreamTrack) => {
        track.stop()
      })
    }
  }

  setScanState(state: ScanState) {
    this.state = state
    if (state !== 'scanning' && state !== 'nodevice') {
      if (this.stateResetTimeoutId !== 0) {
        window.clearTimeout(this.stateResetTimeoutId)
      }
      this.stateResetTimeoutId = window.setTimeout(
        () => (this.state = 'scanning'),
        stateResetMs
      )
    }
  }

  syncronizeCardWidth() {
    const card = document.querySelector('.cards')
    if (card) {
      this.cardShiftWidth = card.clientWidth
    }
  }

  async addCard(book: BookRecord) {
    // 追加直後は追加位置以降のカードを左にずらす
    this.syncronizeCardWidth()
    this.isCardToAppear = true

    this.scannedBooks.splice(this.toAddIndex, 0, book)
    this.$set(this.scannedBooksMap, book.isbn, book)
    this.isCardToAppear = true

    await this.$nextTick()
    this.isCardAppearing = true
    this.isCardToAppear = false
  }

  deleteCardAt(index: number) {
    this.scannedBooks.splice(index, 1)
  }

  async barcodeScanned(scanResult: Result) {
    const isbn = scanResult.getText()
    const prefix = isbn.substring(0, 3)
    if (isbn in this.scannedBooksMap) {
      // スキャンに成功した直後に読み続けている場合はscannedにする
      this.setScanState(this.state === 'scanned' ? 'scanned' : 'known')
      return
    }
    if (prefix !== '978' && prefix !== '979') {
      this.setScanState('incorrect')
      return
    }

    this.setScanState('scanned')

    // 競合するかも?
    this.searchingCount += 1
    await this.$nextTick()
    const searchResult = await this.$store.dispatch('searchBooksByISBN', {
      isbn
    })
    await this.$nextTick()
    this.searchingCount -= 1

    if (searchResult.length !== 1) {
      this.setScanState('noresult')
      return
    }
    if (!(isbn in this.scannedBooksMap)) {
      this.addCard(searchResult[0])
    }
  }

  handleAppearStart() {
    this.isCardAppearing = true
  }

  handleAppear() {
    this.isCardAppearing = false
  }

  handleCardToRemove(index: number) {
    this.syncronizeCardWidth()
    this.toRemoveIndex = index
  }

  async handleCardRemove(index: number) {
    this.toRemoveIndex = -1
    this.cardShiftWidth = 0
    await this.$nextTick()
    this.deleteCardAt(index)
  }

  handleCarouselPageChange(page: number) {
    this.toAddIndex = page
  }

  getCardWrapStyle(index: number) {
    if (this.toRemoveIndex >= 0 && index > this.toRemoveIndex) {
      // カード削除中は削除位置以降を左にずらす
      return {
        transform: `translateX(-${this.cardShiftWidth}px)`
      }
    }
    if (this.isCardToAppear && index > this.toAddIndex) {
      // カード追加直後は追加位置以降を左にずらす
      return {
        transform: `translateX(-${this.cardShiftWidth}px)`
      }
    }
    return {}
  }

  getCardWrapClass(index: number) {
    if (this.toRemoveIndex >= 0 && index > this.toRemoveIndex) {
      return 'to-transition'
    }
    if (this.isCardAppearing && index > this.toAddIndex) {
      return 'to-transition'
    }
    return ''
  }

  get scannerColor() {
    return stateColorMap[this.state]
  }
}
</script>

<style lang="sass" scoped>
#video
  position: absolute
  top: 0
  left: 0
  width: 100%
  height: 100%
  object-fit: cover

#overlay
  position: absolute
  top: 0
  left: 0
  width: 100%
  height: 100%
  background-color: rgba(0, 0, 0, 0.5)
  object-fit: cover
  display: flex
  justify-content: center
  align-items: center

#crop-area
  position: relative
  width: 480px
  height: 270px
  max-width: 80vw
  max-height: 45vw
  border-radius: 8px

.barcode-reader-container
  display: flex
  align-items: center
  justify-content: center

  position: absolute
  top: 0
  left: 0

  width: 100%
  height: 100%

  padding: 16px

.info
  position: relative
  display: flex
  flex-flow: column nowrap
  align-items: center
  top: 0
  height: 100%
  width: 100%
  padding-top: 24px    // .modal-frame-topの高さ分

.scan-message
  display: flex
  align-items: center
  justify-content: center

  width: 86%
  max-width: 400px

  margin:
    top: 48px
    left: 16px
    bottom: 0
    right: 16px

  padding: 16px 32px

  background: rgba(0, 0, 0, 0.1)
  backdrop-filter: blur(8px)
  border-radius: 12px

  color: white
  font:
    weight: bold

  &--error
    background: $tsundoku-red-bg


$card-height: 200px
$card-margin-lg: 16px
$card-margin-sm: 8px
$card-margin-pc: 8px
$card-small-margin: 4px
$carousel-width-sp-sm: 300px
$carousel-width-breakpoint-sp: 500px
$carousel-margin-sp: 24px
$carousel-margin-pc: 48px

.cards
  position: absolute
  left: 0
  bottom: 0
  height: $card-height

  // カルーセルを幅を狭めて表示し、overflow: visibleをカルーセルに指定して横のカードを見せる
  &.is-desktop
    width: calc(100% - #{$carousel-margin-pc * 2} + #{$card-margin-pc * 2})
    margin: 0 $carousel-margin-pc - $card-margin-pc

  &.is-mobile
    // モバイルの場合はカルーセルを300pxで表示し、マージンを調整
    width: $carousel-width-sp-sm + $card-margin-lg * 2
    margin: 0 calc(50% - #{$carousel-width-sp-sm / 2})

    @media (max-width: #{$carousel-width-breakpoint-sp})
      width: $carousel-width-sp-sm + $card-margin-sm * 2

  overflow: visible

.card-wrap
  display: flex
  align-items: center
  justify-content: center

  height: $card-height

  &.is-desktop
    padding: 0 $card-margin-pc

  &.is-mobile
    padding: 0 $card-margin-lg
    @media (max-width: #{$carousel-width-breakpoint-sp})
      padding: 0 $card-margin-sm

  border-radius: 8px

  &.to-transition
    transition: transform 1s $easeInOutQuint
</style>

<style lang="sass">
.barcode-reader-canvas
  border-radius: 16px

.add-books-scan
  .VueCarousel-wrapper
    overflow: visible!important
  .VueCarousel-inner
    overflow: visible!important
</style>
