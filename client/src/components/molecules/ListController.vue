<template lang="pug">
  .list-controller
    .text-input-wrap
      text-input(
        :value="filterText" 
        @input="handleInput"
        placeholder="タイトルや著者など"
        withClearButton)
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

import TextInput from '@/components/atoms/TextInput.vue'

@Component({
  components: { TextInput }
})
export default class ListController extends Vue {
  @Prop({ type: String, required: true })
  private filterText: string = ''

  handleInput(value: string) {
    this.$emit('update:filterText', value)
  }
}
</script>

<style lang="sass" scoped>
.list-controller
  padding:
    top: 16px
    left: 16px
    right: 16px

.text-input-wrap
  max-width: 480px
  margin: auto
</style>
