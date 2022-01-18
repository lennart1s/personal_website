<template>
  <div class="home" style="background: linear-gradient(rgba(2,0,36,1) 0%, rgba(9,9,121,1) 35%, rgba(0,212,255,1) 100%); height: 200vh">
    <div class="title">
      <div class="heading-wrapper">
        <h1>Lennart Sandbothe</h1>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component({
  components: {
  },
})
export default class Home extends Vue {}
</script>

<style lang="scss">

.title {
  width: 100vw;
  height: 100vh;
  background-image: url("../assets/images/titleimg.jpg");
  background-position: center top;
  background-repeat: no-repeat;
  background-size: cover;
}

.heading-wrapper {
  display: flex;
  height: 85%;

  align-items: center;
  justify-content: center;

  & h1 {
    color: white;
    font-size: 4rem;
  }
}

</style>