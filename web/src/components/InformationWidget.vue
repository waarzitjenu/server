<template>
  <div class="information-widget">
    <span class="coordinates">
      <span class="latitude">
        {{ toDMS({ value: latitude, latitude: true }) }}
      </span>
      <span class="longitude">
        {{ toDMS({ value: longitude, latitude: false }) }}
      </span>
    </span>
    <span class="speed">
      <span class="ms">{{ speed.toPrecision(4).replace(".", ",") }} m/s</span>
      <span class="kmh">{{ toKmh(speed) }} km/h</span>
    </span>
  </div>
</template>

<script lang="ts">
  import { Vue } from "vue-class-component";
  import { Prop } from "vue-property-decorator";

  export default class InformationWidget extends Vue {
    toKmh = (speed: number) => (speed * 3.6).toPrecision(4).replace(".", ",");
    toDMS = (input: { value: number; latitude: boolean }) => {
      interface DMS extends Record<string, any> {
        degrees: number;
        minutes: number;
        seconds: number;
        bearing: "N" | "S" | "E" | "W";
      }

      const dms: DMS = {
        degrees: 0,
        minutes: 0,
        seconds: 0,
        bearing: "N"
      };

      dms.degrees = Math.floor(input.value);
      dms.minutes = Math.floor((input.value - dms.degrees) * 60);
      dms.seconds = Math.round(
        ((input.value - dms.degrees) * 60 - dms.minutes) * 60
      );

      // After rounding, the seconds might become 60. These two
      // if-tests are not necessary if no rounding is done.
      if (dms.seconds == 60) {
        dms.minutes++;
        dms.seconds = 0;
      }
      if (dms.minutes == 60) {
        dms.degrees++;
        dms.minutes = 0;
      }

      // Set the bearing (N, S or E, W)
      if (input.latitude == true) {
        if (input.value >= 0) {
          dms.bearing = "N";
        } else {
          dms.bearing = "S";
        }
      } else {
        if (input.value >= 0) {
          dms.bearing = "E";
        } else {
          dms.bearing = "W";
        }
      }
      // Finally, make degrees always positive (so you won't get -60 degees south, for example)
      dms.degrees = Math.abs(dms.degrees);

      return `${dms.degrees}\u00B0 ${dms.minutes}' ${dms.seconds}" ${dms.bearing}`;
    };

    // Props
    @Prop({ required: true, default: 52.132633 }) private latitude!: number;
    @Prop({ required: true, default: 5.291266 }) private longitude!: number;
    @Prop({ required: true, default: 13.3333 }) private speed!: number;
  }
</script>

<style lang="stylus">
.information-widget
  font-size 1em
  font-weight 600
  color rgba(255,255,255,.75)
  background-color rgba(0,0,0,.75)
  display flex
  flex-direction column
  z-index 10
  position absolute
  padding .25em .5em
  border-radius 0 0 .25em
  .coordinates, .speed
    display grid
    grid-auto-flow column
    grid-gap .5rem

</style>
