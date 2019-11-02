<template lang="pug">
  popup-modal-frame(name="privacy" title="プライバシーポリシー")
    .privacy-policy-modal
      
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

import Icon from '@/components/assets/Icon.vue'
import PopupModalFrame from '@/components/atoms/PopupModalFrame.vue'

@Component({ components: { Icon, PopupModalFrame } })
export default class PrivacyPolicyModal extends Vue {}
</script>

<style lang="sass" scoped>
.setting-modal
  display: flex
  flex-flow: column
  align-items: center
</style>
