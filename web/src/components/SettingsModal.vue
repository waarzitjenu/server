<template>
  <div class="settings-modal">
    <span class="title">Settings</span>

    <div class="settings">
      <span class="setting">
        <input
          type="checkbox"
          id="keep-centered"
          v-model="this.$store.state.mapView.settings.keepCentered"
        />
        <label for="keep-centered">Keep map centered</label>
      </span>

      <span class="setting">
        <input
          type="checkbox"
          id="show-widgets"
          v-model="this.$store.state.mapView.settings.showInfoWidget"
        />
        <label for="show-widgets">Show info widget</label>
      </span>
    </div>
  </div>
</template>

<script lang="ts">
  // Imports
  import { Vue } from "vue-class-component";

  // Component
  export default class SettingsModalWindow extends Vue {}
</script>

<style lang="stylus">
.settings-modal
    z-index 10
    background-color rgba(0,0,0,.66)
    color #fff
    display flex
    flex-direction column
    align-items center
    position absolute
    right 10px
    bottom 10px
    border 1px solid rgba(255,255,255,.33)
    border-radius .5rem
    padding 1em 1.75em

    .title
      display inline-block
      font-size 1.25rem
      text-align center
      margin-bottom 1em
      width 100%

    .settings
        display flex
        flex-direction column

        .setting
            display inline-block
            line-height 2rem
            label
                margin-left .25em
</style>
