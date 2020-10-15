<template>
  <div class="map">
    <MapBoxComponent
      :accessToken="this.accessToken"
      :latitude="latitude"
      :longitude="longitude"
      zoom="2"
      :keepCentered="this.$store.state.mapView.settings.keepCentered"
    />
  </div>
</template>

<script lang="ts">
  // Imports
  import { Vue, Options } from "vue-class-component";
  import { Prop, Watch } from "vue-property-decorator";
  import MapBoxComponent from "@/components/Mapbox.vue";

  // Component options
  @Options({ components: { MapBoxComponent } })
  export default class LocationMap extends Vue {
    // Props
    @Prop({ required: true, default: 52.132633 }) private latitude!: number;
    @Prop({ required: true, default: 5.291266 }) private longitude!: number;

    get accessToken(): string {
      return process.env.VUE_APP_MAPBOX_ACCESS_TOKEN || "";
    }
  }
</script>

<style lang="stylus">
.map
  height 100%
</style>