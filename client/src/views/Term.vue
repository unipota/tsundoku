<template lang="pug">
  div
    | term
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'

@Component({
  components: {}
})
export default class Term extends Vue {
  public $store!: ExStore
}
</script>

<style lang="sass" scoped></style>
