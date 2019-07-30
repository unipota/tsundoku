import DecodeHintType from '@zxing/library/esm5/core/DecodeHintType' import
DecodeHintType from '@zxing/library/esm5/core/DecodeHintType'
<template lang="pug">
  modal-frame
    div
      | add books scan
    video#video(autoplay)
    div#overlay
      div#crop-area
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { BrowserBarcodeReader } from '@zxing/library'

import ModalFrame from '@/components/atoms/ModalFrame.vue'
import DecodeHintType from '@zxing/library/esm5/core/DecodeHintType'

const codeReader = new BrowserBarcodeReader(
  500,
  new Map([[DecodeHintType.PURE_BARCODE, true]])
)

interface VideoInputDevice {
  deviceId: string
  groupId: string
  kind: string
  label: string
}

@Component({
  components: { ModalFrame }
})
export default class AddBooksScan extends Vue {
  videoInputDevices: VideoInputDevice[] = []
  captureIntervalID: any = 0

  mounted() {
    console.log(codeReader)
    if (!codeReader.isMediaDevicesSuported) {
      alert('getUserMedia not supported.')
      return
    }

    codeReader
      .listVideoInputDevices()
      .then(videoInputDevices => {
        console.log(videoInputDevices)
        this.videoInputDevices = videoInputDevices
      })
      .catch(err => console.error(err))

    const video = this.$el.querySelector('video')
    if (!video) {
      return
    }

    navigator.mediaDevices
      .getUserMedia({
        video: true,
        audio: false
      })
      .then(stream => {
        video.srcObject = stream

        this.$nextTick(() => {
          const canvas = codeReader.createCaptureCanvas()
          const cropArea = this.$el.querySelector('#crop-area')
          if (!cropArea) {
            return
          }
          canvas.width = cropArea.clientWidth
          canvas.height = cropArea.clientHeight
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
              const sh =
                video.videoHeight * (anyCropArea.offsetHeight / realHeight)
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
                video.videoHeight *
                (anyCropArea.clientHeight / video.clientHeight)

              const realWidth =
                video.clientHeight * (video.videoWidth / video.videoHeight)
              const sx =
                video.videoWidth *
                ((realWidth - anyCropArea.clientWidth) / 2 / realWidth)
              const sw =
                video.videoWidth * (anyCropArea.offsetWidth / realWidth)
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
          this.captureIntervalID = setInterval(() => {
            codeReader
              .decodeFromImageUrl(canvas.toDataURL())
              .then((result: any) => {
                console.log(result)
              })
              .catch(e => {})
          }, 500)
          draw()
        })
      })
  }

  beforeDestroy() {
    codeReader.reset()
    clearInterval(this.captureIntervalID)
    console.log('Reset')
  }
}
</script>

<style lang="sass">
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
  width: 70vw
  height: 20vh
</style>
