<template lang="pug">
  div
    portal(to="modalView")
      transition(name="modal-show")  
        router-view
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'

@Component({
  components: {}
})
export default class Toukei extends Vue {
  public $store!: ExStore
}
</script>

<style lang="sass" scoped></style>
