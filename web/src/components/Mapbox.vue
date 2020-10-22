<template>
  <div class="mapbox">
    <link
      href="https://api.mapbox.com/mapbox-gl-js/v1.12.0/mapbox-gl.css"
      rel="stylesheet"
    />
    <div id="mapContainer" />
  </div>
</template>

<script lang="ts">
  // Imports
  import { Vue, Options } from "vue-class-component";
  import { Prop, Watch } from "vue-property-decorator";
  import mapboxgl from "mapbox-gl";

  // Component options
  @Options({
    components: { mapboxgl }
  })

  // Component
  export default class MapBoxComponent extends Vue {
    // Props
    @Prop({ required: true }) private accessToken!: string;
    @Prop({ required: true, default: 52.132633 }) private latitude!: number;
    @Prop({ required: true, default: 5.291266 }) private longitude!: number;
    @Prop({ required: false, default: 4 }) private zoom!: number;
    @Prop({ required: false, default: true }) private keepCentered!: boolean;

    // Class variables
    private map!: mapboxgl.Map;
    private marker: mapboxgl.Marker = new mapboxgl.Marker();

    // Lifecycle hooks
    mounted() {
      this.map = new mapboxgl.Map({
        accessToken: this.accessToken,
        container: "mapContainer",
        style: "mapbox://styles/ricardobalk/ciyokz20q00382smtd67whbog",
        center: [5.291266, 52.132633], // Center of NL
        zoom: this.zoom
      });
    }

    // Watchers
    @Watch("latitude")
    @Watch("longitude")
    onCoordsChanged() {
      if (this.keepCentered) {
        this.map.setCenter([this.longitude, this.latitude]);
      }
      this.marker.setLngLat([this.longitude, this.latitude]).addTo(this.map);
    }

    @Watch("keepCentered")
    onKeepCenteredChanged() {
      if (this.keepCentered) {
        this.map.setCenter([this.longitude, this.latitude]);
      }
    }
  }
</script>

<style lang="stylus">
.mapbox
  height 100%
  #mapContainer {
    height 100%
  }
</style>