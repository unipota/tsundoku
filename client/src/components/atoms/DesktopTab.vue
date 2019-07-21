<template lang="pug">
  nav.tab-bar
    router-link.tab(
      v-for="tab in Object.keys(tabs)"
      :key="tab"
      :class="{ [tab]: true, 'active': isActive(tab) }"
      :to="{ name: tabs[tab].to, hash: $route.hash }"
    )
      span.icon(:class="tab")
        component(
          :is="`icon-${tab}`"
          :color="isActive(tab) ? undefined : tabs[tab].inactiveColor"
          :height="48"
        )
      span.label
        | {{ $t(tab) }}
</template>

<script lang="ts">
import { Vue, Prop, Component } from 'vue-property-decorator'
import IconKidoku from '@/components/assets/IconKidoku.vue'
import IconTsundoku from '@/components/assets/IconTsundoku.vue'
import IconToukei from '@/components/assets/IconToukei.vue'

@Component({
  components: {
    IconKidoku,
    IconTsundoku,
    IconToukei
  }
})
export default class DesktopTab extends Vue {
  @Prop({ type: String, default: '' }) selectedTab!: string

  private tabs = {
    tsundoku: {
      to: 'tsundoku',
      inactiveColor: 'var(--tsundoku-red-bg)'
    },
    kidoku: {
      to: 'kidoku',
      inactiveColor: 'var(--kidoku-blue-bg)'
    },
    toukei: {
      to: 'toukei',
      inactiveColor: 'var(--toukei-black-bg)'
    }
  }

  private isActive(tab: string) {
    return this.selectedTab === tab
  }
}
</script>

<style lang="sass" scoped>
.tab-bar
  display: flex
  flex-direction: column
  width: 100%
.tab
  display: flex
  align-items: center
  height: 72px
  width: 100%
  padding: 0 32px
  border-radius: 9999px
  &:not(:last-child)
    margin-bottom: 24px

.icon
  margin-right: 24px

.tsundoku
  color: var(--tsundoku-red-bg)
  &.active
    color: var(--tsundoku-red)
    background-color: var(--tsundoku-red-bg)
.kidoku
  color: var(--kidoku-blue-bg)
  &.active
    color: var(--kidoku-blue)
    background-color: var(--kidoku-blue-bg)
.toukei
  color: var(--toukei-black-bg)
  &.active
    color: var(--toukei-black)
    background-color: var(--toukei-black-bg)

.label
  font-weight: bold
</style>
