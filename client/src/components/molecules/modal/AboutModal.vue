<template lang="pug">
  popup-modal-frame(name="about" title="ツンドクについて")
    .about-modal
      
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

import Icon from '@/components/assets/Icon.vue'
import PopupModalFrame from '@/components/atoms/PopupModalFrame.vue'

@Component({ components: { Icon, PopupModalFrame } })
export default class AboutModal extends Vue {}
</script>

<style lang="sass" scoped>
.about-modal
  display: flex
  flex-flow: column
  align-items: center
</style>
