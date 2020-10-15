<template>
  <div class="mapview">
    <InformationWidget
      v-if="this.$store.state.mapView.settings.showInfoWidget"
      :latitude="locations[0].Data.latitude"
      :longitude="locations[0].Data.longitude"
      :speed="locations[0].Data.speed"
    />
    <LocationMap
      :latitude="locations[0].Data.latitude"
      :longitude="locations[0].Data.longitude"
    />
    <SettingsModalWindow />
  </div>
</template>

<script lang="ts">
  import { Options, Vue } from "vue-class-component";
  import InformationWidget from "@/components/InformationWidget.vue";
  import LocationMap from "@/components/LocationMap.vue";
  import SettingsModalWindow from "@/components/SettingsModal.vue";

  // Interfaces
  export interface DatabaseEntry extends Record<string, any> {
    ID: number;
    Timestamp: number;
    Data: LocationData;
  }

  export interface LocationData extends Record<string, number> {
    latitude: number;
    longitude: number;
    timestamp: number;
    hdop: number;
    altitude: number;
    speed: number;
  }

  @Options({
    components: {
      InformationWidget,
      LocationMap,
      SettingsModalWindow
    }
  })
  export default class MapView extends Vue {
    // Class variables
    private apiEndpoint = this.getApiEndpoint;
    private TimerId!: number;
    private locations: Array<DatabaseEntry> = [
      {
        ID: 0,
        Timestamp: 0,
        Data: {
          latitude: 0,
          longitude: 0,
          timestamp: 0,
          hdop: 0,
          altitude: 0,
          speed: 0
        }
      }
    ];

    // Functions
    retrieveLocationUpdates = async () => {
      fetch(this.apiEndpoint)
        .then(response => response.json())
        .then(data => (this.locations = data));
    };

    get getApiEndpoint(): string {
      return process.env.VUE_APP_LOCATION_API_ENDPOINT || "http://localhost:8080";
    }

    // Lifecycle hooks
    created() {
      this.retrieveLocationUpdates();
    }

    mounted() {
      this.TimerId = setInterval(() => this.retrieveLocationUpdates(), 2000);
    }

    beforeUnmount() {
      clearInterval(this.TimerId);
    }
  }
</script>

<style lang="stylus">
.mapview
  height 100%
</style>