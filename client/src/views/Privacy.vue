<template lang="pug">
  div
    | privacy
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'

@Component({
  components: {}
})
export default class Privacy extends Vue {
  public $store!: ExStore
}
</script>

<style lang="sass" scoped></style>
