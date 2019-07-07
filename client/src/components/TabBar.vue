<template lang="pug">
  div.wrapper
    div.tab-bar
      span.tab.tsundoku(
        :class="{'selected': selectedTab === 'tsundoku'}"
        @click="selectTab(tabs.tsundoku.to)"
      )
        span.icon.tsundoku
          IconTsundoku(:color="selectedTab === 'tsundoku' ? undefined: 'var(--tsundoku-red-bg)'")
        span.label(v-if="selectedTab === 'tsundoku'")
          | {{ tabs.tsundoku.label }}

      span.tab.kidoku(
        :class="{'selected': selectedTab === 'kidoku'}"
        @click="selectTab(tabs.kidoku.to)"
      )
        span.icon.tsundoku
          IconKidoku(:color="selectedTab==='kidoku' ? undefined: 'var(--kidoku-blue-bg)'").icon.kidoku
        span.label(v-if="selectedTab === 'kidoku'")
          | {{ tabs.kidoku.label }}
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import IconKidoku from '@/components/Icons/IconKidoku.vue'
import IconTsundoku from '@/components/Icons/IconTsundoku.vue'

@Component({
  components: {
    IconKidoku,
    IconTsundoku,
  },
})
export default class TabBar extends Vue {
  @Prop({ type: String, default: 'tsundoku' })
  private selectedTab!: string

  private tabs = {
    tsundoku: {
      name: 'tsundoku',
      label: 'ツンドク',
      to: 'tsundoku',
    },
    kidoku: {
      name: 'kidoku',
      label: 'キドク',
      to: 'kidoku',
    },
  }

  private selectTab(routeName: string): void {
    this.$router.push({
      name: routeName,
      params: { view: this.$route.params.view },
    })
  }

}
</script>

<style lang="sass">
.wrapper
  width: 100%
  // fix to bottom
  bottom: 0
  position: fixed
  z-index: 10
  @media (min-width: 750px)
    // desktop
    bottom: auto
    top: 0

  .tab-bar
    padding: 12px 18px
    display: flex
    width: auto
    max-width: 375px - (18px * 2);
    margin: auto;

    .tab
      display: flex
      width: max-content
      height: 72px
      border-radius: 44px
      padding: 0 35px
      cursor: pointer

      &.tsundoku
        color: $tsundoku-red
        &.selected
          background-color: $tsundoku-red-bg
          margin:
            right: auto
      &.kidoku
        color: $kidoku-blue
        &.selected
          background-color: $kidoku-blue-bg
          margin:
            left: auto

      .icon
        display: flex
        margin: 
          top: auto
          right: auto
          bottom: auto
          left: 0

      .label
        margin:
          top: auto
          right: 0
          bottom: auto
          left: 27px
        font:
          size: 20px
          weight: bold
</style>
