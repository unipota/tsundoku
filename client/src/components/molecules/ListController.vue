<template lang="pug">
  .list-controller
    .text-input-wrap
      text-input(
        v-model="filterText" 
        placeholder="タイトルや著者でフィルター"
        withClearButton)
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

import TextInput from '@/components/atoms/TextInput.vue'

@Component({
  components: { TextInput }
})
export default class ListController extends Vue {
  filterText: string = ''
}
</script>

<style lang="sass" scoped>
.list-controller
  padding:
    top: 16px
    left: 16px
    right: 16px

.text-input-wrap
</style>
