<template lang="pug">
  modal-frame.add-books-scan(close-color="white" no-padding)
    video#video(autoplay muted playsinline)
    #overlay
      #crop-area
        .barcode-reader-container
          add-book-scan-barcode-reader(:color="scannerColor")
    .info
      carousel.cards(
        :paginationEnabled="false"
        :perPage="1"
        :class="$store.getters.viewTypeClass"
      )
        slide.card-wrap(
          v-for="( book, i ) in scannedBooks"
          :style="getCardWrapStyle(i)"
          :class="getCardWrapClass(i)"
          :key="book.isbn"
        )
          add-book-card.card(
            :book="book"
            type="scan"
            @to-remove="handleCardToRemove(i)"
            @remove="handleCardRemove(i)"
          )
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import AddBookScanBarcodeReader from '@/components/atoms/AddBookScanBarcodeReader.vue'
import AddBookCard from '@/components/molecules/AddBookCard.vue'
import ModalFrame from '@/components/atoms/ModalFrame.vue'
import { Carousel, Slide } from 'vue-carousel'

import api from '@/store/general/api'

import {
  BrowserBarcodeReader,
  BarcodeFormat,
  DecodeHintType,
  Result
} from '@zxing/library'
import { BookRecord } from '../types/Book'

const codeReader = new BrowserBarcodeReader(
  500,
  new Map([[DecodeHintType.POSSIBLE_FORMATS, [BarcodeFormat.EAN_13]]])
)

const stateResetMs = 1000

interface VideoInputDevice {
  deviceId: string
  groupId: string
  kind: string
  label: string
}

type ScanState = 'nodevice' | 'scanning' | 'incorrect' | 'scanned' | 'known' | 'noresult'

const stateColorMap: Record<ScanState, string> = {
  'nodevice': 'rgba(255, 255, 255, 0.5)',
  'scanning': 'white',
  'incorrect': 'var(--tsundoku-red)',
  'scanned': 'var(--succeed-blue)',
  'known': 'rgba(255, 255, 255, 0.5)',
  'noresult': 'var(--tsundoku-red)',
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
  videoInputDevices: VideoInputDevice[] = []
  captureIntervalID = 0
  stateResetTimeoutId = 0

  scannedBooksMap: Record<string, BookRecord> = {}
  scannedBooks: BookRecord[] = []

  state: ScanState = 'scanning'

  toRemoveIndex = -1
  cardShiftWidth = 0

  async mounted() {
    ;(window as any).addByIsbn = (isbn: string) => this.barcodeScanned({ getText() { return isbn } } as any)
    console.log(codeReader)
    if (!codeReader.isMediaDevicesSuported) {
      alert('getUserMedia not supported.')
      this.setScanState('nodevice')
      return
    }

    try {
      const videoInputDevices = await codeReader.listVideoInputDevices()
      console.log(videoInputDevices)
      this.videoInputDevices = videoInputDevices
    } catch (err) {
      console.error(err)
      this.setScanState('nodevice')
    }

    const video = this.$el.querySelector('video')
    if (!video) {
      return
    }

    const stream = await navigator.mediaDevices.getUserMedia({
      video: {
        facingMode: {
          ideal: 'environment'
        }
      },
      audio: false
    })
    video.srcObject = stream

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
        const sh =
          video.videoHeight * (anyCropArea.offsetHeight / realHeight)

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
        const sw =
          video.videoWidth * (anyCropArea.offsetWidth / realWidth)

        return [sx, sy, sw, sh]
      }
    }

    const draw = () => {
      const [sx, sy, sw, sh] = getDrawArea()
      ctx!!.drawImage(
        video,
        sx,
        sy,
        sw,
        sh,
        0,
        0,
        canvas.width,
        canvas.height
      )
      requestAnimationFrame(draw)
    }

    this.captureIntervalID = window.setInterval(async () => {
      try {
        const scanResult = await codeReader.decodeFromImageUrl(canvas.toDataURL())
        this.barcodeScanned(scanResult)
      }
      catch (_) {}
    }, 500)
    draw()
  }

  beforeDestroy() {
    codeReader.reset()
    clearInterval(this.captureIntervalID)
    console.log('Reset')
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
    const searchResult = await api.searchBooksByISBN(isbn)

    if (searchResult.data.length !== 1) {
      this.setScanState('noresult')
      return
    }
    if (!(isbn in this.scannedBooksMap)) {
      this.scannedBooks.push(searchResult.data[0])
      this.$set(this.scannedBooksMap, isbn, searchResult.data[0])
    }
  }

  handleCardToRemove(index: number) {
    this.toRemoveIndex = index
    const card = document.querySelector('.card-wrap')
    if (card) {
      this.cardShiftWidth = card.clientWidth
    }
  }

  async handleCardRemove(index: number) {
    this.toRemoveIndex = -1
    this.cardShiftWidth = 0
    await this.$nextTick()
    this.scannedBooks.splice(index, 1)
  }

  getCardWrapStyle(index: number) {
    if (this.toRemoveIndex < 0 || index <= this.toRemoveIndex) {
      return {}
    }
    return {
      transform: `translateX(-${this.cardShiftWidth}px)`,
    }
  }

  getCardWrapClass(index: number) {
    if (this.toRemoveIndex < 0 || index <= this.toRemoveIndex) {
      return ''
    }
    return 'to-transition'
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
  top: 0
  height: 100%
  width: 100%

$card-height: 200px
$card-margin: 8px
$card-small-margin: 4px
$carousel-margin-sp: 24px
$carousel-margin-pc: 40px

.cards
  position: absolute
  bottom: 0
  height: $card-height

  // カルーセルを幅を狭めて表示し、overflow: visibleをカルーセルに指定して横のカードを見せる
  &.is-desktop
    width: calc(100% - #{$carousel-margin-pc * 2})
    margin: 0 $carousel-margin-pc

  &.is-mobile
    width: calc(100% - #{$carousel-margin-sp * 2})
    margin: 0 $carousel-margin-sp

    @media (max-width: #{300px + $carousel-margin-sp * 2})
      // 小さいデバイスの場合はカルーセルを300pxで表示し、マージンを調整
      width: 300px
      margin: 0 calc(50% - 150px)

  overflow: visible

.card-wrap
  display: flex
  align-items: center
  justify-content: center

  height: $card-height

  padding: 0 $card-margin

  &.is-mobile
    @media (max-width: #{300px + $carousel-margin-sp * 2})
      padding: 0 $card-small-margin

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
