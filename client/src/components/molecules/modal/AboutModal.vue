<template lang="pug">
  popup-modal-frame(name="about" title="ツンドクについて")
    .about-modal
      .tsundoku-logo
        icon(name="logo" :width="38" :height="38")
        .tsundoku-logo-text
          | ツンドク
      .copy-wrap
        .tsundoku-icon
          icon(name="tsundoku" :width="42" :height="42")
        | を
        .kidoku-icon
          icon(name="kidoku" :width="42" :height="42")
        | に
      .app-description
        p せっかく買った本、積まれていませんか？
        p お金を出したのに読んでないなんて勿体ない
        p ツンドクはそんな「未読負債」をみえるようにするアプリです
        p ツンドク額をみて焦ったり
        p キドク額をみて自分の読書を振り返ったり
        p お金をきっかけに本の価値について考え直してみましょう
      .developers
        p.title
          | ツンドク開発チーム
        .dev-item
          .dev-name
            | うにぽた
            .links
              a(href="https://github.com/unipota" target="_blank")
                icon(name="github" :width="20" :height="20")
              a(href="https://twitter.com/unipota" target="_blank")
                icon(name="twitter" :width="22" :height="22")
          .dev-contributions
            | Direction / Design / Frontend
        .dev-item
          .dev-name
            | すぱ
            .links
              a(href="https://twitter.com/__spaspa__" target="_blank")
                icon(name="twitter" :width="22" :height="22")
          .dev-contributions
            | Design / Frontend
        .dev-item
          .dev-name
            | 60°
            .links
          .dev-contributions
            | Design / Frontend
        .dev-item
          .dev-name
            | とーふとふ
            .links
          .dev-contributions
            | Server / Infrastructure
        .dev-item
          .dev-name
            | oribe
            .links
              a(href="https://github.com/oribe1115" target="_blank")
                icon(name="github" :width="20" :height="20")
          .dev-contributions
            | Server / Infrastructure
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

import Icon from '@/components/assets/Icon.vue'
import PopupModalFrame from '@/components/atoms/PopupModalFrame.vue'

@Component({ components: { Icon, PopupModalFrame } })
export default class AboutModal extends Vue {}
</script>

<style lang="sass" scoped>
.about-modal
  display: flex
  flex-flow: column
  align-items: center

.tsundoku-logo
  display: flex
  align-items: center

.tsundoku-logo-text
  font:
    weight: bold
    size: 32px
  padding:
    left: 16px
  color: var(--text-black)

.copy-wrap
  display: inline-flex
  align-items: center
  margin:
    top: 64px
    bottom: 48px
  font:
    size: 18px
    weight: bold

.tsundoku-icon
  position: relative
  margin:
    right: 24px

  &::before, &::after
    position: absolute
    left: 50%
    white-space: nowrap
    color: var(--tsundoku-red)
    font:
      size: 18px

  &::before
    top: 0
    content: 'ツンドク'
    transform: translate(-50%,-100%)

  &::after
    bottom: 0
    content: '負債'
    transform: translate(-50%,100%)

.kidoku-icon
  position: relative
  margin:
    left: 24px
    right: 24px

  &::before, &::after
    position: absolute
    left: 50%
    white-space: nowrap
    color: var(--kidoku-blue)
    font:
      size: 18px

  &::before
    top: 0
    content: 'キドク'
    transform: translate(-50%,-100%)

  &::after
    bottom: 0
    content: '財産'
    transform: translate(-50%,100%)

.app-description
  text-align: center
  font:
    weight: bold
  color: var(text-black)
  padding: 8px 4px

  p
    margin: 12px auto

.developers
  border:
    top: solid var(--border-gray) 2px
  padding:
    top: 24px
  margin:
    top: 16px
    bottom: 16px
  width: 100%

  .title
    font:
      weight: bold
    text-align: center
    margin:
      bottom: 24px

.dev-item
  width: calc(100% - 12px)
  margin: 8px auto
  padding: 12px 16px
  background: var(--bg-gray)
  border:
    radius: 16px

  .dev-name
    display: flex
    justify-content: space-between
    font:
      size: 18px
      weight: bold

    .links
      display: inline-flex
      align-items: center

      a
        display: inline-flex
        margin:
          right: 8px

  .dev-contributions
    margin:
      top: 4px
    font:
      size: 14px
      weight: bold
</style>
