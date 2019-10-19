<template lang="pug">
  div
    portal(to="modalView")
      transition(name="modal-show")
        router-view
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class UserDetails extends Vue {}
</script>

<style lang="sass"></style>
