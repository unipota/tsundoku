<template lang="pug">
  modal-frame
    div
      | add books search
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

import ModalFrame from '@/components/atoms/ModalFrame.vue'

@Component({
  components: { ModalFrame }
})
export default class AddBooksSearch extends Vue {}
</script>

<style lang="sass"></style>
