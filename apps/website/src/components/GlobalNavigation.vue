<template>
  <nav id="navigationbar">
    <a href="/">Home</a>
    <a onclick="alert('Diese Seite ist leider noch in der Entwicklung.')">Projekte</a>
    <a href="/#donate">Spenden</a>
    <a href="/contact">Kontakt</a>
  </nav>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component({
  components: {
  },
})
export default class GlobalNavigation extends Vue {}
</script>

<style lang="scss">

  #navigationbar {
    width: 100%;
    z-index: 100;
  
    & > a {
      margin: 0 1.5rem;

      color: black;
      font-size: 20px;
      text-decoration: none;
      transition: 0.1s;

      &.hover {
        transform: scale(1.2);
      }
    }
  }

</style>
