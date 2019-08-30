<template>
  <span class="badge" :class="badgeClass">
    <slot></slot>
  </span>
</template>

<script>
export default {
  name: 'TagButton',
  props: {
    tag: String
  },
  data() {
    return {
      badgeClass: `badge-${this.tag}`
    }
  }
}
</script>