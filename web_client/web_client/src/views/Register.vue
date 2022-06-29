<template>
  <div class="about">
    <h1>This is an test page</h1>
    <Left></Left>
  </div>
</template>

<script>
import Left from '@/components/HelloWorld.vue';
async function getData() {
  const res = await API({
    url: '/test',
    method: 'get',
  });
  console.log(res);
}
export default {
  name: 'About',
  components: {
    Left,
  },
  setup() {
    return {
      getData,
    };
  },
};



</script>