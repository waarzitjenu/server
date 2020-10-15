<template>
  <div class="console">
    <pre><slot /></pre>
  </div>
</template>

<script lang="ts">
  import { Vue } from "vue-class-component";

  export default class Console extends Vue {}
</script>

<style lang="stylus">
.console
    pre
        color #fff
</style>