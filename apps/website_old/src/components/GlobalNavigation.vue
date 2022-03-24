<template>
  <nav :class="{fixed: fixed}">
    <div class="left-side-wrapper">
      <a href="/">Home</a>
      <a onclick="alert('Diese Seite ist leider noch in der Entwicklung.')">Projekte</a>
      <a href="/#donate">Spenden</a>
    </div>
    <div class="right-side-wrapper">
      <a href="/contact">Kontakt</a>
    </div>
  </nav>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import $ from "jquery";

@Component({
  components: {
  },
})
export default class GlobalNavigation extends Vue {

  fixed = false;
  initialNavBottom = 0;

  mounted() {
    let nav = $('nav');
    this.initialNavBottom = nav.offset().top+nav.height();
    this.fixed = window.pageYOffset > this.initialNavBottom;
    document.onscroll = () => {
      this.fixed = window.pageYOffset > this.initialNavBottom;
    }
  }

}
</script>

<style lang="scss" scoped>

nav {
  position: absolute;
  display: flex;
  flex-flow: row wrap;
  justify-content: space-evenly;
  width: 100%;
  z-index: 100;
  padding: 120px 0 20px 0;
  color: white;

  &.fixed {
    position: fixed;
    background-color: white;
    color: black;
    padding: 40px 0 10px 0;
  }
}

.left-side-wrapper {
  display: flex;
  flex-flow: row wrap;
  column-gap: 7vw;
}

a {
  color: inherit;
  font-size: 20px;
  text-decoration: none;
  transition: 0.1s;
  &.hover {
    transform: scale(1.2);
  }
}

</style>
