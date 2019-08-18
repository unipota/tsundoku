<template lang="pug">
router-link.login-button(to="/login")
  | {{ $t('title_login') }}
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class LoginButton extends Vue {}
</script>

<style lang="sass" scoped>
.login-button
  background: var(--bg-gray)
  color: var(--text-black)
  border-radius: 29px
  padding: 7px 30px
  font-weight: bold
</style>
