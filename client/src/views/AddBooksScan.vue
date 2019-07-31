<template lang="pug">
  modal-frame(close-color="white" no-padding)
    video#video(autoplay muted playsinline)
    #overlay
      #crop-area
        .barcode-reader-container
          add-book-scan-barcode-reader(:color="scannerColor")
    .info
      vue-scroll-snap.cards(horizontal)
        .scan-card-wrap(v-for="book in scannedBooks" :key="book.isbn")
          scan-card(:book="book")
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import AddBookScanBarcodeReader from '@/components/atoms/AddBookScanBarcodeReader.vue'
import ScanCard from '@/components/molecules/ScanCard.vue'
import ModalFrame from '@/components/atoms/ModalFrame.vue'
import VueScrollSnap from 'vue-scroll-snap'

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

type ScanState = 'scanning' | 'incorrect' | 'scanned' | 'known' | 'noresult'

const stateColorMap: Record<ScanState, string> = {
  scanning: 'white',
  incorrect: 'var(--tsundoku-red)',
  scanned: 'var(--succeed-blue)',
  known: 'rgba(255, 255, 255, 0.5)',
  noresult: 'rgba(255, 255, 255, 0.5)'
}

@Component({
  components: {
    VueScrollSnap,
    ModalFrame,
    ScanCard,
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

  async mounted() {
    console.log(codeReader)
    if (!codeReader.isMediaDevicesSuported) {
      alert('getUserMedia not supported.')
      return
    }

    try {
      const videoInputDevices = await codeReader.listVideoInputDevices()
      console.log(videoInputDevices)
      this.videoInputDevices = videoInputDevices
    } catch (err) {
      console.error(err)
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

    const draw = () => {
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
      }
      requestAnimationFrame(draw)
    }

    this.captureIntervalID = window.setInterval(() => {
      codeReader
        .decodeFromImageUrl(canvas.toDataURL())
        .then(result => {
          this.barcodeScanned(result)
        })
        .catch(e => {})
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
    if (state !== 'scanning') {
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
    this.scannedBooks.push(searchResult.data[0])
    this.$set(this.scannedBooksMap, isbn, searchResult.data[0])
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
  width: 320px
  height: 180px
  max-width: 80vw
  max-height: 45vw
  border-radius: 8px

.barcode-reader-container
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

.cards
  position: absolute
  bottom: 0
  height: 160px
  width: 100%
  display: flex
  flex-direction: row
  overflow-x: scroll
  align-items: flex-end

.scan-card-wrap
  margin: 0 16px

  &:first-child
    margin-left: 0
    padding-left: 48px
  &:last-child
    margin-right: 0
    padding-right: 48px
</style>

<style lang="sass">
.barcode-reader-canvas
  border-radius: 16px
</style>
