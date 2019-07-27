<template lang="pug">
  modal-frame
    div
      | add books scan
    video#video
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { BrowserBarcodeReader } from '@zxing/library'

import ModalFrame from '@/components/atoms/ModalFrame.vue'

const codeReader = new BrowserBarcodeReader()

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

    codeReader
      .decodeFromInputVideoDevice(undefined, 'video')
      .then(result => alert(result))
      .catch(err => console.error(err))
  }

  beforeDestroy() {
    codeReader.reset()
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
</style>
