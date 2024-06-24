<template>
  <a-space size="large">
    <div></div>
    <a-dropdown @select="handleSelect" :popup-max-height="false">
      <a-button> Admin / User <icon-down/></a-button>
      <template #content>
        <a-doption>Option 1</a-doption>
        <a-doption>Option 3</a-doption>
        <a-doption>Log out</a-doption>
      </template>
    </a-dropdown>
  </a-space>
</template>

<script>
export default {
  setup() {
    const handleSelect = (v) => {
      console.log(v)
    };

    return {
      handleSelect
    }
  },
}
</script>

<style>
.arco-dropdown-open .arco-icon-down {
  transform: rotate(180deg);
}
</style>
